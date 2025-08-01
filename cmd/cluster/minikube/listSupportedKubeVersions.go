package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var ListSupportedVersionsCmd = &cobra.Command{
	Use:   "list-k8s-versions",
	Short: "List supportted k8s version",
	RunE:  runListSupportedVersions,
}

func runListSupportedVersions(cmd *cobra.Command, args []string) error {
	return mksrc.ListSupportedKubeVersions(binaryName)
}
