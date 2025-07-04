package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	podman "infra-lab-cli/cmd/podman"
)

var listClustersCmd = &cobra.Command{
	Use:   "status",
	Short: "Status over clusters and machines",
	RunE:  runStatus,
}

func runStatus(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("Podman machines:")
	err = podman.ListCmd.RunE(podman.ListCmd, args)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
