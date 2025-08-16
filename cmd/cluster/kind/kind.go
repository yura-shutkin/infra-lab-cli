package kind

import (
	"infra-lab-cli/config"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "kind",
	Short: "Manage kind clusters",
}

var cfg config.ILCConfig

func init() {
	cfg = *config.GetConfig()

	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Kind.Binary, "binary", "b", cfg.Apps.Kind.Binary, "Binary to use")
	RootCmd.Flags().StringVarP(&cfg.Apps.Kind.ClusterName, "cluster", "c", cfg.Apps.Kind.ClusterName, "Name of the cluster to use")
	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Kind.ConfigPath, "config", "", cfg.Apps.Kind.ConfigPath, "Path to kind config")

	RootCmd.AddCommand(CreateClustersCmd)
	RootCmd.AddCommand(DeleteClustersCmd)
	RootCmd.AddCommand(RecreateClustersCmd)
	RootCmd.AddCommand(ListClustersCmd)
}
