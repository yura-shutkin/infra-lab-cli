package podman

import (
	"fmt"
	"infra-lab-cli/config"
	"os/exec"
)

// StopMachine stops the specified podman machine.
func StopMachine(binaryName, machineName string) error {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	out, err := exec.Command("podman", "machine", "stop", machineName).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
