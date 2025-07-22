package podman

import (
	"fmt"
	"infra-lab-cli/utils"
)

func GetMachineStatus(binaryName, machineName string) error {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
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
