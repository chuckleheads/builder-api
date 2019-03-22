package config

type Config struct {
	Port      int `mapstructure:"port"`
	Datastore DBConfig
	Oauth     OauthConfig
}
