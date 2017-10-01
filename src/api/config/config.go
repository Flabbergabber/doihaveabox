package config

import (
	"io/ioutil"
	"log"
	"sync"
)

type RiotApiKey string

const secretKeyFilePath string = "secret/secret_key"

type riotApiConfig struct {
	SecretKey RiotApiKey
}

func (config *riotApiConfig) GetRiotApiKey() RiotApiKey {

	if len(config.SecretKey) == 0 {
		if secretKeyBuf, err := ioutil.ReadFile(secretKeyFilePath); err != nil {
			log.Fatal("Failure reading Secret Riot API Key file.")
		} else {
			config.SecretKey = RiotApiKey(secretKeyBuf)
		}
	}

	return config.SecretKey
}

var instance *riotApiConfig
var once sync.Once

func GetInstance() *riotApiConfig {
	once.Do(func() {
		instance = &riotApiConfig{}
	})
	return instance
}