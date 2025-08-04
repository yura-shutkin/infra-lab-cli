package vm

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/vm/podman"
)

var RootCmd = &cobra.Command{
	Use:     "vm",
	Aliases: []string{"v"},
	Short:   "Manage VMs",
}

func init() {
	RootCmd.AddCommand(podman.RootCmd)
}
