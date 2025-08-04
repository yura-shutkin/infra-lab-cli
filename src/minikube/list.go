package minikube

import (
	"fmt"
	"infra-lab-cli/src/common"
)

// TODO: different style. Yet easier to implement

func ListProfiles(binaryName string) (err error) {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = common.ExecBinaryCommand(
		binaryName,
		"profile list",
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
