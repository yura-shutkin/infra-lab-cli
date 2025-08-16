package webhook

import (
	"infra-lab-cli/config"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "HTTP webhook",
}

var cfg config.ILCConfig

func init() {
	cfg = *config.GetConfig()

	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.Binary, "binary", "b", cfg.Apps.Webhook.Binary, "Binary to use")
	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.Secret, "secret", "s", cfg.Apps.Webhook.Secret, "Secret to use")
	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.ListenAddr, "addr", "", cfg.Apps.Webhook.ListenAddr, "Listen addr")
	RootCmd.PersistentFlags().IntVarP(&cfg.Apps.Webhook.ListenPort, "port", "", cfg.Apps.Webhook.ListenPort, "Listen port")
	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.WebhooksPath, "config", "c", cfg.Apps.Webhook.WebhooksPath, "Path to config file")
	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.Prefix, "prefix", "p", cfg.Apps.Webhook.Prefix, "Prefix")
	RootCmd.PersistentFlags().StringVarP(&cfg.Apps.Webhook.ExtraArgs, "extra-args", "", "", "Extra args not covered by the infra-lab-cli")

	RootCmd.AddCommand(StartWebhookCmd)
}
