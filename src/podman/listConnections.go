package podman

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func getConnections(binaryName string, connections *[]Connection) (err error) {
	args := []string{"system", "connection", "list", "--format", "json"}
	out, err := exec.Command(binaryName, args...).CombinedOutput()
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, &connections)
	if err != nil {
		return err
	}

	return nil
}

func GetConnections(binaryName string, connections *[]Connection) (err error) {
	err = getConnections(binaryName, connections)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func GetDefaultMachineName(connections *[]Connection, machineName *string) (err error) {
	// TODO: what if list of connections is empty?
	// TODO: probably this function need to execute another to update connections
	prefix := "-root"
	for _, conn := range *connections {
		if conn.IsDefault {
			*machineName = strings.ReplaceAll(conn.Name, prefix, "")
			break
		}
	}

	return nil
}

func GetMachineNames(connections *[]Connection) (machineNames []string, err error) {
	// TODO: what if list of connections is empty?
	// TODO: probably this function need to execute another to update connections
	postfix := "-root"
	for _, conn := range *connections {
		if strings.Contains(conn.Name, postfix) {
			machineNames = append(machineNames, strings.ReplaceAll(conn.Name, postfix, ""))
		}
	}

	return machineNames, nil
}
