package main

import (
	"log"
	"github.com/spf13/viper"
)

// define a struct that will contain the allowed configurations. 
// In our case, it is the port number and the MySQL connection string.
type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}
// define the AppConfig variable that will be accessed by other files and packages within the application code.
var AppConfig *Config
// use Viper to load configurations from the config.json file  
// and assign its values to the AppConfig variable.
// it will call the LoadAppConfig function from the main program which in turn will be loading the data from the JSON file into the AppConfig variable. Pretty neat, 
func LoadAppConfig(){
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")//tells Viper that our configuration file is named config.
	viper.SetConfigType("json")//Tells Viper that our config file is of JSON type.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}