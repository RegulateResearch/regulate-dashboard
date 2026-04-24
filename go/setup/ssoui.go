package setup

import (
	"frascati/comp/ssoui"
	"frascati/config"
)

func setupSsoUiClient() ssoui.Client {
	url := config.GetSsoUiUrl()
	client := ssoui.NewSsoClient(url)
	return client
}
