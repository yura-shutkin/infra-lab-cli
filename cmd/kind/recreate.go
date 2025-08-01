package kind

import (
	"github.com/spf13/cobra"
	kindsrc "infra-lab-cli/src/kind"
)

var RecreateClustersCmd = &cobra.Command{
	Use:     "recreate",
	Aliases: []string{"r"},
	Short:   "Recreate cluster",
	RunE:    runRecreateClusters,
}

func runRecreateClusters(cmd *cobra.Command, args []string) error {
	return kindsrc.RecreateCluster(binaryName, cluster)
}
