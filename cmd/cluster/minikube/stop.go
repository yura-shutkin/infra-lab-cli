package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var StopClusterCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop cluster",
	RunE:  runStopCluster,
}

func runStopCluster(cmd *cobra.Command, args []string) error {
	return mksrc.StopCluster(binaryName, cluster)
}
