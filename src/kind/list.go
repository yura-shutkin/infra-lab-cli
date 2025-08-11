package kind

import (
	"fmt"
	"infra-lab-cli/src/utils"
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
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
