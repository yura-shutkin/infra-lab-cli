package podman

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/src/common"
	"strings"
)

func InspectMachine(binaryName string, machineName string) (machine *InspectedMachine, err error) {
	stdout, _, err := common.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("machine inspect %s", machineName),
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	var machines []InspectedMachine
	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &machines)
	if err != nil {
		return nil, err
	}

	if len(machines) == 0 {
		return nil, fmt.Errorf("machine %s not found", machineName)
	}

	return &machines[0], nil
}
