package podman

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "podman",
	Short: "Manage podman machines",
}

var machineName string

func init() {
	RootCmd.PersistentFlags().StringVarP(&machineName, "name", "n", "podman-machine-default", "Name of the podman machine")
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(startCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(restartCmd)
	RootCmd.AddCommand(configCmd)
}
