package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List podman machines",
	RunE:  runList,
}

func runList(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "list").CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
