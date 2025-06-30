package podman

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

type PMachine struct {
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

var inspectCmd = &cobra.Command{
	Use:   "start",
	Short: "Start podman machine",
	RunE:  runInspect,
}

var machines []PMachine

func runInspect(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "inspect", machineName).CombinedOutput()
	if err != nil {

		fmt.Println("Error:", err)
	}

	err = json.Unmarshal(out, &machines)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
