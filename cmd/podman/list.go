package podman

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List podman machines",
	RunE:  runList,
}

func runList(cmd *cobra.Command, args []string) (err error) {
	out, err := exec.Command("podman", "machine", "list").CombinedOutput()
	fmt.Print(string(out))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
