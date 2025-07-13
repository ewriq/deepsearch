package utils

import (
	"github.com/go-ini/ini"
	"log"
)


type ConfigList struct {
	Port     			string
	Serpapi 			string
	Database 			string
	Gemini 			string
	Prompt 			string
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
		Database: cfg.Section("db").Key("dsn").String(),
		Gemini: cfg.Section("ai").Key("gemini").String(),
		Prompt: cfg.Section("ai").Key("prompt").String(),
	}
	

	return Config
}
