package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var RestartClusterCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart cluster",
	RunE:  runRestartCluster,
}

func runRestartCluster(cmd *cobra.Command, args []string) error {
	return mksrc.RestartCluster(binaryName, cluster)
}
