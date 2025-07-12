package utils

import (
	"github.com/go-ini/ini"
	"log"
)


type ConfigList struct {
	Port     			string
	Serpapi 			string
}

var Config ConfigList

func LoadConfig(path string) ConfigList {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}
	
	Config := ConfigList{
		Port:     			cfg.Section("api").Key("port").MustString(":3000"),
		Serpapi: cfg.Section("serpapi").Key("key").String(),
	}
	

	return Config
}
