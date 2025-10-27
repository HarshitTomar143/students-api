package config

import (
	"flag"
	"log"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

// Loading the yaml file values into struct fields
type HTTPServer struct {
	Address string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

//Defining the function
func MustLoad() *Config{

	//Getting the config path from the environment variables
	configPath := os.Getenv("CONFIG_PATH")

	//Checking if the config path exist or not, if not then the if condition will applicable
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		//checking whether the config path is still empty or not 
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	// Checking for the error in the config path 
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err:= cleanenv.ReadConfig(configPath, &cfg)
	
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())	
	}

	return &cfg
}
