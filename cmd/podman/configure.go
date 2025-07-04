package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strconv"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure podman machine",
	RunE:  runConfig,
}

func convertToMiB(size string) (string, error) {
	// Could be 2048, 2048M, 2048m, 2G, 2.5G, 2g
	result, err := strconv.Atoi(size)
	if err != nil {
		coefficient := 1.0
		if strings.Contains(strings.ToLower(size), "g") {
			coefficient = 1024.0
			size = strings.Replace(strings.ToLower(size), "g", "", 1)
		}
		if strings.Contains(strings.ToLower(size), "m") {
			size = strings.Replace(strings.ToLower(size), "m", "", 1)
		}
		var sizeFloat float64
		sizeFloat, err = strconv.ParseFloat(size, 32)
		if err != nil {
			return "0", err
		}
		return strconv.Itoa(int(sizeFloat * coefficient)), err
	}
	return strconv.Itoa(result), err
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
	err = inspectCmd.RunE(inspectCmd, args)
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
	// fmt.Printf("%#v\n", cmd.LocalNonPersistentFlags())
	if cpus == "0" && memory == "0" && diskSize == "0" {
		return cmd.Help()
	}
	isChanged, err := isConfigChanged(cmd, args)
	if err != nil {
		return err
	}

	if isChanged {
		if machines[0].State == "running" {
			err = stopCmd.RunE(stopCmd, args)
			if err != nil {
				return err
			}
		}

		// Update CPU
		if strconv.Itoa(machines[0].Resources.CPUs) != cpus && cpus != "0" {
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
		memMiB, err := convertToMiB(memory)
		if memMiB != strconv.Itoa(machines[0].Resources.Memory) && memMiB != "0" {
			out, err := exec.Command(
				"podman",
				"machine",
				"set",
				"--memory", memMiB,
			).CombinedOutput()
			fmt.Print(string(out))

			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Printf("Memory was updated from %d to %s\n", machines[0].Resources.Memory, memory)
		}

		// Update Disk Size
		if strconv.Itoa(machines[0].Resources.DiskSize) != diskSize && diskSize != "0" {
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
			err = startCmd.RunE(startCmd, args)
			if err != nil {
				return err
			}
		}

		err = ListCmd.RunE(ListCmd, args)
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
	configCmd.Flags().StringVarP(&cpus, "cpus", "c", "0", "Number of CPUs to allocate to the podman machine")
	configCmd.Flags().StringVarP(&memory, "memory", "m", "0", "Memory in GiB or in MiB to allocate to the podman machine")
	configCmd.Flags().StringVarP(&diskSize, "disk-size", "d", "0", "Disk size for the podman machine")
}
