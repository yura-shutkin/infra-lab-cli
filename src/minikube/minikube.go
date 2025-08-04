package minikube

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/src/common"
	"strings"
)

// TODO: these functions look very similar and duplicated

func GetSupportedKubeVersions(binaryName string) (versions []string, err error) {
	stdout, _, err := common.ExecBinaryCommand(
		binaryName,
		"config defaults kubernetes-version -o json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list of supported k8s versions: %v", err)
	}

	return versions, nil
}

func ListSupportedKubeVersions(binaryName string) (err error) {
	versions, err := GetSupportedKubeVersions(binaryName)
	if err != nil {
		return err
	}
	for _, version := range versions {
		fmt.Println(version)
	}
	return nil
}

func GetSupportedDrivers(binaryName string) (versions []string, err error) {
	stdout, _, err := common.ExecBinaryCommand(
		binaryName,
		"config defaults driver -o json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list of supported k8s versions: %v", err)
	}

	return versions, nil
}

func getClusters(binaryName string) (clusters []Cluster, err error) {
	stdout, _, err := common.ExecBinaryCommand(
		binaryName,
		"profile list -o json",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	var mkList MkList
	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &mkList)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}
	clusters = mkList.Valid

	return clusters, nil
}

func getClusterIfExists(newCluster Cluster, clusters []Cluster) *Cluster {
	for _, cluster := range clusters {
		if cluster.Name == newCluster.Name {
			return &cluster
		}
	}
	return nil
}
