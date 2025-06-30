package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strconv"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure podman machine",
	RunE:  runConfig,
}

func isConfigChanged(cmd *cobra.Command, args []string) (isChanged bool, err error) {
	err = inspectCmd.RunE(inspectCmd, args)
	if err != nil {
		return isChanged, err
	}

	if strconv.Itoa(machines[0].Resources.CPUs) != cpus {
		return true, nil
	}

	if strconv.Itoa(machines[0].Resources.Memory) != memory {
		return true, nil
	}

	if strconv.Itoa(machines[0].Resources.DiskSize) != diskSize {
		return true, nil
	}

	return false, err
}

func runConfig(cmd *cobra.Command, args []string) (err error) {
	isChanged, err := isConfigChanged(cmd, args)
	if err != nil {
		return err
	}

	if isChanged {
		err = stopCmd.RunE(stopCmd, args)
		if err != nil {
			return err
		}

		// Update CPU
		if strconv.Itoa(machines[0].Resources.CPUs) != cpus {
			out, err := exec.Command(
				"podman",
				"machine",
				"set",
				"--cpus", cpus,
			).CombinedOutput()
			fmt.Print(string(out))

			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("CPU was updated from %d to %s\n", machines[0].Resources.CPUs, cpus)
		}

		// Update Memory
		if strconv.Itoa(machines[0].Resources.Memory) != memory {
			out, err := exec.Command(
				"podman",
				"machine",
				"set",
				"--memory", memory,
			).CombinedOutput()
			fmt.Print(string(out))

			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Memory was updated from %d to %s\n", machines[0].Resources.Memory, memory)
		}

		// Update Disk Size
		if strconv.Itoa(machines[0].Resources.DiskSize) != diskSize {
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

		err = startCmd.RunE(startCmd, args)
		if err != nil {
			return err
		}

		err = listCmd.RunE(listCmd, args)
		if err != nil {
			return err
		}

		return err
	} else {
		fmt.Println("No changes detected in configuration.")
	}
	return err
}

var cpus, memory, diskSize string

func init() {
	configCmd.Flags().StringVar(&cpus, "cpus", "2", "Number of CPUs to allocate to the podman machine")
	// TODO: Allow to use GB. Need to convert GB to MiB
	configCmd.Flags().StringVar(&memory, "memory", "2048", "Memory in MiB to allocate to the podman machine")
	configCmd.Flags().StringVar(&diskSize, "disk-size", "40", "Disk size for the podman machine")
}
