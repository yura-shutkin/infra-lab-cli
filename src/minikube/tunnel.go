package minikube

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func Tunnel(binaryName string, cluster Cluster) (err error) {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	_, _, err = common.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("-p %s tunnel", cluster.Name),
		true,
		true,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}
