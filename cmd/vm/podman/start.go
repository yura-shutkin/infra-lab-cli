package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var StartMachineCmd = &cobra.Command{
	Use:   "start",
	Short: "Start podman machine",
	RunE:  runStartMachine,
}

func runStartMachine(cmd *cobra.Command, args []string) error {
	return podmansrc.StartMachine(binaryName, machineName)
}
