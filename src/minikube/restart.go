package minikube

func RestartCluster(binaryName string, cluster Cluster) (err error) {
	err = StopCluster(binaryName, cluster)
	if err != nil {
		return err
	}

	err = StartCluster(binaryName, cluster)
	if err != nil {
		return err
	}

	return nil
}
