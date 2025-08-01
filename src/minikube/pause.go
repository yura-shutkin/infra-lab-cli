package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

func pauseCluster(binaryName, clusterName string) (err error) {
	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("-p %s pause", clusterName),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func PauseCluster(binaryName string, cluster Cluster) (err error) {
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
		err = pauseCluster(binaryName, cluster.Name)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("Cluster %s does not exist.\n", cluster.Name)
	}

	return nil
}
