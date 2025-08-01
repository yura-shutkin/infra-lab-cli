package vm

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/cluster/kind"
	"infra-lab-cli/cmd/cluster/minikube"
)

var RootCmd = &cobra.Command{
	Use:     "vm",
	Aliases: []string{"v"},
	Short:   "Manage VMs",
}

func init() {
	RootCmd.AddCommand(minikube.RootCmd)
	RootCmd.AddCommand(kind.RootCmd)
}
