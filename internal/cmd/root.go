package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/chuckleheads/builder-api/internal/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "originsrv",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./secrets/builder-api.toml)")
	viper.SetDefault("port", 7001)
	viper.SetDefault("datastore.host", "localhost")
	viper.SetDefault("datastore.port", 5432)
	viper.SetDefault("datastore.database", "builder")
	viper.SetDefault("datastore.username", "elliott.davis")
	viper.SetDefault("datastore.password", "")
	viper.SetDefault("datastore.ssl-mode", "disable")
	viper.SetDefault("oauth.provider", "github")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		pwd, _ := os.Getwd()
		home := filepath.Dir(filepath.Join(pwd, "secrets/builder-api.toml"))

		// Search config in home directory with name ".config" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// ConfigFromViper fetches database config from viper
func ConfigFromViper() (*config.Config, error) {
	cfg := &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(err.Error())
	}
	return cfg, nil
}
