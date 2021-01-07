package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/zapr"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-projects/pkg/version"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log"
	zaputil "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	LOG_LEVEL  = "LOG_LEVEL"
	DEBUG_MODE = "DEBUG_MODE"
)

func CreateRootContext(customCtx context.Context, name string) context.Context {
	rootCtx := customCtx
	if rootCtx == nil {
		rootCtx = context.Background()
	}

	setupLogging(rootCtx, name, zap.String("version", version.Version))
	return setupSignalHandler(rootCtx)
}

func setupLogging(ctx context.Context, name string, fields ...zap.Field) {

	// Default to info level logging
	level := zapcore.InfoLevel
	var devMode bool
	if envLogLevel := os.Getenv(LOG_LEVEL); envLogLevel != "" {
		// If log level is set in ENV use that log level
		// Available levels can be found here: https://pkg.go.dev/go.uber.org/zap/zapcore?tab=doc#Level
		// Can use either upper case or lower case, ex: (DEBUG/debug)
		if err := (&level).Set(envLogLevel); err != nil {
			contextutils.LoggerFrom(ctx).Infof("Could not set log level from env %s=%s, available levels"+
				"can be found here: https://pkg.go.dev/go.uber.org/zap/zapcore?tab=doc#Level",
				LOG_LEVEL,
				envLogLevel,
				zap.Error(err),
			)
		}
	}
	// If set will override LOG_LEVEL set above
	if debugMode := os.Getenv(DEBUG_MODE); debugMode != "" {
		// if DEBUG_MODE is set, enable debug mode in zap logger, and set level to debug
		level = zapcore.DebugLevel
		devMode = true
	}

	atomicLevel := zap.NewAtomicLevelAt(level)
	baseLogger := zaputil.NewRaw(
		zaputil.Level(&atomicLevel),
		// Only set debug mode if specified. This will use a non-json (human readable) encoder which makes it impossible
		// to use any json parsing tools for the log. Should only be enabled explicitly
		zaputil.UseDevMode(devMode),
		zaputil.RawZapOpts(zap.Fields(fields...)),
	).Named(name)

	// klog
	zap.ReplaceGlobals(baseLogger)
	// controller-runtime
	log.SetLogger(zapr.NewLogger(baseLogger))
	// go-utils
	contextutils.SetFallbackLogger(baseLogger.Sugar())
}

func setupSignalHandler(parentCtx context.Context) context.Context {
	ctx, cancel := context.WithCancel(parentCtx)
	c := make(chan os.Signal, 2)
	signal.Notify(c, []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}...)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(129) // second signal. Exit directly.
	}()

	return ctx
}
