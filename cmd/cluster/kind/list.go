package kind

import (
	"github.com/spf13/cobra"
	kindsrc "infra-lab-cli/src/kind"
)

var ListClustersCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List clusters",
	RunE:    runListClusters,
}

func runListClusters(cmd *cobra.Command, args []string) error {
	return kindsrc.ListClusters(binaryName)
}
