package adaptors

import (
	"os"

	"fmt"

	"github.com/spf13/viper"
)

// GetEnvironmentVariables returns config based on your set environment
// run `ENV=DEV go run ...` on root directory
func GetEnvironmentVariables() (port string, hashString string, db map[string]string) {
	env := os.Getenv("ENV")
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config file not found")
	} else {
		switch env {
		case "PRODUCTION":
			port = viper.GetString("production.port")
			hashString = viper.GetString("production.hash_string")
			db = map[string]string{
				"host":     viper.GetString("production.db_host"),
				"name":     viper.GetString("production.db_name"),
				"user":     viper.GetString("production.db_user"),
				"password": viper.GetString("production.db_password"),
			}
		case "TEST":
			port = viper.GetString("test.port")
			hashString = viper.GetString("test.hash_string")
			db = map[string]string{
				"host":     viper.GetString("test.db_host"),
				"name":     viper.GetString("test.db_name"),
				"user":     viper.GetString("test.db_user"),
				"password": viper.GetString("test.db_password"),
			}
		default:
			port = viper.GetString("development.port")
			hashString = viper.GetString("development.hash_string")
			db = map[string]string{
				"host":     viper.GetString("development.db_host"),
				"name":     viper.GetString("development.db_name"),
				"user":     viper.GetString("development.db_user"),
				"password": viper.GetString("development.db_password"),
			}
		}
	}

	fmt.Printf("ENVIRONMENT: %s \n", env)

	return
}
