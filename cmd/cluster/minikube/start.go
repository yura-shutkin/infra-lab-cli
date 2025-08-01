package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var StartClusterCmd = &cobra.Command{
	Use:   "start",
	Short: "Start cluster if exist",
	RunE:  runStartCluster,
}

func runStartCluster(cmd *cobra.Command, args []string) error {
	return mksrc.StartCluster(binaryName, cluster)
}
