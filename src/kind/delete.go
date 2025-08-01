package kind

import (
	"fmt"
	"infra-lab-cli/utils"
)

func deleteCluster(binaryName, clusterName string) (err error) {
	_, _, err = utils.ExecBinaryCommand(
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
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	clusters, err := getClusters(binaryName)
	if err != nil {
		return err
	}

	if utils.IfStringInSlice(cluster.Name, clusters) {
		err = deleteCluster(binaryName, cluster.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
