package kind

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func deleteCluster(binaryName, clusterName string) (err error) {
	_, _, err = common.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("delete cluster --name %s", clusterName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCluster(binaryName string, cluster Cluster) (err error) {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	clusters, err := getClusters(binaryName)
	if err != nil {
		return err
	}

	if common.IfStringInSlice(cluster.Name, clusters) {
		err = deleteCluster(binaryName, cluster.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
