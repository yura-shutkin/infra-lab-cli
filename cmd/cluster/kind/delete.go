package kind

import (
	"github.com/spf13/cobra"
	kindsrc "infra-lab-cli/src/kind"
)

var DeleteClustersCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete cluster",
	RunE:    runDeleteClusters,
}

func runDeleteClusters(cmd *cobra.Command, args []string) error {
	return kindsrc.DeleteCluster(binaryName, cluster)
}
