package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/podman"
	"infra-lab-cli/utils"
	"os"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status over clusters and machines",
	RunE:  runStatus,
}

func runStatus(cmd *cobra.Command, args []string) (err error) {
	fmt.Println("Podman machines:")
	// TODO: DRY. Need to find a way to avoid repeating this check in every command
	if utils.IsBinaryInPath("podman") {
		err = podman.StatusCmd.RunE(podman.StatusCmd, args)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return err
	} else {
		cmd.Printf("'%s' not found\n", "podman")
		os.Exit(1)
	}
	return nil
}
