package podman

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func RestartMachine(binaryName, machineName string) (err error) {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	machine, err := InspectMachine(binaryName, machineName)
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
