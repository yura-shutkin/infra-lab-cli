package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var StopMachineCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop podman machine",
	RunE:  runStopMachine,
}

func runStopMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.StopMachine(binaryName, machineName)
}
