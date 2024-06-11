package cmd

import (
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
	Use: "ollamit",
	RunE: func(cmd *cobra.Command, args []string) error {
		g := git.New()
		diff, err := g.DiffStaged()
		if err != nil {
			return err
		}

		cfg := &ollamit.Config{
			DryRun:       flagDryRun,
			Model:        flagModel,
			GitClient:    g,
			OllamaClient: ollama.New(),
		}

		m := ollamit.New(cfg)
		if err := m.Start(diff); err != nil {
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
	rootCmd.Flags().StringVarP(&flagModel, "model", "m", "", "model name")
	_ = rootCmd.MarkFlagRequired("model")
}
