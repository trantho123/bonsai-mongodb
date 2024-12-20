package utils

import "github.com/spf13/viper"

// Config is the application configuration
type Config struct {
	DBSource       string `mapstructure:"DB_SOURCE"`
	HTTPServerPort string `mapstructure:"HTTP_SERVER_ADDRESS"`
	DBName         string `mapstructure:"DB_NAME"`
	AccessTokenKey string `mapstructure:"ACCESS_TOKEN_KEY"`
	SMTPHost       string `mapstructure:"SMTP_HOST"`
	SMTPPort       string `mapstructure:"SMTP_PORT"`
	EmailFrom      string `mapstructure:"EMAIL_FROM"`
	EmailPassword  string `mapstructure:"EMAIL_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = viper.Unmarshal(&config)
	return
}
