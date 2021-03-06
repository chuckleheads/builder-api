package config

type Config struct {
	Port      int         `mapstructure:"port"`
	Datastore DBConfig    `mapstructure:"datastore"`
	Oauth     OauthConfig `mapstructure:"oauth"`
}
