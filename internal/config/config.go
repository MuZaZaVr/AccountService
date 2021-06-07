package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Mongo MongoConfig
	HTTP  HTTPConfig
	JWT   JWTConfig
}

type MongoConfig struct {
	URI string

	Name string
	Host string
	Port int
	Dialect string
}

type HTTPConfig struct {
	Port int
	MaxHeaderBytes int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type JWTConfig struct {
	SigningKey string
}

func (cfg *Config) PrintConfig() {
	fmt.Printf("\nConfig: %v \n\n", cfg)

	fmt.Print("	MongoConfig: \n")
	fmt.Printf("MONGO_URI: %s\n", cfg.Mongo.URI)
	fmt.Printf("MONGO_PORT: %v\n", cfg.Mongo.Port)
	fmt.Printf("MONGO_HOST: %s\n", cfg.Mongo.Host)
	fmt.Printf("MONGO_NAME: %s\n", cfg.Mongo.Name)
	fmt.Printf("MONGO_DIALECT: %s\n", cfg.Mongo.Dialect)

	fmt.Print("\n	HTTPConfig: \n")
	fmt.Printf("HTTP_PORT: %v\n", cfg.HTTP.Port)
	fmt.Printf("HTTP_MAX_HEADER_BYTES: %v\n", cfg.HTTP.MaxHeaderBytes)
	fmt.Printf("HTTP_READ_TIMEOUT: %s\n", cfg.HTTP.ReadTimeout)
	fmt.Printf("HTTP_WRITE_TIMEOUT: %s\n", cfg.HTTP.WriteTimeout)

	fmt.Print("\n	JWTConfig: \n")
	fmt.Printf("JWT_SIGNING_KEY: %s\n", cfg.JWT.SigningKey)
}

func Init(configPath string) (*Config, error) {
	var config Config

	if err := loadFilesIntoViper(configPath); err != nil {
		return nil, err
	}

	if err := parseFiles(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func loadFilesIntoViper(configPath string) error {
	currentDirectoryPath, err := os.Getwd()
	if err != nil {
		return err
	}

	pathSplit := strings.SplitAfter(currentDirectoryPath, "account-service")[0]
	envFilePath := pathSplit + "/" + ".env"
	configPathSplit := strings.Split(configPath, "/")
	configFilePath := pathSplit + "/" + configPathSplit[0]
	configFileName := configPathSplit[1]

	loadConfigFiles(configFilePath, configFileName)
	loadEnvFile(envFilePath)

	return viper.ReadInConfig()
}

func parseFiles(cfg *Config) error {
	if err := parseConfigFiles(cfg); err != nil {
		return err
	}

	if err := parseEnvFile(cfg); err != nil {
		return err
	}

	return nil
}

func loadConfigFiles(configFilePath, configFileName string) {
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configFileName)
}

func loadEnvFile(envFilePath string) {
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error load .env file: %s", err)
	}
}

//@TODO Remove if section?
func parseConfigFiles(cfg *Config) error {

	// Parse Pg variables:
	if viper.IsSet("mongo.databaseName") {
		if err := viper.UnmarshalKey("mongo.databaseName", &cfg.Mongo.Name); err != nil {
			return err
		}
	} else {
		log.Fatal("mongo.databaseName is config file has not specified")
	}

	if viper.IsSet("mongo.databaseDialect") {
		if err := viper.UnmarshalKey("mongo.databaseDialect", &cfg.Mongo.Dialect); err != nil {
			return err
		}
	} else {
		log.Fatal("mongo.databaseDialect is config file has not specified")
	}

	// Parse Http variables:
	if viper.IsSet("http.port") {
		if err := viper.UnmarshalKey("http.port", &cfg.HTTP.Port); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.port is config file has not specified")
	}

	if viper.IsSet("http.maxHeaderBytes") {
		if err := viper.UnmarshalKey("http.maxHeaderBytes", &cfg.HTTP.MaxHeaderBytes); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.maxHeaderBytes is config file has not specified")
	}

	if viper.IsSet("http.readTimeout") {
		if err := viper.UnmarshalKey("http.readTimeout", &cfg.HTTP.ReadTimeout); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.readTimeout is config file has not specified")
	}

	if viper.IsSet("http.writeTimeout") {
		if err := viper.UnmarshalKey("http.writeTimeout", &cfg.HTTP.WriteTimeout); err != nil {
			return err
		}
	} else {
		log.Fatal("Http.writeTimeout is config file has not specified")
	}

	return nil
}

func parseEnvFile(cfg *Config) error {
	err := parseValues()
	if err != nil {
		return err
	}

	setValues(cfg)

	return nil
}

func parseValues() error {
	err := parseMongoValues()
	if err != nil {
		return err
	}

	err = parseJWTValues()
	if err != nil {
		return err
	}

	return nil
}

func setValues(cfg *Config) {
	setPgValues(cfg)
	setJWTValues(cfg)
}

func parseMongoValues() error {
	viper.SetEnvPrefix("mongo")

	if err := viper.BindEnv("host"); err != nil {
		return err
	}

	if err := viper.BindEnv("port"); err != nil {
		return err
	}

	return nil
}

func parseJWTValues() error {
	viper.SetEnvPrefix("jwt")

	if err := viper.BindEnv("signing_key"); err != nil {
		return err
	}

	return nil
}

func setPgValues(cfg *Config) {
	cfg.Mongo.Host = viper.GetString("host")
	cfg.Mongo.Port = viper.GetInt("port")
}

func setJWTValues(cfg *Config) {
	cfg.JWT.SigningKey = viper.GetString("signing_key")
}