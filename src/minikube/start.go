package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

func startCluster(binaryName string, cluster Cluster) (err error) {
	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("-p %s start", cluster.Name),
		true,
		false,
		[]string{},
	)
	if err != nil {
		return err
	}

	return nil
}

func StartCluster(binaryName string, cluster Cluster) error {
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
		fmt.Printf("Cluster %s already exists.\n", cluster.Name)
		// TODO: not sure which else statuses are possible and this is for sure a hardcode
		if existingCluster.Status == "Stopped" {
			fmt.Printf("Cluster %s is stopped. Trying to start it\n", cluster.Name)
			err = startCluster(binaryName, cluster)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Cluster %s is not stopped. Do nothing\n", cluster.Name)
		}
	} else {
		fmt.Printf("Cluster %s does not exist.\n", cluster.Name)
	}

	return nil
}
