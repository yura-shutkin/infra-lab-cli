package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start podman machine",
	RunE:  runStart,
}

func runStart(cmd *cobra.Command, args []string) error {
	return podmansrc.StartMachine(binaryName, machineName)
}
