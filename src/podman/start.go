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

	out, err := exec.Command(binaryName, "machine", "start", machineName).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
