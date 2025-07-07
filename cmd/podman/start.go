package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start podman machine",
	RunE:  runStart,
}

func runStart(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "start", machineName).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
