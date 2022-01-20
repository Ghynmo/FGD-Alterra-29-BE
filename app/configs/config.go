package configs

import "github.com/spf13/viper"

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	DBPort        string `mapstructure:"DB_PORT"`
	CTXTimeout    int    `mapstructure:"CTXTIMEOUT"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpired    int    `mapstructure:"JWT_EXPIRED"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
