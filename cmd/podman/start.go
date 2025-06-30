package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start podman machine",
	RunE:  runStart,
}

func runStart(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "start", machineName).CombinedOutput()
	fmt.Print(string(out))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
