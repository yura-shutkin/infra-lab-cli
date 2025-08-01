package podman

import (
	"fmt"
	"infra-lab-cli/utils"
)

func StartMachine(binaryName, machineName string) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("machine start %s", machineName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
