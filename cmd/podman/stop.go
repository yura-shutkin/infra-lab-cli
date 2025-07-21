package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"infra-lab-cli/config"
	"os/exec"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop podman machine",
	RunE:  runStop,
}

func runStop(cmd *cobra.Command, args []string) (err error) {
	if !config.IsBinaryInPath(binaryName) {
		fmt.Print(config.BinaryNotFoundError(binaryName))
		return nil
	}

	out, err := exec.Command("podman", "machine", "stop", machineName).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
