package viper

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var config *viper.Viper = viper.New()

func TestViper(t *testing.T) {
	assert.NotNil(t, config)
}

func TestJson(t *testing.T) {
 	config.SetConfigType("json")
 	config.SetConfigName("config")
 	config.AddConfigPath(".")
 	err := config.ReadInConfig()
 	assert.NoError(t, err)

 	appName := config.GetString("app.name")
 	assert.Equal(t, "Viper", appName)

 	port := config.GetInt("database.port")
 	assert.Equal(t, 3306, port)

 	debug := config.GetBool("database.debug")
 	assert.Equal(t, true, debug)
}

func TestYaml(t *testing.T) {
 	//config.SetConfigType("yaml")
 	//config.SetConfigName("config")
 	config.SetConfigFile("config.yaml")
 	config.AddConfigPath(".")
 	err := config.ReadInConfig()
 	assert.NoError(t, err)

 	appName := config.GetString("app.name")
 	assert.Equal(t, "Viper", appName)

 	port := config.GetInt("database.port")
 	assert.Equal(t, 3306, port)

 	debug := config.GetBool("database.debug")
 	assert.Equal(t, true, debug)
}

func TestEnvFile(t *testing.T) {
 	config.SetConfigFile("config.env")
 	config.AddConfigPath(".")
 	err := config.ReadInConfig()
 	assert.NoError(t, err)

 	appName := config.GetString("APP_NAME")
 	assert.Equal(t, "Viper", appName)

 	port := config.GetInt("DB_PORT")
 	assert.Equal(t, 3306, port)

 	debug := config.GetBool("DB_DEBUG")
 	assert.Equal(t, true, debug)
}

func TestEnvVar(t *testing.T) {
 	config.AddConfigPath(".")
	config.AutomaticEnv() // Automatically read environment variables

 	err := config.ReadInConfig()
	assert.NoError(t, err)

	os.Setenv("APP_ENV", "production") // export APP_ENV=production
	assert.Equal(t, "production", config.GetString("APP_ENV"))
}