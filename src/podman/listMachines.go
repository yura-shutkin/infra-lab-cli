package podman

import (
	"encoding/json"
	"fmt"
	"infra-lab-cli/utils"
	"reflect"
	"strconv"
	"strings"
)

var machineFields = []MachineField{
	{"Name", "Machine name", 30},
	{"Running", "Running", 10},
	{"VMType", "VMType", 10},
	{"CPUs", "CPUs", 6},
	{"Memory", "Memory", 9},
	{"DiskSize", "Disk size", 9},
}

func ListMachines(binaryName string) error {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	machines, err := GetMachineList(binaryName)
	if err != nil {
		fmt.Printf("Error listing machines: %v\n", err)
		return nil
	}

	printMachines(machines)
	return nil
}

func GetMachineList(binaryName string) ([]ListedMachine, error) {
	stdout, _, err := utils.ExecBinaryCommand(
		binaryName,
		"machine list --format json --all-providers",
		false,
		false,
		[]string{},
	)
	if err != nil {
		return nil, err
	}

	var machines []ListedMachine
	data := strings.Join(stdout, "\n")
	jsonBytes := []byte(data)
	err = json.Unmarshal(jsonBytes, &machines)
	if err != nil {
		return nil, fmt.Errorf("failed to parse podman machine list: %v", err)
	}

	return machines, nil
}

func printMachines(machines []ListedMachine) {
	for _, f := range machineFields {
		fmt.Printf("%-*s", f.Width, f.Header)
	}
	fmt.Println()

	for _, m := range machines {
		val := reflect.ValueOf(m)
		for _, f := range machineFields {
			value := ""
			switch f.Name {
			case "Memory", "DiskSize":
				var sizeInt int64
				var sizeStr string

				sizeStr = fmt.Sprintf("%s", val.FieldByName(f.Name).Interface())
				sizeInt, err := strconv.ParseInt(sizeStr, 10, 64)
				if err != nil {
					value = "N/A"
				} else {
					value = utils.ByteCountIEC(sizeInt)
				}
			default:
				value = fmt.Sprintf("%v", val.FieldByName(f.Name).Interface())
			}
			fmt.Printf("%-*v", f.Width, value)
		}
		fmt.Println()
	}
}
