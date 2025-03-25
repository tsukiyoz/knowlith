package cmd

import (
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/tsukiyoz/knowlith/pkg/version"
)

var (
	configFile   string
	registerLock = new(sync.Mutex)
	registry     = make(map[string]*cobra.Command)
)

var rootCommand = &cobra.Command{
	Use:   "knowlith",
	Short: "knowlith",
	Long:  `knowlith is a server about ai and rag backend service written in Go.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Parent() == nil {
			return
		}
		serverInitialize()
	},
}

func NewRootCommand() *cobra.Command {
	addServers()
	rootCommand.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the configuration file.")

	version.AddFlags(rootCommand.PersistentFlags())

	return rootCommand
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := NewRootCommand().Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Register(name string, srv *cobra.Command) {
	registerLock.Lock()
	defer registerLock.Unlock()

	if _, ok := registry[name]; ok {
		panic("duplicate server entry: " + name)
	}

	registry[name] = srv
}

func ListServers() map[string]*cobra.Command {
	registerLock.Lock()
	defer registerLock.Unlock()

	return registry
}

func addServers() error {
	for _, srv := range ListServers() {
		rootCommand.AddCommand(srv)
	}
	return nil
}
