package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tsukiyoz/knowlith/cmd/apiserver/app/options"
	"github.com/tsukiyoz/knowlith/internal/pkg/log"
	"github.com/tsukiyoz/knowlith/pkg/bootstrap"
	"github.com/tsukiyoz/knowlith/pkg/version"
)

var configFile string

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

	cobra.OnInitialize(bootstrap.OnInitialize(&configFile, "KNOWLITH", searchDirs(), defaultConfigName))

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to config file.")

	version.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run(opts *options.ServerOptions) error {
	version.PrintAndExitIfRequested()

	log.Init(logOptions())
	defer log.Sync()

	if err := viper.Unmarshal(opts); err != nil {
		return err
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

func logOptions() *log.Options {
	opts := log.NewOptions()
	if viper.IsSet("log.disable-caller") {
		opts.DisableCaller = viper.GetBool("log.disable-caller")
	}
	if viper.IsSet("log.disable-stacktrace") {
		opts.DisableStacktrace = viper.GetBool("log.disable-stacktrace")
	}
	if viper.IsSet("log.level") {
		opts.Level = viper.GetString("log.level")
	}
	if viper.IsSet("log.format") {
		opts.Format = viper.GetString("log.format")
	}
	if viper.IsSet("log.output-paths") {
		opts.OutputPaths = viper.GetStringSlice("log.output-paths")
	}
	return opts
}
