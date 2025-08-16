package webhook

import (
	"infra-lab-cli/config"
	webhookSrc "infra-lab-cli/src/webhook"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "HTTP webhook",
}

var cfg config.ILCConfig
var webhook webhookSrc.Webhook

func init() {
	cfg = *config.GetConfig()

	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.Binary, "binary", "b", cfg.Apps.Webhook.Binary, "Binary to use")
	RootCmd.PersistentFlags().StringVarP(&webhook.Secret, "secret", "s", cfg.Apps.Webhook.Secret, "Secret to use")
	RootCmd.PersistentFlags().StringVarP(&webhook.ListenAddr, "addr", "", cfg.Apps.Webhook.ListenAddr, "Listen addr")
	RootCmd.PersistentFlags().IntVarP(&webhook.ListenPort, "port", "", cfg.Apps.Webhook.ListenPort, "Listen port")
	RootCmd.PersistentFlags().StringVarP(&webhook.WebhooksPath, "config", "c", cfg.Apps.Webhook.WebhooksPath, "Path to config file")
	RootCmd.PersistentFlags().StringVarP(&webhook.UrlPrefix, "prefix", "p", cfg.Apps.Webhook.Prefix, "Prefix")
	RootCmd.PersistentFlags().StringVarP(&webhook.ExtraArgs, "extra-args", "", "", "Extra args not covered by the infra-lab-cli")

	RootCmd.AddCommand(StartWebhookCmd)
}
