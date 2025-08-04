package minikube

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func RecreateCluster(binaryName string, cluster Cluster) (err error) {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
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
