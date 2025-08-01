package cluster

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/cluster/kind"
	"infra-lab-cli/cmd/cluster/minikube"
)

var RootCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"c"},
	Short:   "Manage K8S clusters",
}

func init() {
	RootCmd.AddCommand(kind.RootCmd)
	RootCmd.AddCommand(minikube.RootCmd)
}
