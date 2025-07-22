package podman

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/utils"
)

// GetMachineStatus retrieves and displays the status of the specified podman machine
func GetMachineStatus(binaryName, machineName string) error {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	machines, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	if len(machines) == 0 {
		return fmt.Errorf("machine %s not found", machineName)
	}

	machine := machines[0]
	fmt.Printf("%s\t %s\t %d cpu\t %.1f GiB\t %d GiB\n",
		machine.Name, machine.State,
		machine.Resources.CPUs,
		utils.ConvertMiBToGiB(machine.Resources.Memory),
		machine.Resources.DiskSize)

	return nil
}
