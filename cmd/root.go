package cmd

import (
	"errors"
	"os"

	"github.com/koki-develop/ollamit/internal/git"
	"github.com/koki-develop/ollamit/internal/ollama"
	"github.com/koki-develop/ollamit/internal/ollamit"
	"github.com/spf13/cobra"
)

var (
	flagDryRun bool   // --dry-run
	flagModel  string // --model
)

var rootCmd = &cobra.Command{
	Use:  "ollamit",
	Long: "A command-line tool to generate commit messages with ollama.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if flagModel == "" {
			return errors.New("model name is required")
		}

		cfg := &ollamit.Config{
			DryRun:       flagDryRun,
			Model:        flagModel,
			GitClient:    git.New(),
			OllamaClient: ollama.New(),
		}

		m := ollamit.New(cfg)
		if err := m.Start(); err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&flagDryRun, "dry-run", false, "dry run")
	rootCmd.Flags().StringVarP(&flagModel, "model", "m", os.Getenv("OLLAMIT_MODEL"), "model name")
}
