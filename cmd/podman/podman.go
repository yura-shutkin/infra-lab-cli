package podman

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/utils"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "podman",
	Short: "Manage podman machines",
	// TODO: this PreRun does not prevent subcommands from running if the binary is not found
	PreRun: func(cmd *cobra.Command, args []string) {
		if !utils.IsBinaryInPath(binaryName) {
			cmd.Printf("'%s' not found\n", binaryName)
			os.Exit(1)
		}
	},
}

var machineName string
var binaryName = "podman"

func init() {
	// TODO: select default machine as default value, if no default machine is set, use the very first one
	// TODO: Add possibility to autocomplete machine name when using the `--name` flag
	RootCmd.PersistentFlags().StringVarP(&machineName, "name", "n", "podman-machine-default", "Name of the podman machine")
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(StartCmd)
	RootCmd.AddCommand(StopCmd)
	RootCmd.AddCommand(RestartCmd)
	RootCmd.AddCommand(ConfigCmd)
	RootCmd.AddCommand(StatusCmd)
}
