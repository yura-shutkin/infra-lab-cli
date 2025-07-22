package podman

import (
	"encoding/json"
	"os/exec"
)

func InspectMachine(machineName string) ([]InspectedMachine, error) {
	out, err := exec.Command("podman", "machine", "inspect", machineName).CombinedOutput()
	if err != nil {
		return nil, err
	}

	var machines []InspectedMachine
	err = json.Unmarshal(out, &machines)
	if err != nil {
		return nil, err
	}

	return machines, nil
}
