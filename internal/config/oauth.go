package config

type OauthConfig struct {
	Provider     string `mapstructure:"provider"`
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-id"`
	RedirectURL  string `mapstructure:"redirect-url"`
	TokenURL     string `mapstructure:"token-url"`
	UserinfoURL  string `mapstructure:"userinfo-url"`
}
