package create

import (
	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	"github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/argsutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/flagutils"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/surveyutils"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/cliutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	optionsExt "github.com/solo-io/solo-projects/projects/gloo/cli/pkg/cmd/options"
	flagutilsExt "github.com/solo-io/solo-projects/projects/gloo/cli/pkg/flagutils"
	surveyutilsExt "github.com/solo-io/solo-projects/projects/gloo/cli/pkg/surveyutils"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/ratelimit"
	ratelimit2 "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/ratelimit"
	"github.com/spf13/cobra"
)

func VSCreate(opts *options.Options, optionsFunc ...cliutils.OptionsFunc) *cobra.Command {

	optsExt := &optionsExt.RateLimit{}
	cmd := &cobra.Command{
		// Use command constants to aid with replacement.
		Use:     constants.VIRTUAL_SERVICE_COMMAND.Use,
		Aliases: constants.VIRTUAL_SERVICE_COMMAND.Aliases,
		Short:   "Create a Virtual Service",
		Long: "A virtual service describes the set of routes to match for a set of domains. \n" +
			"Virtual services are containers for routes assigned to a domain or set of domains. \n" +
			"Virtual services must not have overlapping domains, as the virtual service to match a request " +
			"is selected by the Host header (in HTTP1) or :authority header (in HTTP2). " +
			"The routes within a virtual service ",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if opts.Top.Interactive {
				if err := surveyutils.AddVirtualServiceFlagsInteractive(&opts.Create.VirtualService); err != nil {
					return err
				}
				if err := surveyutilsExt.AddVirtualServiceFlagsInteractive(optsExt); err != nil {
					return err
				}
			}
			err := argsutils.MetadataArgsParse(opts, args)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return createVirtualService(opts, optsExt, args)
		},
	}

	pflags := cmd.PersistentFlags()
	flagutils.AddMetadataFlags(pflags, &opts.Metadata)
	flagutils.AddVirtualServiceFlags(pflags, &opts.Create.VirtualService)
	flagutilsExt.AddVirtualServiceFlags(pflags, optsExt)
	cliutils.ApplyOptions(cmd, optionsFunc)

	return cmd
}

func createVirtualService(opts *options.Options, optsExt *optionsExt.RateLimit, args []string) error {
	vs, err := virtualServiceFromOpts(opts.Metadata, opts.Create.VirtualService, *optsExt)
	if err != nil {
		return err
	}
	virtualServiceClient := helpers.MustVirtualServiceClient()
	vs, err = virtualServiceClient.Write(vs, clients.WriteOpts{})
	if err != nil {
		return err
	}

	helpers.PrintVirtualServices(v1.VirtualServiceList{vs}, opts.Top.Output)

	return nil
}

func virtualServiceFromOpts(meta core.Metadata, input options.InputVirtualService, rl optionsExt.RateLimit) (*v1.VirtualService, error) {
	if len(input.Domains) == 0 {
		input.Domains = constants.DefaultDomains
	}
	vs := &v1.VirtualService{
		Metadata: meta,
		VirtualHost: &gloov1.VirtualHost{
			Domains: input.Domains,
		},
	}
	if rl.Enable {
		if vs.VirtualHost.VirtualHostPlugins == nil {
			vs.VirtualHost.VirtualHostPlugins = &gloov1.VirtualHostPlugins{}
		}
		timeUnit, ok := ratelimit.RateLimit_Unit_value[rl.TimeUnit]
		if !ok {
			return nil, errors.Errorf("invalid time unit specified: %v", rl.TimeUnit)
		}
		ingressRateLimit := &ratelimit.IngressRateLimit{
			AnonymousLimits: &ratelimit.RateLimit{
				Unit:            ratelimit.RateLimit_Unit(timeUnit),
				RequestsPerUnit: rl.RequestsPerTimeUnit,
			},
		}
		ingressRateLimitStruct, err := envoyutil.MessageToStruct(ingressRateLimit)
		if err != nil {
			return nil, errors.Wrapf(err, "Error marshalling ingress rate limit")
		}
		if vs.VirtualHost.VirtualHostPlugins.Extensions == nil {
			vs.VirtualHost.VirtualHostPlugins.Extensions = new(gloov1.Extensions)
		}
		if vs.VirtualHost.VirtualHostPlugins.Extensions.Configs == nil {
			vs.VirtualHost.VirtualHostPlugins.Extensions.Configs = make(map[string]*types.Struct)
		}
		vs.VirtualHost.VirtualHostPlugins.Extensions.Configs[ratelimit2.ExtensionName] = ingressRateLimitStruct
	}

	return vs, nil
}
