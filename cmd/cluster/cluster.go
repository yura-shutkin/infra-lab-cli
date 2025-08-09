package cluster

import (
	"infra-lab-cli/cmd/cluster/kind"
	"infra-lab-cli/cmd/cluster/minikube"

	"github.com/spf13/cobra"
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
