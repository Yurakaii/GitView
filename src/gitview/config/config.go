package config

import (
	"encoding/json"
	"log"
	"os"
)


type Config struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Directory string `json:"directory"`
	RemoteURL string `json:"remoteURL"`
	ForcePush bool `json:"forcePush"`
}

const configLocation string = "./gitview.json"
func ReadConfig() Config {
	file, err := os.Open(configLocation)
	if(err != nil) {
		log.Fatalf("Failed to read config: %v", err)
	}
	decoder := json.NewDecoder(file)
	var data Config
	decoder.Decode(&data)
	return data
}