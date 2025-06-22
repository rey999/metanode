package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SecretKey  string `json:"secretKey"`
	SqliteName string `json:"sqliteName"`
}

var ConfigData Config = Config{}

func LoadConfig() {
	file, err := os.Open("../config.json")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ConfigData); err != nil {
		return
	}
}
