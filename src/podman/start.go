package podman

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func StartMachine(binaryName, machineName string) (err error) {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = common.ExecBinaryCommand(
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
