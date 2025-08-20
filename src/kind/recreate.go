package kind

import (
	"fmt"
	"infra-lab-cli/src/utils"
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

	if utils.IfStringInSlice(cluster.Name, clusters) {
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
