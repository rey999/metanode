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
	file, err := os.Open("config.json")
	if err != nil {
		panic("无法打开 config.json: " + err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ConfigData); err != nil {
		panic("解析 config.json 出错: " + err.Error())
	}
}
