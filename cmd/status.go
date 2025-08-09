package cmd

import (
	"fmt"
	"infra-lab-cli/cmd/cluster/minikube"
	"infra-lab-cli/cmd/vm/podman"

	"github.com/spf13/cobra"
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
