package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/config"
	"infra-lab-cli/utils"
	"os/exec"
	"strconv"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure podman machine",
	RunE:  runConfig,
}

func isParamChanged(param string, currentValue string) bool {
	if param != "0" {
		if param != currentValue {
			return true
		}
	}
	return false
}

func isConfigChanged(cmd *cobra.Command, args []string) (isChanged bool, err error) {
	err = InspectCmd.RunE(InspectCmd, args)
	if err != nil {
		return isChanged, err
	}

	if isParamChanged(cpus, strconv.Itoa(machines[0].Resources.CPUs)) {
		return true, nil
	}

	if isParamChanged(memory, strconv.Itoa(machines[0].Resources.Memory)) {
		return true, nil
	}

	if isParamChanged(diskSize, strconv.Itoa(machines[0].Resources.DiskSize)) {
		return true, nil
	}

	return false, nil
}

func runConfig(cmd *cobra.Command, args []string) (err error) {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	if cpus == "0" && memory == "0" && diskSize == "0" {
		return cmd.Help()
	}

	isChanged, err := isConfigChanged(cmd, args)
	if err != nil {
		return err
	}

	if isChanged {
		if machines[0].State == "running" {
			err = StopCmd.RunE(StopCmd, args)
			if err != nil {
				return err
			}
		}

		// Update CPU
		if isParamChanged(cpus, strconv.Itoa(machines[0].Resources.CPUs)) {
			out, err := exec.Command(
				"podman", "machine", "set", "--cpus", cpus).CombinedOutput()
			fmt.Print(string(out))

			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("CPU was updated from %d to %s\n", machines[0].Resources.CPUs, cpus)
		}

		// Update Memory
		memMiB, err := utils.ConvertToMiB(memory)
		if isParamChanged(memMiB, strconv.Itoa(machines[0].Resources.Memory)) {
			out, err := exec.Command("podman", "machine", "set", "--memory", memMiB).CombinedOutput()
			fmt.Print(string(out))

			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Memory was updated from %d to %s\n", machines[0].Resources.Memory, memory)
		}

		// Update Disk Size
		if isParamChanged(diskSize, strconv.Itoa(machines[0].Resources.DiskSize)) {
			// TODO: new size must be greater than current size
			out, err := exec.Command(
				"podman",
				"machine",
				"set",
				"--disk-size", diskSize,
			).CombinedOutput()
			fmt.Print(string(out))

			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Disk size was updated from %d to %s\n", machines[0].Resources.DiskSize, diskSize)
		}

		if machines[0].State == "running" {
			err = StartCmd.RunE(StartCmd, args)
			if err != nil {
				return err
			}
		}

		err = ListCmd.RunE(ListCmd, args)
		if err != nil {
			return err
		}

		return nil
	} else {
		fmt.Println("No changes detected in configuration.")
	}
	return nil
}

var cpus, memory, diskSize string

func init() {
	ConfigCmd.Flags().StringVarP(&cpus, "cpus", "c", "0", "Number of CPUs to allocate to the podman machine")
	ConfigCmd.Flags().StringVarP(&memory, "memory", "m", "0", "Memory in GiB or in MiB to allocate to the podman machine")
	ConfigCmd.Flags().StringVarP(&diskSize, "disk-size", "d", "0", "Disk size for the podman machine")
}
