package pkg

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	path = "config/config.yaml"
)

type Mailer struct {
	ApiMailer string `yaml:"api_mailer"`
}

type Yookassa struct {
	ID     string `yaml:"id"`
	ApiKey string `yaml:"api_key"`
}

type Database struct {
	URL string `yaml:"url"`
}

type Server struct {
	Port        string `yaml:"port"`
	Secret      string `yaml:"secret"`
	Env         string `yaml:"env"`
	Debug       string `yaml:"debug"`
	SmsApiKey   string `yaml:"sms_api_key"`
	ServiceHost string `yaml:"service_host"`
	Discount    string `yaml:"discount"`
}

type AppConfig struct {
	Database
	Server
	Mailer
	Yookassa
}

func CreateConfig() (*AppConfig, error) {
	err := godotenv.Load(".env.dist")
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content = []byte(os.ExpandEnv(string(content)))

	appConfig := AppConfig{}
	err = yaml.Unmarshal(content, &appConfig)
	if err != nil {
		return nil, err
	}

	return &appConfig, nil
}
