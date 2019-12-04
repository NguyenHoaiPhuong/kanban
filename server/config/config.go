package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Configurations includes all configurations for the App
type Configurations struct {
	filePath  string
	fileName string
	MGOConfig MongoConfigurations `json:"MGOConfig"`
}

// MongoConfigurations includes configurations for Mongo
type MongoConfigurations struct {
	ServerHost     string `json:"ServerHost"`
	ServerPort     string `json:"ServerPort"`
	ServerUsername string `json:"ServerUsername"`
	ServerPassword string `json:"ServerPassword"`
	DatabaseName   string `json:"DatabaseName"`
}

// SetupConfig : init configurations
func SetupConfig(filePath, fileName string) *Configurations {	
	globalConfig.filePath = filePath
	globalConfig.fileName = fileName
	globalConfig.init()

	return globalConfig
}

func (c *Configurations) init() {
	strSlice := strings.Split(c.fileName, ".")
	fileName := strSlice[0]
	ext := strSlice[1]
	// Set the file name of the configurations file
	viper.SetConfigName(fileName)
	viper.SetConfigType(ext)

	// Set the path to look for the configurations file
	viper.AddConfigPath(c.filePath)

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s\n", err)
	}

	// Set undefined variables
	viper.SetDefault("MGOConfig.ServerHost", "localhost")
	viper.SetDefault("MGOConfig.ServerPort", "27017")
	viper.SetDefault("MGOConfig.ServerUsername", "")
	viper.SetDefault("MGOConfig.ServerPassword", "")
	viper.SetDefault("MGOConfig.DatabaseName", "test")

	err := viper.Unmarshal(c)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v\n", err)
	}

	// c.Print()
}

// Print : prints all configurations on the terminal
func (c *Configurations) Print() {
	// Reading variables using the model
	fmt.Println("Reading variables using the model.")
	fmt.Println("ServerHost is\t\t", c.MGOConfig.ServerHost)
	fmt.Println("ServerPort is\t\t", c.MGOConfig.ServerPort)
	fmt.Println("ServerUsername is\t", c.MGOConfig.ServerUsername)
	fmt.Println("ServerPassword is\t", c.MGOConfig.ServerPassword)
	fmt.Println("DatabaseName is\t\t", c.MGOConfig.DatabaseName)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model.")
	fmt.Println("ServerHost is\t\t", viper.GetString("MGOConfig.ServerHost"))
	fmt.Println("ServerPort is\t\t", viper.GetString("MGOConfig.ServerPort"))
	fmt.Println("ServerUsername is\t", viper.GetString("MGOConfig.ServerUsername"))
	fmt.Println("ServerPassword is\t", viper.GetString("MGOConfig.ServerPassword"))
	fmt.Println("DatabaseName is\t\t", viper.GetString("MGOConfig.DatabaseName"))
}

// GlobalConfig init new config and return it
func GlobalConfig() *Configurations {
	return globalConfig
}

var globalConfig *Configurations

func init() {
	globalConfig = new(Configurations)
}