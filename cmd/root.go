package cmd

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/cluster"
	"infra-lab-cli/cmd/vm"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "infra-lab-cli",
	Short: "Local sandbox automation CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(vm.RootCmd)
	rootCmd.AddCommand(cluster.RootCmd)
	rootCmd.AddCommand(statusCmd)
}
