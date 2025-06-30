package podman

import (
	"github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart podman machine",
	RunE:  runRestart,
}

func runRestart(cmd *cobra.Command, args []string) (err error) {
	// Execute StopCmd
	err = StopCmd.RunE(StopCmd, args)
	if err != nil {
		return err
	}
	// Execute StartCmd
	err = StartCmd.RunE(StartCmd, args)
	if err != nil {
		return err
	}
	return err
}
