package kind

import (
	"fmt"
	"infra-lab-cli/src/common"
)

func createCluster(binaryName string, cluster Cluster) (err error) {
	args := fmt.Sprintf("create cluster --name %s", cluster.Name)
	if cluster.Config != "" {
		args += fmt.Sprintf(" --config=%s", cluster.Config)
	}

	_, _, err = common.ExecBinaryCommand(
		binaryName,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

func CreateCluster(binaryName string, cluster Cluster) error {
	if !common.IsBinaryInPath(binaryName) {
		fmt.Print(common.BinaryNotFoundError(binaryName))
		return nil
	}

	clusters, err := getClusters(binaryName)
	if err != nil {
		return err
	}

	if common.IfStringInSlice(cluster.Name, clusters) {
		fmt.Printf("Cluster %s already exists. Please use recreate command instead\n", cluster.Name)
	} else {
		err = createCluster(binaryName, cluster)
		if err != nil {
			return err
		}
	}

	return nil
}
