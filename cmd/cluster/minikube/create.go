package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var CreateClusterCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new cluster or recreate existing by using --recreate",
	Aliases: []string{"c"},
	RunE:    runCreateCluster,
}

var recreate bool

func runCreateCluster(cmd *cobra.Command, args []string) error {
	if recreate {
		return mksrc.RecreateCluster(binaryName, cluster)
	} else {
		return mksrc.CreateCluster(binaryName, cluster)
	}
}

func init() {
	// TODO: somehow I could not figure out how to use `cluster.KubeVersions, _ = mksrc.GetSupportedKubeVersions(binaryName)` because it adds a list to completion as plain text
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.KubeConfig.KubeVersion, "kubeVersion", "v", "v1.30.0", "Version of Kubernetes cluster")
	_ = CreateClusterCmd.RegisterFlagCompletionFunc("kubeVersion", getSupportedKubeVersions)

	// TODO: Should the default value be gathered from config in HOME dir?
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.Driver, "driver", "", "podman", "Which driver to use")
	_ = CreateClusterCmd.RegisterFlagCompletionFunc("driver", getSupportedDrivers)

	CreateClusterCmd.Flags().StringVarP(&cluster.CNI, "cni", "", "auto", "The CNI binary to use")
	_ = CreateClusterCmd.RegisterFlagCompletionFunc("cni", getSupportedCNIs)

	CreateClusterCmd.Flags().StringVarP(&cluster.Config.CPUsFlag, "cpus", "c", "2", "The amount of CPUs to use")
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.MemoryFlag, "memory", "m", "2G", "The amount of memory to use")
	CreateClusterCmd.Flags().StringVarP(&cluster.Config.DiskSizeFlag, "disk-size", "d", "10G", "Disk size allocated to the minikube VM")
	CreateClusterCmd.Flags().StringVarP(&cluster.CIDR, "cidr", "", "172.16.0.0/16", "The CIDR to use")
	CreateClusterCmd.Flags().IntVarP(&cluster.NodesCount, "nodes", "n", 1, "Number of nodes")
	CreateClusterCmd.Flags().BoolVarP(&recreate, "recreate", "", false, "Recreate existing cluster")
}
