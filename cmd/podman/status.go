package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of podman machine",
	RunE:  runStatus,
}

func runStatus(cmd *cobra.Command, args []string) error {
	return podmansrc.GetMachineStatus(binaryName, machineName)
}
