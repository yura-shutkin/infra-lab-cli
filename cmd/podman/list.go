package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List podman machines",
	RunE:  runList,
}

func runList(cmd *cobra.Command, args []string) error {
	return podmansrc.ListMachines(binaryName)
}
