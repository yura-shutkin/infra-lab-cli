package kind

import (
	"infra-lab-cli/utils"
)

func getClusters(binaryName string) (stdout []string, err error) {
	stdout, _, err = utils.ExecBinaryCommand(
		binaryName,
		"get clusters",
		false,
		false,
	)
	if err != nil {
		return nil, err
	}

	return stdout, nil
}
