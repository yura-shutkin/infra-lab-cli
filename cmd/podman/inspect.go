package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var InspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect podman machine",
	RunE:  runInspect,
}

var inspectedMachines []podmansrc.InspectedMachine

func runInspect(cmd *cobra.Command, args []string) error {
	machines, err := podmansrc.InspectMachine(machineName)
	if err != nil {
		return err
	}
	inspectedMachines = machines
	return nil
}
