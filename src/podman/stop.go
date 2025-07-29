package podman

import (
	"fmt"
	"infra-lab-cli/utils"
)

func StopMachine(binaryName, machineName string) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("machine stop %s", machineName),
		true,
	)
	if err != nil {
		return err
	}

	return nil
}
