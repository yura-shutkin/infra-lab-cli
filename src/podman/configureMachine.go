package podman

import (
	"fmt"
	"infra-lab-cli/utils"
	"os/exec"
	"strconv"
)

func ConfigureMachine(binaryName, machineName string, params ConfigParams) error {
	// TODO: is it wise to move this check to a function, or this action would not help with code duplication?
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	machine, err := InspectMachine(binaryName, machineName)
	if err != nil {
		return err
	}

	checkIfParamsWereChanged(params, machine)
	if !params.CPUs.IsChanged && !params.Memory.IsChanged && !params.DiskSize.IsChanged {
		fmt.Println("No changes detected in configuration.")
		return nil
	}

	isRunning := machine.State == "running"
	if isRunning {
		err := StopMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	if params.CPUs.IsChanged {
		cmd := fmt.Sprintf("%s machine set --cpus %d %s", binaryName, params.CPUs.Value, machineName)
		_, err := exec.Command(cmd).CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("CPU was updated from %d to %d\n", machine.Resources.CPUs, params.CPUs.Value)
	}

	if params.Memory.IsChanged {
		cmd := fmt.Sprintf("%s machine set --memory %d %s", binaryName, params.Memory.Value, machineName)
		_, err := exec.Command(cmd).CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Memory was updated from %.1fG to %.1f\n", utils.ConvertMiBToGiB(machine.Resources.Memory), utils.ConvertMiBToGiB(params.Memory.Value))
	}

	if params.DiskSize.IsChanged {
		if params.DiskSize.Value > machine.Resources.DiskSize {
			cmd := fmt.Sprintf("%s machine set --disk-size %d %s", binaryName, params.DiskSize.Value, machineName)
			_, err := exec.Command(cmd).CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Disk size was updated from %d to %d\n", machine.Resources.DiskSize, params.DiskSize.Value)
		} else {
			fmt.Println("Disk size must be greater than current size.")
		}
	}

	if isRunning {
		err = StartMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkIfMemoryChanged(param *ConfigParam, currentValue int) {
	var err error
	param.Value, err = utils.ConvertToMiB(param.ValueFlag)
	if err != nil {
		fmt.Printf("Invalid memory value: %v", err)
	}
	if param.Value != currentValue {
		param.IsChanged = true
	}
}

func checkIfParamChanged(param *ConfigParam, currentValue int) {
	var err error
	param.Value, err = strconv.Atoi(param.ValueFlag)
	if err != nil {
		fmt.Printf("Invalid value, should be of Int type: %v\n", err)
	}
	if param.Value != currentValue {
		param.IsChanged = true
	}
}

func checkIfParamsWereChanged(params ConfigParams, machine *InspectedMachine) {
	if params.CPUs.IsProvided {
		checkIfParamChanged(&params.CPUs, machine.Resources.CPUs)
	}

	if params.Memory.IsProvided {
		checkIfMemoryChanged(&params.Memory, machine.Resources.Memory)
	}

	if params.DiskSize.IsProvided {
		checkIfParamChanged(&params.DiskSize, machine.Resources.DiskSize)
	}
}
