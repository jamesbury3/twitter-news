package configuration

import (
	"fmt"
	"just_news/models"

	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config yaml file: %w", err)
	}

	return nil
}

func Token() string {
	return viper.GetString("token")
}

func SearchUrl() string {
	return viper.GetString("searchUrl")
}

func Query() models.Query {

	required := viper.GetStringSlice("query.required")
	optional := viper.GetStringSlice("query.optional")

	return models.Query{Required: required, Optional: optional}
}
