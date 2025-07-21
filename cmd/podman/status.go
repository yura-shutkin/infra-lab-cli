package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/config"
	"infra-lab-cli/utils"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of podman machine",
	RunE:  runStatus,
}

func runStatus(cmd *cobra.Command, args []string) (err error) {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	err = InspectCmd.RunE(InspectCmd, args)
	if err != nil {
		return err
	}
	fmt.Printf("%s\t %s\t %d cpu\t %.1f GiB\t %d GiB\n",
		inspectedMachines[0].Name, inspectedMachines[0].State,
		inspectedMachines[0].Resources.CPUs, utils.ConvertMiBToGiB(inspectedMachines[0].Resources.Memory), inspectedMachines[0].Resources.DiskSize)
	return nil
}
