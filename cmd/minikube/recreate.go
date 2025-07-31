package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var RecreateClusterCmd = &cobra.Command{
	Use:   "recreate",
	Short: "Recreate cluster",
	RunE:  runRecreateCluster,
}

func runRecreateCluster(cmd *cobra.Command, args []string) error {
	return mksrc.RecreateCluster(binaryName, cluster)
}
