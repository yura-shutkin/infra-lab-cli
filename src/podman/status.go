package podman

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func GetMachineStatus(binaryName, machineName string) error {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	machine, err := InspectMachine(binaryName, machineName)
	if err != nil {
		return err
	}

	fmt.Printf("%s\t %s\t %d cpu\t %.1f GiB\t %d GiB\n",
		machine.Name, machine.State,
		machine.Resources.CPUs,
		common.ConvertMiBToGiB(machine.Resources.Memory),
		machine.Resources.DiskSize)

	return nil
}
