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

func ConfigureMachine(binaryName, machineName string, params ConfigureParams) error {
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

	wasRunning := machine.State == "running"
	if wasRunning {
		err = StopMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	if isParamChanged(params.CPUs, strconv.Itoa(machine.Resources.CPUs)) {
		out, err := exec.Command("podman", "machine", "set", "--cpus", params.CPUs).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("CPU was updated from %d to %s\n", machine.Resources.CPUs, params.CPUs)
	}

	memMiB, err := utils.ConvertToMiB(params.Memory)
	if err == nil && isParamChanged(memMiB, strconv.Itoa(machine.Resources.Memory)) {
		out, err := exec.Command("podman", "machine", "set", "--memory", memMiB).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Memory was updated from %d to %s\n", machine.Resources.Memory, params.Memory)
	}

	if isParamChanged(params.DiskSize, strconv.Itoa(machine.Resources.DiskSize)) {
		// TODO: new size must be greater than current
		out, err := exec.Command("podman", "machine", "set", "--disk-size", params.DiskSize).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Disk size was updated from %d to %s\n", machine.Resources.DiskSize, params.DiskSize)
	}

	if wasRunning {
		err = StartMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	return nil
}

func isConfigChanged(params ConfigureParams, machine *InspectedMachine) bool {
	if isParamChanged(params.CPUs, strconv.Itoa(machine.Resources.CPUs)) {
		return true
	}

	if isParamChanged(params.Memory, strconv.Itoa(machine.Resources.Memory)) {
		return true
	}

	if isParamChanged(params.DiskSize, strconv.Itoa(machine.Resources.DiskSize)) {
		return true
	}

	return false
}
