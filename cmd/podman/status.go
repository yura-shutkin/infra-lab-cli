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
		machines[0].Name, machines[0].State,
		machines[0].Resources.CPUs, utils.ConvertMiBToGiB(machines[0].Resources.Memory), machines[0].Resources.DiskSize)
	return nil
}
