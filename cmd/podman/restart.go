package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart podman machine",
	RunE:  runRestart,
}

func runRestart(cmd *cobra.Command, args []string) error {
	return podmansrc.RestartMachine(binaryName, machineName)
}
