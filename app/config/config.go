package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const DefaultConfig string = "./config/params-local.json"

type MysqlConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"db"`
}

type ParamsLocal struct {
	SecretKey string      `json:"secret_key"`
	ApiKey    string      `json:"api_key"`
	Mysql     MysqlConfig `json:"mysql"`
}

func NewConfig() ParamsLocal {
	config := ParamsLocal{}

	file, err := os.Open(DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Unn", err)
	}

	return config
}
