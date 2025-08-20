package vm

import (
	"infra-lab-cli/cmd/vm/podman"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "vm",
	Aliases: []string{"v"},
	Short:   "Manage VMs",
}

func init() {
	RootCmd.AddCommand(podman.RootCmd)
}
