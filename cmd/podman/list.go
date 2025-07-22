package podman

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/config"
	"infra-lab-cli/utils"
	"os/exec"
	"reflect"
	"strconv"
)

type listedMachine struct {
	Name               string `json:"Name"`
	Default            bool   `json:"Default"`
	Created            string `json:"Created"`
	Running            bool   `json:"Running"`
	Starting           bool   `json:"Starting"`
	LastUp             string `json:"LastUp"`
	Stream             string `json:"Stream"`
	VMType             string `json:"VMType"`
	CPUs               int    `json:"CPUs"`
	Memory             string `json:"Memory"`
	DiskSize           string `json:"DiskSize"`
	Port               int    `json:"Port"`
	RemoteUsername     string `json:"RemoteUsername"`
	IdentityPath       string `json:"IdentityPath"`
	UserModeNetworking bool   `json:"UserModeNetworking"`
}

type machineField struct {
	Name   string
	Header string
	Width  int
}

var listedMachines []listedMachine

var machineFields = []machineField{
	{"Name", "Machine name", 30},
	{"Running", "Running", 10},
	{"VMType", "VMType", 10},
	{"Default", "Default", 10},
	{"CPUs", "CPUs", 6},
	{"Memory", "Memory", 9},
	{"DiskSize", "Disk size", 9},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List podman machines",
	RunE:  runList,
}

func printMachines() {
	for _, f := range machineFields {
		fmt.Printf("%-*s", f.Width, f.Header)
	}
	fmt.Println()

	for _, m := range listedMachines {
		val := reflect.ValueOf(m)
		for _, f := range machineFields {
			value := ""
			switch f.Name {
			case "Memory", "DiskSize":
				var sizeInt int64
				var sizeStr string

				sizeStr = fmt.Sprintf("%s", val.FieldByName(f.Name).Interface())
				// TODO: handle error
				sizeInt, _ = strconv.ParseInt(sizeStr, 10, 64)
				value = utils.ByteCountIEC(sizeInt)
			default:
				value = fmt.Sprintf("%v", val.FieldByName(f.Name).Interface())
			}
			fmt.Printf("%-*v", f.Width, value)
		}
		fmt.Println()
	}
}

func listMachines() (err error) {
	out, err := exec.Command("podman", "machine", "list", "--format", "json", "--all-providers").CombinedOutput()
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, &listedMachines)
	if err != nil {
		fmt.Printf("Failed to parse podman machine list: %v\n", err)
		return err
	}

	return nil
}

func runList(cmd *cobra.Command, args []string) (err error) {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	err = listMachines()
	if err != nil {
		fmt.Printf("Error listing machines: %v\n", err)
		return nil
	}

	printMachines()
	return nil
}
