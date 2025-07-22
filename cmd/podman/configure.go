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
	Use:   "config machine",
	Short: "Configure podman machine",
	RunE:  runConfig,
}

func runConfig(cmd *cobra.Command, args []string) error {
	if cpus == "0" && memory == "0" && diskSize == "0" {
		return cmd.Help()
	}

	params := podmansrc.ConfigureParams{
		CPUs:     cpus,
		Memory:   memory,
		DiskSize: diskSize,
	}

	return podmansrc.ConfigureMachine(binaryName, machineName, params)
}

func init() {
	ConfigCmd.Flags().StringVarP(&cpus, "cpus", "c", "0", "Number of CPUs to allocate to the podman machine")
	ConfigCmd.Flags().StringVarP(&memory, "memory", "m", "0", "Memory in GiB or in MiB to allocate to the podman machine")
	ConfigCmd.Flags().StringVarP(&diskSize, "disk-size", "d", "0", "Disk size for the podman machine")
}
