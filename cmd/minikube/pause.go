package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var PauseClusterCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause cluster if exist",
	RunE:  runPauseCluster,
}

func runPauseCluster(cmd *cobra.Command, args []string) error {
	return mksrc.PauseCluster(binaryName, cluster)
}
