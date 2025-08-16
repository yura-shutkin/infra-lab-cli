package kind

import (
	"fmt"
	"infra-lab-cli/config"
	"infra-lab-cli/src/utils"
)

func createCluster(cluster config.Kind) (err error) {
	args := fmt.Sprintf("create cluster --name %s", cluster.ClusterName)
	if cluster.ConfigPath != "" {
		args += fmt.Sprintf(" --config=%s", cluster.ConfigPath)
	}

	_, _, err = utils.ExecBinaryCommand(
		cluster.Binary,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

func CreateCluster(cluster config.Kind) (err error) {
	if !utils.IsBinaryInPath(cluster.Binary) {
		fmt.Print(utils.BinaryNotFoundError(cluster.Binary))
		return nil
	}

	// TODO: check if VM is running, but which podman, docker, colima, what if we have multiple online or want to run kind only in specific env?

	clusters, err := getClusters(cluster.Binary)
	if err != nil {
		return err
	}

	if utils.IfStringInSlice(cluster.ClusterName, clusters) {
		fmt.Printf("Cluster %s already exists. Please use recreate command instead\n", cluster.ClusterName)
	} else {
		err = createCluster(cluster)
		if err != nil {
			return err
		}
	}

	return nil
}
