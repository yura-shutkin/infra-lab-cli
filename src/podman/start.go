package podman

import (
	"fmt"
	"infra-lab-cli/utils"
	"os/exec"
)

func StartMachine(binaryName, machineName string) error {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	args := []string{"machine", "start", machineName}
	out, err := exec.Command(binaryName, args...).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
