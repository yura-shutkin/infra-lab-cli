package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/minikube"
	"infra-lab-cli/cmd/podman"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status over clusters and machines",
	Run:   runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	fmt.Println("Podman machines:")
	_ = podman.StatusCmd.RunE(podman.StatusCmd, args)
	fmt.Println("Minikube clusters:")
	_ = minikube.ListProfilesCmd.RunE(minikube.ListProfilesCmd, args)
}
