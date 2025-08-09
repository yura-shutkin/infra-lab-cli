package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "minikube",
	Aliases: []string{"mk"},
	Short:   "Manage minikube clusters",
}

var cluster mksrc.Cluster
var binaryName = "minikube"

func getSupportedKubeVersions(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	kubeVersions, err := mksrc.GetSupportedKubeVersions(binaryName)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	return kubeVersions, cobra.ShellCompDirectiveNoFileComp
}

func getSupportedDrivers(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	supportedDrivers, err := mksrc.GetSupportedDrivers(binaryName)
	if err != nil {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}
	return supportedDrivers, cobra.ShellCompDirectiveNoFileComp
}

func getSupportedCNIs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return []string{"auto", "bridge", "calico", "cilium", "flannel", "kindnet"}, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	RootCmd.Flags().StringVarP(&cluster.Name, "cluster", "c", "local", "Name of the cluster to use")
	RootCmd.PersistentFlags().StringVarP(&binaryName, "binary", "b", binaryName, "Binary to use")
	RootCmd.PersistentFlags().StringVarP(&cluster.ExtraArgs, "extra-args", "", "", "Extra args to pass to minikube")

	RootCmd.AddCommand(StartClusterCmd)
	RootCmd.AddCommand(StopClusterCmd)
	RootCmd.AddCommand(RestartClusterCmd)
	RootCmd.AddCommand(CreateClusterCmd)
	RootCmd.AddCommand(DeleteClusterCmd)
	RootCmd.AddCommand(ListProfilesCmd)
	RootCmd.AddCommand(ListSupportedVersionsCmd)
	RootCmd.AddCommand(PauseClusterCmd)
	RootCmd.AddCommand(UnpauseClusterCmd)
	RootCmd.AddCommand(TunnelCmd)
}
