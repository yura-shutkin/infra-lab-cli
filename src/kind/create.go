package kind

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func createCluster(binaryName string, cluster Cluster) (err error) {
	args := fmt.Sprintf("create cluster --name %s", cluster.Name)
	if cluster.ConfigPath != "" {
		args += fmt.Sprintf(" --config=%s", cluster.ConfigPath)
	}

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		args,
		true,
		false,
		[]string{},
	)

	return err
}

func CreateCluster(binaryName string, cluster Cluster) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	// TODO: check if VM is running, but which podman, docker, colima, what if we have multiple online or want to run kind only in specific env?

	clusters, err := getClusters(binaryName)
	if err != nil {
		return err
	}

	if utils.IfStringInSlice(cluster.Name, clusters) {
		fmt.Printf("Cluster %s already exists. Please use recreate command instead\n", cluster.Name)
	} else {
		err = createCluster(binaryName, cluster)
		if err != nil {
			return err
		}
	}

	return nil
}
