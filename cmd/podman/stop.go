package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop podman machine",
	RunE:  runStop,
}

func runStop(cmd *cobra.Command, args []string) error {
	return podmansrc.StopMachine(binaryName, machineName)
}
