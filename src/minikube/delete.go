package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

func deleteCluster(binaryName, clusterName string) (err error) {
	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("-p %s delete", clusterName),
		true,
		false,
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

	existingCluster := getClusterIfExists(cluster, clusters)
	if existingCluster != nil {
		err = deleteCluster(binaryName, cluster.Name)
		if err != nil {
			return err
		}
	}

	return nil
}
