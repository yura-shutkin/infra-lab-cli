package cmd

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/podman"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "infra-lab-cli",
	Short: "Local sandbox automation CLI",
}

var machineName string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(podman.RootCmd)
}
