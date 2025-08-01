package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

// TODO: different style. Yet easier to implement

func ListProfiles(binaryName string) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
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
