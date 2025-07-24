package podman

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "podman",
	Short: "Manage podman machines",
}

var machineName string
var binaryName = "podman"

func init() {
	// TODO: Select the default machine name based on the default system connection
	// TODO: Add possibility to autocomplete machine name when using the `--name` flag. Correlated with the previous TODO.
	RootCmd.PersistentFlags().StringVarP(&machineName, "name", "n", "podman-machine-default", "Name of the podman machine")
	RootCmd.AddCommand(ListMachinesCmd)
	RootCmd.AddCommand(StartMachineCmd)
	RootCmd.AddCommand(StopMachineCmd)
	RootCmd.AddCommand(RestartMachineCmd)
	RootCmd.AddCommand(ConfigMachineCmd)
	RootCmd.AddCommand(StatusCmd)
}
