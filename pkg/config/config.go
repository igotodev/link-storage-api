package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	PostgresDB PostgresDB `yaml:"postgres_db"`
	Listen     Listen     `yaml:"listen"`
}

type PostgresDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

type Listen struct {
	Port   string `yaml:"port" env-default:"8080"`
	BindIP string `yaml:"bind_ip" env-default:"0.0.0.0"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
