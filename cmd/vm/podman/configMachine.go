package podman

import (
	podmansrc "infra-lab-cli/src/podman"

	"github.com/spf13/cobra"
)

var (
	cpus     string
	memory   string
	diskSize string
)

var ConfigMachineCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "Configure podman machine",
	RunE:    runConfigMachine,
}

func runConfigMachine(cmd *cobra.Command, args []string) error {
	if !cmd.Flags().Changed("cpus") &&
		!cmd.Flags().Changed("memory") &&
		!cmd.Flags().Changed("disk-size") {
		return cmd.Help()
	}

	params := podmansrc.ConfigParams{
		CPUs:     podmansrc.ConfigParam{ValueFlag: cpus, IsProvided: cmd.Flags().Changed("cpus")},
		Memory:   podmansrc.ConfigParam{ValueFlag: memory, IsProvided: cmd.Flags().Changed("memory")},
		DiskSize: podmansrc.ConfigParam{ValueFlag: diskSize, IsProvided: cmd.Flags().Changed("disk-size")},
	}

	return podmansrc.ConfigureMachine(binaryName, machineName, params)
}

func init() {
	ConfigMachineCmd.Flags().StringVarP(&cpus, "cpus", "c", "2", "Number of CPUs to allocate to the podman machine")
	ConfigMachineCmd.Flags().StringVarP(&memory, "memory", "m", "2G", "Memory in GiB or in MiB to allocate to the podman machine. E.g. 2G, 2048, 2048M")
	ConfigMachineCmd.Flags().StringVarP(&diskSize, "disk-size", "d", "40", "Disk size in GiB for the podman machine")
}
