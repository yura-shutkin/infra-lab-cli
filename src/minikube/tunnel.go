package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

func Tunnel(binaryName string, cluster Cluster) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("-p %s tunnel", cluster.Name),
		true,
		true,
	)
	if err != nil {
		return err
	}

	return nil
}
