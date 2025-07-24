package podman

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func InspectMachine(binaryName string, machineName string) (machine *InspectedMachine, err error) {
	out, err := exec.Command(binaryName, "machine", "inspect", machineName).CombinedOutput()
	if err != nil {
		return nil, err
	}

	var machines []InspectedMachine
	err = json.Unmarshal(out, &machines)
	if err != nil {
		return nil, err
	}

	if len(machines) == 0 {
		return nil, fmt.Errorf("machine %s not found", machineName)
	}

	return &machines[0], nil
}
