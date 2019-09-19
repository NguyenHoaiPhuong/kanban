package config

import (
	"github.com/paked/configure"
)

// Config includes all configurations for the App
type Config struct {
	MongoConfig
}

// MongoConfig includes configurations for Mongo
type MongoConfig struct {
	Host     *string
	Port     *string
	UserName *string
	Password *string
	DBName   *string
}

var cf *Config
var conf *configure.Configure

func init() {
	conf = configure.New()

	// Default configurations
	cf = &Config{
		MongoConfig: MongoConfig{
			Host:     conf.String("ServerHost", "localhost", "MongoDB server host"),
			Port:     conf.String("ServerPort", "27017", "MongoDB port"),
			UserName: conf.String("ServerUsername", "", "MongoDB username"),
			Password: conf.String("ServerPassword", "", "MongoDB password"),
			DBName:   conf.String("DatabaseName", "test", "MongoDB database name"),
		},
	}
}

// SetupConfig parses app 's configurations
func SetupConfig(fileName string) *Config {
	conf.Use(configure.NewFlag())
	conf.Use(configure.NewEnvironment())
	conf.Use(configure.NewJSONFromFile(fileName))
	conf.Parse()
	return cf
}
