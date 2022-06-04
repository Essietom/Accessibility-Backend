package utilities

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort      string `mapstructure:"PORT"`
    DBHost      string `mapstructure:"DB_HOST"`
    DBName      string `mapstructure:"DB_NAME"`
    DBConnectionUri string `mapstructure:"DB_CONNECTION_URI"`
}


func LoadConfig(path string) (config Config, err error) {
    viper.AddConfigPath(path)
  //  viper.SetConfigName("testApp")
    viper.SetConfigType("env")

	viper.AutomaticEnv()

    err = viper.ReadInConfig()
    if err != nil {
        log.Print("Error reading config file, ", err)

        viper.SetDefault("DB_HOST",os.Getenv("DB_HOST"))
        viper.SetDefault("DB_PORT",os.Getenv("DB_PORT"))
        viper.SetDefault("DB_NAME",os.Getenv("DB_NAME"))
        viper.SetDefault("DB_CONNECTION_URI",os.Getenv("DB_CONNECTION_URI"))
        viper.SetDefault("PORT",os.Getenv("PORT"))
        //return
    }

    err = viper.Unmarshal(&config)
    return
}