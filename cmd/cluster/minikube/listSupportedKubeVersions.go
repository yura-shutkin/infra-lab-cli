package minikube

import (
	mksrc "infra-lab-cli/src/minikube"

	"github.com/spf13/cobra"
)

var ListSupportedVersionsCmd = &cobra.Command{
	Use:   "list-k8s-versions",
	Short: "List supported k8s version",
	RunE:  runListSupportedVersions,
}

func runListSupportedVersions(cmd *cobra.Command, args []string) error {
	return mksrc.ListSupportedKubeVersions(binaryName)
}
