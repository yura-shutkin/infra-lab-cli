package podman

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/utils"
	"os/exec"
	"strconv"
)

// ConfigureParams holds the configuration parameters for a machine
type ConfigureParams struct {
	CPUs     string
	Memory   string
	DiskSize string
}

// ConfigureMachine configures the specified podman machine with the given parameters
func ConfigureMachine(binaryName, machineName string, params ConfigureParams) error {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	// Get current machine state
	machines, err := InspectMachine(machineName)
	if err != nil {
		return err
	}

	if len(machines) == 0 {
		return fmt.Errorf("machine %s not found", machineName)
	}

	machine := machines[0]

	// Check if any configuration has changed
	isChanged := isConfigChanged(params, machine)
	if !isChanged {
		fmt.Println("No changes detected in configuration.")
		return nil
	}

	// Stop machine if running
	wasRunning := machine.State == "running"
	if wasRunning {
		err = StopMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	// Update CPU configuration
	if isParamChanged(params.CPUs, strconv.Itoa(machine.Resources.CPUs)) {
		out, err := exec.Command("podman", "machine", "set", "--cpus", params.CPUs).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("CPU was updated from %d to %s\n", machine.Resources.CPUs, params.CPUs)
	}

	// Update memory configuration
	memMiB, err := utils.ConvertToMiB(params.Memory)
	if err == nil && isParamChanged(memMiB, strconv.Itoa(machine.Resources.Memory)) {
		out, err := exec.Command("podman", "machine", "set", "--memory", memMiB).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Memory was updated from %d to %s\n", machine.Resources.Memory, params.Memory)
	}

	// Update disk size configuration
	if isParamChanged(params.DiskSize, strconv.Itoa(machine.Resources.DiskSize)) {
		// TODO: new size must be greater than current
		out, err := exec.Command("podman", "machine", "set", "--disk-size", params.DiskSize).CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Disk size was updated from %d to %s\n", machine.Resources.DiskSize, params.DiskSize)
	}

	// Restart machine if it was running
	if wasRunning {
		err = StartMachine(binaryName, machineName)
		if err != nil {
			return err
		}
	}

	return nil
}

// IsParamChanged checks if a parameter has changed from its current value
func IsParamChanged(param string, currentValue string) bool {
	if param != "0" && param != "" {
		return param != currentValue
	}
	return false
}

// isParamChanged is an internal helper function
func isParamChanged(param string, currentValue string) bool {
	return IsParamChanged(param, currentValue)
}

// isConfigChanged checks if any configuration parameter has changed
func isConfigChanged(params ConfigureParams, machine InspectedMachine) bool {
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
