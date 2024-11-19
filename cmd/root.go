package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"natthan.com/go-play/cmd/core"
)

var (
	env string

	// TODO: Read Cobra Doc
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "serve services",
	}
)

func init() {
	// TODO: Read Cobra PersistentFlags
	// TODO: Read Cobra AddCommand
	serveCmd.PersistentFlags().
		StringVar(&env, "env", "", "environment of the application: local, dev, prod")

	serveCmd.AddCommand(core.RootCmd)
}

func Execute() {
	// TODO: Read Cobra Execute
	if err := serveCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// TODO: Read sync.WaitGroup
var wg sync.WaitGroup

// TODO: Read cobra.Command
var ServeAllCmd = &cobra.Command{
	Use:   "all",
	Short: "serve all services",
	Run: func(cmd *cobra.Command, args []string) {
		wg.Add(4)

		go func() {
			core.RootCmd.Run(core.RootCmd, core.RootCmd.Flags().Args())
			wg.Done()
		}()
		wg.Wait()
	},
}
