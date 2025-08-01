package kind

import (
	"fmt"
	"infra-lab-cli/utils"
)

// TODO: different style. Yet easier to implement

func ListClusters(binaryName string) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		"get clusters",
		true,
		false,
	)
	if err != nil {
		return err
	}

	return nil
}
