package kind

import (
	kindsrc "infra-lab-cli/src/kind"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kind",
	Short: "Manage kind clusters",
}

var cluster kindsrc.Cluster
var binaryName = "kind"

func init() {
	RootCmd.Flags().StringVarP(&cluster.Name, "cluster", "c", "local", "Name of the cluster to use")
	RootCmd.PersistentFlags().StringVarP(&cluster.Config, "config", "", "", "Path to kind config")
	RootCmd.PersistentFlags().StringVarP(&binaryName, "binary", "b", binaryName, "Binary to use")

	RootCmd.AddCommand(CreateClustersCmd)
	RootCmd.AddCommand(DeleteClustersCmd)
	RootCmd.AddCommand(RecreateClustersCmd)
	RootCmd.AddCommand(ListClustersCmd)
}
