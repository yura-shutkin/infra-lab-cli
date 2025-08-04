package kind

import (
	"infra-lab-cli/src/common"
)

func getClusters(binaryName string) (stdout []string, err error) {
	stdout, _, err = common.ExecBinaryCommand(
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
