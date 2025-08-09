package webhook

import (
	webhooksrc "infra-lab-cli/src/webhook"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "HTTP webhook",
}

var binaryName = "webhook"
var webhook webhooksrc.Webhook

func init() {
	RootCmd.PersistentFlags().StringVarP(&binaryName, "binary", "b", binaryName, "Binary to use")
	RootCmd.PersistentFlags().StringVarP(&webhook.Secret, "secret", "s", "", "Secret to use")
	RootCmd.PersistentFlags().StringVarP(&webhook.ListenAddr, "addr", "", "127.0.0.1", "Listen addr")
	RootCmd.PersistentFlags().IntVarP(&webhook.ListenPort, "port", "", 8080, "Listen port")
	RootCmd.PersistentFlags().StringVarP(&webhook.WebhooksPath, "config", "c", "", "Path to config file")
	RootCmd.PersistentFlags().StringVarP(&webhook.UrlPrefix, "prefix", "p", "hooks", "Prefix")
	RootCmd.PersistentFlags().StringVarP(&webhook.ExtraArgs, "extra-args", "", "", "Extra args not covered by the infra-lab-cli")

	RootCmd.AddCommand(StartWebhookCmd)
}
