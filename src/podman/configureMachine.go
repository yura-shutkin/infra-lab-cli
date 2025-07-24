package podman

import (
	"fmt"
	"infra-lab-cli/utils"
	"os/exec"
	"strconv"
)

func isParamChanged(param string, currentValue string) bool {
	if param != "0" && param != "" {
		return param != currentValue
	}
	return false
}

func ConfigureMachine(binaryName, machineName string, params ConfigParams) error {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	machine, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	isChanged := isConfigChanged(params, machine)
	if !isChanged {
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

	if isParamChanged(params.CPUs.Value, strconv.Itoa(machine.Resources.CPUs)) {
		_, err := exec.Command("podman", "machine", "set", "--cpus", params.CPUs.Value, machineName).CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("CPU was updated from %d to %s\n", machine.Resources.CPUs, params.CPUs.Value)
	}

	memMiB, err := utils.ConvertToMiB(params.Memory.Value)
	if err != nil {
		return fmt.Errorf("invalid memory value: %v", err)
	}
	if isParamChanged(memMiB, strconv.Itoa(machine.Resources.Memory)) {
		_, err := exec.Command("podman", "machine", "set", "--memory", memMiB, machineName).CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Memory was updated from %.1fG to %s\n", utils.ConvertMiBToGiB(machine.Resources.Memory), params.Memory.Value)
	}

	if isParamChanged(params.DiskSize.Value, strconv.Itoa(machine.Resources.DiskSize)) {
		// TODO: new size must be greater than current
		if params.DiskSize.Value > strconv.Itoa(machine.Resources.DiskSize) {
			_, err := exec.Command("podman", "machine", "set", "--disk-size", params.DiskSize.Value, machineName).CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Disk size was updated from %d to %s\n", machine.Resources.DiskSize, params.DiskSize.Value)
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

func isConfigChanged(params ConfigParams, machine *InspectedMachine) bool {
	if isParamChanged(params.CPUs.Value, strconv.Itoa(machine.Resources.CPUs)) && params.CPUs.IsProvided {
		return true
	}

	memProvided, err := utils.ConvertToMiB(params.Memory.Value)
	if err != nil {
		fmt.Printf("Invalid memory value: %v\n", err)
		return false
	}

	if isParamChanged(memProvided, strconv.Itoa(machine.Resources.Memory)) && params.Memory.IsProvided {
		return true
	}

	if isParamChanged(params.DiskSize.Value, strconv.Itoa(machine.Resources.DiskSize)) && params.DiskSize.IsProvided {
		return true
	}

	return false
}
