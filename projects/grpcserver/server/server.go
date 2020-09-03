package server

import (
	"context"
	"net"

	"github.com/solo-io/solo-projects/projects/grpcserver/server/logging"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/client"

	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/solo-projects/projects/grpcserver/api/v1"
	"google.golang.org/grpc"
)

type GlooGrpcService struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGlooGrpcService(
	ctx context.Context,
	listener net.Listener,
	upstreamService v1.UpstreamApiServer,
	upstreamGroupService v1.UpstreamGroupApiServer,
	artifactService v1.ArtifactApiServer,
	configService v1.ConfigApiServer,
	secretService v1.SecretApiServer,
	virtualService v1.VirtualServiceApiServer,
	routeTable v1.RouteTableApiServer,
	gatewayService v1.GatewayApiServer,
	proxyService v1.ProxyApiServer,
	envoyService v1.EnvoyApiServer,
	clientUpdater client.Updater,
) *GlooGrpcService {

	logger := contextutils.LoggerFrom(ctx).Desugar()

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(logger)

	glooServer := &GlooGrpcService{
		server: grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				// Logs generic info about responses (e.g. errors, code, etc)
				grpc_zap.UnaryServerInterceptor(logger, logOkAtDebugLevel()),
				// Logs request/response payloads at the debug level
				logging.RequestResponseDebugInterceptor(logger),
			),
		),
		listener: listener,
	}

	v1.RegisterUpstreamApiServer(glooServer.server, upstreamService)
	v1.RegisterUpstreamGroupApiServer(glooServer.server, upstreamGroupService)
	v1.RegisterArtifactApiServer(glooServer.server, artifactService)
	v1.RegisterConfigApiServer(glooServer.server, configService)
	v1.RegisterSecretApiServer(glooServer.server, secretService)
	v1.RegisterVirtualServiceApiServer(glooServer.server, virtualService)
	v1.RegisterRouteTableApiServer(glooServer.server, routeTable)
	v1.RegisterGatewayApiServer(glooServer.server, gatewayService)
	v1.RegisterProxyApiServer(glooServer.server, proxyService)
	v1.RegisterEnvoyApiServer(glooServer.server, envoyService)

	// just responsible for kicking off the settings watch loop that rebuilds all the clients
	// (the client updater has to be passed somewhere, otherwise wire complains about an unused provider)
	clientUpdater.StartWatch(ctx)

	return glooServer
}

func (s *GlooGrpcService) Run(ctx context.Context) error {
	contextutils.LoggerFrom(ctx).Infow("Starting gloo grpc server")
	return s.server.Serve(s.listener)
}

func (s *GlooGrpcService) Stop() {
	s.server.Stop()
}

// Option to log OK responses only at Debug level to avoid flooding the logs
func logOkAtDebugLevel() grpc_zap.Option {
	return grpc_zap.WithLevels(func(code codes.Code) zapcore.Level {
		if code == codes.OK {
			return zapcore.DebugLevel
		}
		return grpc_zap.DefaultCodeToLevel(code)
	})
}
