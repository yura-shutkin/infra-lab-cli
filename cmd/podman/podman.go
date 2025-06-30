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
	RootCmd.PersistentFlags().StringVarP(&machineName, "machine-name", "m", "podman-machine-default", "Name of the podman machine")
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(StartCmd)
	RootCmd.AddCommand(StopCmd)
	RootCmd.AddCommand(RestartCmd)
}
