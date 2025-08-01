package minikube

import (
	"github.com/spf13/cobra"
	mksrc "infra-lab-cli/src/minikube"
)

var ListProfilesCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List profiles",
	RunE:    runListProfiles,
}

func runListProfiles(cmd *cobra.Command, args []string) error {
	return mksrc.ListProfiles(binaryName)
}
