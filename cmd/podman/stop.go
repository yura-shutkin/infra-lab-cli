package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop podman machine",
	RunE:  runStop,
}

func runStop(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "stop", machineName).CombinedOutput()
	fmt.Print(string(out))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
