package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "podman",
	Short: "Manage podman machines",
}

// TODO: consider to define config variable
var machineName string
var binaryName = "podman"
var defaultMachineName string
var connections []podmansrc.Connection

func machineNameCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	machineNames, err := podmansrc.GetMachineNames(&connections)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	return machineNames, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	// TODO: Select the default machine name based on the default system connection
	// TODO: Add possibility to autocomplete machine name when using the `--name` flag. Correlated with the previous TODO.
	RootCmd.PersistentFlags().StringVarP(&binaryName, "binary", "b", binaryName, "Binary to use")

	_ = podmansrc.GetConnections(binaryName, &connections)
	_ = podmansrc.GetDefaultMachineName(&connections, &defaultMachineName)

	RootCmd.PersistentFlags().StringVarP(&machineName, "name", "n", defaultMachineName, "Name of the podman machine")
	_ = RootCmd.RegisterFlagCompletionFunc("name", machineNameCompletion)

	RootCmd.AddCommand(ListMachinesCmd)
	RootCmd.AddCommand(StartMachineCmd)
	RootCmd.AddCommand(StopMachineCmd)
	RootCmd.AddCommand(RestartMachineCmd)
	RootCmd.AddCommand(ConfigMachineCmd)
	RootCmd.AddCommand(StatusCmd)
}
