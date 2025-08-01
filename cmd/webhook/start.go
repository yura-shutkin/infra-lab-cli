package webhook

import (
	"github.com/spf13/cobra"
	webhooksrc "infra-lab-cli/src/webhook"
)

var StartWebhookCmd = &cobra.Command{
	Use:   "start",
	Short: "Start webhook",
	RunE:  runStartWebhook,
}

func runStartWebhook(cmd *cobra.Command, args []string) error {
	return webhooksrc.StartWebhook(binaryName, webhook)
}
