package podman

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"os/exec"
)

type InspectedMachine struct {
	ConfigDir struct {
		Path string `json:"Path"`
	} `json:"ConfigDir"`
	ConnectionInfo struct {
		PodmanSocket struct {
			Path string `json:"Path"`
		} `json:"PodmanSocket"`
		PodmanPipe interface{} `json:"PodmanPipe"`
	} `json:"ConnectionInfo"`
	Created   string `json:"Created"`
	LastUp    string `json:"LastUp"`
	Name      string `json:"Name"`
	Resources struct {
		CPUs     int           `json:"CPUs"`
		DiskSize int           `json:"DiskSize"`
		Memory   int           `json:"Memory"`
		USBs     []interface{} `json:"USBs"`
	} `json:"Resources"`
	SSHConfig struct {
		IdentityPath   string `json:"IdentityPath"`
		Port           int    `json:"Port"`
		RemoteUsername string `json:"RemoteUsername"`
	} `json:"SSHConfig"`
	State              string `json:"State"`
	UserModeNetworking bool   `json:"UserModeNetworking"`
	Rootful            bool   `json:"Rootful"`
	Rosetta            bool   `json:"Rosetta"`
}

var InspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect podman machine",
	RunE:  runInspect,
}

var inspectedMachines []InspectedMachine

func runInspect(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "inspect", machineName).CombinedOutput()
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, &inspectedMachines)
	if err != nil {
		return err
	}

	return nil
}
