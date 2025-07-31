package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var UnpauseClusterCmd = &cobra.Command{
	Use:   "unpause",
	Short: "Unpause cluster if exist",
	RunE:  runUnpauseCluster,
}

func runUnpauseCluster(cmd *cobra.Command, args []string) error {
	return mksrc.UnpauseCluster(binaryName, cluster)
}
