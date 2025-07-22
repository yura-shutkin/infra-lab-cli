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

	// TODO: 14-24 Duplicated code (configure.go)
	var machines []InspectedMachine
	machines, err = InspectMachine(machineName)
	if err != nil {
		return err
	}

	if len(machines) == 0 {
		return fmt.Errorf("machine %s not found", machineName)
	}

	machine := machines[0]

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
