package kind

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
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

func DeleteCluster(cluster config.Kind) (err error) {
	if !utils.IsBinaryInPath(cluster.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cluster.Binary))
		return nil
	}

	clusters, err := getClusters(cluster.Binary)
	if err != nil {
		return err
	}

	if utils.IfStringInSlice(cluster.ClusterName, clusters) {
		err = deleteCluster(cluster.Binary, cluster.ClusterName)
		if err != nil {
			return err
		}
	}

	return nil
}
