package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

func RecreateCluster(binaryName string, cluster Cluster) (err error) {
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

	err = createCluster(binaryName, cluster)
	if err != nil {
		return err
	}

	return nil
}
