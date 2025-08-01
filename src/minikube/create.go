package minikube

import (
	"fmt"
	"infra-lab-cli/utils"
)

func createCluster(binaryName string, cluster Cluster) (err error) {
	fmt.Printf("Creating cluster: \n")
	fmt.Printf("\tName: %s\n", cluster.Name)
	fmt.Printf("\tNodes: %d\n", cluster.NodesCount)
	fmt.Printf("\tCPUs: %s\n", cluster.Config.CPUsFlag)
	fmt.Printf("\tMemory: %s\n", cluster.Config.MemoryFlag)
	fmt.Printf("\tDiskSize: %s\n", cluster.Config.DiskSizeFlag)
	fmt.Printf("\tK8S version: %s\n", cluster.Config.KubeConfig.KubeVersion)
	fmt.Printf("\tCNI: %s\n", cluster.CNI)
	fmt.Printf("\tDriver: %s\n", cluster.Config.Driver)
	fmt.Printf("\tCIDR: %s\n", cluster.CIDR)

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		// TODO: I dislike how fragile this construction is:
		//   * It is hard to extend and read (long line)
		//   * It is possible to make a mess by just changing args order (accidentally or intentionally)
		fmt.Sprintf("-p %s start --cpus=%s --memory=%s --disk-size=%s --nodes=%d --kubernetes-version=%s --extra-config=kubeadm.pod-network-cidr=%s --driver=%s %s",
			cluster.Name,
			cluster.Config.CPUsFlag,
			cluster.Config.MemoryFlag,
			cluster.Config.DiskSizeFlag,
			cluster.NodesCount,
			cluster.Config.KubeConfig.KubeVersion,
			cluster.CIDR,
			cluster.Config.Driver,
			cluster.ExtraArgs,
		),
		true,
		false,
		[]string{},
	)

	return err
}

func CreateCluster(binaryName string, cluster Cluster) error {
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
		fmt.Printf("Cluster %s already exists. Please use recreate command instead\n", cluster.Name)
	} else {
		err = createCluster(binaryName, cluster)
		if err != nil {
			return err
		}
	}

	return nil
}
