package ratelimit

import (
	"github.com/gogo/protobuf/types"
	editOptions "github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/edit/options"

	"github.com/solo-io/gloo/pkg/cliutil"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/utils"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"github.com/solo-io/solo-projects/projects/gloo/cli/pkg/constants"
	ratelimitpb "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/ratelimit"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/ratelimit"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func RateLimitConfig(opts *editOptions.EditOptions, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {

	optsExt := &RateLimitSettings{}

	cmd := &cobra.Command{
		// Use command constants to aid with replacement.
		Use:     constants.CONFIG_RATELIMIT_COMMAND.Use,
		Aliases: constants.CONFIG_RATELIMIT_COMMAND.Aliases,
		Short:   "Configure rate limit settings (Enterprise)",
		Long:    "Let gloo know the location of the rate limit server. This is a Gloo Enterprise feature.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Top.Interactive {
				if err := AddSettingsInteractive(optsExt); err != nil {
					return err
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return editSettings(opts, optsExt, args)
		},
	}

	AddConfigFlagsSettings(cmd.Flags(), optsExt)
	cliutils.ApplyOptions(cmd, optionsFunc)

	cmd.AddCommand(RateLimitCustomConfig(opts))
	return cmd
}

func editSettings(opts *editOptions.EditOptions, optsExt *RateLimitSettings, args []string) error {
	settingsClient := helpers.MustSettingsClient()
	settings, err := settingsClient.Read(opts.Metadata.Namespace, opts.Metadata.Name, clients.ReadOpts{})
	if err != nil {
		return errors.Wrapf(err, "Error reading settings")
	}

	var rlSettings ratelimitpb.Settings
	err = utils.UnmarshalExtension(settings, ratelimit.ExtensionName, &rlSettings)
	if err != nil {
		if err != utils.NotFoundError {
			return err
		}
	}
	if rlSettings.RatelimitServerRef == nil {
		rlSettings.RatelimitServerRef = new(core.ResourceRef)
	}
	if optsExt.RateLimitServerUpstreamRef.Name != "" {
		rlSettings.RatelimitServerRef.Name = optsExt.RateLimitServerUpstreamRef.Name
	}
	if optsExt.RateLimitServerUpstreamRef.Namespace != "" {
		rlSettings.RatelimitServerRef.Namespace = optsExt.RateLimitServerUpstreamRef.Namespace
	}

	if settings.Extensions == nil {
		settings.Extensions = &gloov1.Extensions{}
	}

	if settings.Extensions.Configs == nil {
		settings.Extensions.Configs = make(map[string]*types.Struct)
	}

	rlStruct, err := protoutils.MarshalStruct(&rlSettings)
	if err != nil {
		return err
	}
	settings.Extensions.Configs[ratelimit.ExtensionName] = rlStruct

	_, err = settingsClient.Write(settings, clients.WriteOpts{OverwriteExisting: true})
	return err
}

func AddSettingsInteractive(opts *RateLimitSettings) error {

	if err := cliutil.GetStringInput("name of the ratelimit server upstream: ", &opts.RateLimitServerUpstreamRef.Name); err != nil {
		return err
	}
	if err := cliutil.GetStringInput("namespace of the ratelimit server upstream: ", &opts.RateLimitServerUpstreamRef.Namespace); err != nil {
		return err
	}
	return nil
}

type RateLimitSettings struct {
	RateLimitServerUpstreamRef core.ResourceRef
}

func AddConfigFlagsSettings(set *pflag.FlagSet, opts *RateLimitSettings) {
	set.StringVar(&opts.RateLimitServerUpstreamRef.Name, "ratelimit-server-name", "", "name of the ext rate limit upstream")
	set.StringVar(&opts.RateLimitServerUpstreamRef.Namespace, "ratelimit-server-namespace", "", "namespace of the ext rate limit upstream")
}
