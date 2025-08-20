package kind

import (
	"infra-lab-cli/src/utils"
)

func getClusters(binaryName string) (stdout []string, err error) {
	stdout, _, err = utils.ExecBinaryCommand(
		binaryName,
		"get clusters",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
