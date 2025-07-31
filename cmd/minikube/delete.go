package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var DeleteClusterCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete cluster",
	Aliases: []string{"d"},
	RunE:    runDeleteCluster,
}

func runDeleteCluster(cmd *cobra.Command, args []string) error {
	return mksrc.DeleteCluster(binaryName, cluster)
}
