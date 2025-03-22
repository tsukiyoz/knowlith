package cmd

import (
	"os"

	"github.com/spf13/cobra"
	apiserver "github.com/tsukiyoz/knowlith/cmd/app/apiserver"
	"github.com/tsukiyoz/knowlith/pkg/version"
)

var configFile string

func NewRootCommand() *cobra.Command {
	cobra.OnInitialize(onInitialize)

	command := &cobra.Command{
		Use:   "knowlith",
		Short: "knowlith",
		Long:  `knowlith is a server about ai and rag backend service written in Go.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	command.AddCommand(apiserver.NewAPIServerCommand())
	command.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the configuration file.")

	version.AddFlags(command.PersistentFlags())

	return command
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}
