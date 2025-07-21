package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/podman"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status over clusters and machines",
	Run:   runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	fmt.Println("Podman machines:")
	err := podman.StatusCmd.RunE(podman.StatusCmd, args)
	if err != nil {
		fmt.Println(err)
	}
}
