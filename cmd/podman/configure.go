package podman

import (
	"github.com/spf13/cobra"
	podmansrc "infra-lab-cli/src/podman"
)

var (
	cpus     string
	memory   string
	diskSize string
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure podman machine",
	RunE:  runConfig,
}

func runConfig(cmd *cobra.Command, args []string) error {
	if !cmd.Flags().Changed("cpus") &&
		!cmd.Flags().Changed("memory") &&
		!cmd.Flags().Changed("disk-size") {
		return cmd.Help()
	}

	params := podmansrc.ConfigParams{
		CPUs:     podmansrc.ConfigParam{Value: cpus, IsProvided: cmd.Flags().Changed("cpus")},
		Memory:   podmansrc.ConfigParam{Value: memory, IsProvided: cmd.Flags().Changed("memory")},
		DiskSize: podmansrc.ConfigParam{Value: diskSize, IsProvided: cmd.Flags().Changed("disk-size")},
	}

	return podmansrc.ConfigureMachine(binaryName, machineName, params)
}

func init() {
	ConfigCmd.Flags().StringVarP(&cpus, "cpus", "c", "", "Number of CPUs to allocate to the podman machine. E.g. 2")
	ConfigCmd.Flags().StringVarP(&memory, "memory", "m", "", "Memory in GiB or in MiB to allocate to the podman machine. E.g. 2G")
	ConfigCmd.Flags().StringVarP(&diskSize, "disk-size", "d", "", "Disk size for the podman machine. E.g. 10G")
}
