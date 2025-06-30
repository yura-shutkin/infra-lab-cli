package podman

import (
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart podman machine",
	RunE:  runRestart,
}

func runRestart(cmd *cobra.Command, args []string) (err error) {
	// Execute StopCmd
	err = stopCmd.RunE(stopCmd, args)
	if err != nil {
		return err
	}
	// Execute StartCmd
	err = startCmd.RunE(startCmd, args)
	if err != nil {
		return err
	}
	return err
}
