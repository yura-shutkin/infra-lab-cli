package minikube

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func unpauseCluster(binaryName, clusterName string) (err error) {
	_, _, err = common.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("-p %s unpause", clusterName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func UnpauseCluster(binaryName string, cluster Cluster) (err error) {
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
		err = unpauseCluster(binaryName, cluster.Name)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Cluster %s does not exist.\n", cluster.Name)
	}

	return nil
}
