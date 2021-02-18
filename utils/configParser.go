package util

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port struct {
		Http string `json:"http"`
	} `json:"port"`
	Redis struct {
		Uri      string `json:"uri"`
		Password string `json:"password"`
	} `json:"redis"`
	MySql struct {
		Uri string `json:"uri"`
	} `json:"mysql"`
	Tcp_mode           string `json:"tcp_mode"` // tcp, tcp4, tcp6
	Tcp_max_connection int    `json:"tcp_max_connection"`
}

var instance *Config

func GetConfigInstance() *Config {
	if instance == nil {
		inst, err := instance.loadConfig()
		if err != nil {
			log.Println(err)
			return nil
		}

		instance = inst
	}

	return instance
}

func (c *Config) loadConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("configure.json")

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal(err)
	}

	return config, err
}
