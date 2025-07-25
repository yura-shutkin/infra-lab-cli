package podman

import (
	"fmt"
	"infra-lab-cli/utils"
	"os/exec"
)

func StopMachine(binaryName, machineName string) error {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	args := []string{"machine", "stop", machineName}
	out, err := exec.Command(binaryName, args...).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
