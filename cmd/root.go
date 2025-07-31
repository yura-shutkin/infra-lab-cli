package cmd

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/minikube"
	"infra-lab-cli/cmd/podman"
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
	rootCmd.AddCommand(podman.RootCmd)
	rootCmd.AddCommand(minikube.RootCmd)
	rootCmd.AddCommand(statusCmd)
}
