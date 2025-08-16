package webhook

import (
	"fmt"
	"infra-lab-cli/src/utils"
)

func startWebhook(binaryName string, webhook Webhook) (err error) {
	var envs []string
	args := fmt.Sprintf("-hooks %s -template -ip %s -port %d -urlprefix %s -hotreload -logfile /dev/stdout",
		webhook.WebhooksPath,
		webhook.ListenAddr,
		webhook.ListenPort,
		webhook.UrlPrefix,
	)

	if webhook.ExtraArgs != "" {
		args = fmt.Sprintf("%s %s", args, webhook.ExtraArgs)
	}

	if webhook.Secret != "" {
		envs = []string{fmt.Sprintf("WEBHOOK_SECRET=%s", webhook.Secret)}
	}

	_, _, err = utils.ExecBinaryCommand(
		binaryName,
		args,
		true,
		false,
		envs,
	)

	return err
}

func StartWebhook(binaryName string, webhook Webhook) (err error) {
	if !utils.IsBinaryInPath(binaryName) {
		fmt.Print(utils.BinaryNotFoundError(binaryName))
		return nil
	}

	err = startWebhook(binaryName, webhook)
	if err != nil {
		return err
	}

	return nil
}
