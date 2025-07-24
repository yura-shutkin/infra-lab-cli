package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var RestartMachineCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart podman machine",
	RunE:  runRestartMachine,
}

func runRestartMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.RestartMachine(binaryName, machineName)
}
