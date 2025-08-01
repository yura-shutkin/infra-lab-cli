package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var ListMachinesCmd = &cobra.Command{
	Use:   "list",
	Short: "List podman machines",
	RunE:  runListMachines,
}

func runListMachines(cmd *cobra.Command, args []string) error {
	return podmansrc.ListMachines(binaryName)
}
