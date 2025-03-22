package apiserver

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tsukiyoz/knowlith/cmd/app/apiserver/options"
	"github.com/tsukiyoz/knowlith/pkg/version"
)

var ServerName = "apiserver"

func NewAPIServerCommand() *cobra.Command {
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		Use:          "apiserver",
		Short:        "apiserver",
		Long:         `apiserver serve api requests.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		Args: cobra.NoArgs,
	}

	return cmd
}

func run(opts *options.ServerOptions) error {
	version.PrintAndExitIfRequested()

	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	if serverCfg := viper.Sub(ServerName); serverCfg != nil {
		if err := serverCfg.Unmarshal(opts); err != nil {
			return err
		}
	}

	if err := opts.Validate(); err != nil {
		return err
	}

	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	return server.Run()
}
