package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Token    string
	ClientID string
}

type config struct {
	AppBotToken string `envconfig:"APP_BOT_TOKEN"`
	ClientID    string `envconfig:"CLIENT_ID"`
}

func New() Config {
	var c config
	envconfig.Process("", &c)

	return Config{
		Token:    fmt.Sprintf("Bot %s", c.AppBotToken),
		ClientID: c.ClientID,
	}
}
