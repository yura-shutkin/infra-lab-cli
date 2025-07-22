package podman

import (
	"fmt"
	"infra-lab-cli/utils"
)

func RestartMachine(binaryName, machineName string) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	machine, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	if machine.State == "running" {
		err = StopMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	err = StartMachine(binaryName, machineName)
	if err != nil {
		return err
	}

	return nil
}
