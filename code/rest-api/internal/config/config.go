package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"go.mod/pkg/logging"
)

type Config struct{
	IsDebug *bool			`yaml:"is_debug"`
	Listen struct {
		Type string			`yaml:"type"`
		BindIP string		`yaml:"bind_ip"`
		Port string			`yaml:"port"`
	}						`yaml:"Listen"`
	Storage StorageConfig 	`yaml:"storage"`
}

type StorageConfig struct{
	Host 		string `json:"host"`
	Port 		string `json:"port"`
	Database 	string `json:"database"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config{
	once.Do(func()  {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		//ЗАМЕНИТЬ "/home/serpc/projects/TP_L/Project_Leonid/codes2.0/rest-api/logs" на использование переменных 
		if err := cleanenv.ReadConfig("/home/serpc/projects/TP_L/Project_L/codes2.0/rest-api/config.yml", instance); err != nil{
			help,_ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}