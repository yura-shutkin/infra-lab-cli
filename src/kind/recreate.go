package kind

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func RecreateCluster(cluster config.Kind) (err error) {
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

	err = createCluster(cluster)
	if err != nil {
		return err
	}

	return nil
}
