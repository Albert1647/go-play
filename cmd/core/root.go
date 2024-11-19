package core

import (
	"log/slog"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/spf13/cobra"
	// "google.golang.org/api/config/v1"
)

const (
	defaultIdleTimeout    = time.Minute
	defaultReadTimeout    = 5 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	defaultShutdownPeriod = 30 * time.Second
)

var RootCmd = &cobra.Command{
	Use:   "core",
	Short: "serve core service",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			logger = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
					Level: slog.LevelDebug,
				}),
			)
			env = cmd.Flag("env").Value.String()
		)

		if err := run(logger, config.Env(env)); err != nil {
			trace := string(debug.Stack())
			logger.Error(err.Error(), "trace", trace)
			os.Exit(1)
		}
	},
}

func run(logger *slog.Logger, env config.Env) error {
	cfg := config.New(env)
	return nil
}

type application struct {
	logger *slog.Logger
	config *configs.Config

	userRepo *userrepo.Repository
	wg       sync.WaitGroup
}
