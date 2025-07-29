package podman

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/utils"
)

func InspectMachine(binaryName string, machineName string) (machine *InspectedMachine, err error) {
	var stdout []byte
	stdout, _, err = utils.ExecBinaryCommand(
		binaryName,
		fmt.Sprintf("machine inspect %s", machineName),
		false,
	)
	if err != nil {
		return nil, err
	}

	var machines []InspectedMachine
	err = json.Unmarshal(stdout, &machines)
	if err != nil {
		return nil, err
	}

	if len(machines) == 0 {
		return nil, fmt.Errorf("machine %s not found", machineName)
	}

	return &machines[0], nil
}
