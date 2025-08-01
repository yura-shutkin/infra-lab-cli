package cluster

import (
	"github.com/spf13/cobra"
	"infra-lab-cli/cmd/vm/podman"
)

var RootCmd = &cobra.Command{
	Use:     "cluster",
	Aliases: []string{"c"},
	Short:   "Manage K8S clusters",
}

func init() {
	RootCmd.AddCommand(podman.RootCmd)
}
