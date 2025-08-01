package webhook

type Webhook struct {
	ListenAddr   string
	ListenPort   int
	WebhooksPath string
	Secret       string
	UrlPrefix    string
	ExtraArgs    string
}
