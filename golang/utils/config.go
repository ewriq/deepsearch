package utils

import (
	"log"

	"github.com/go-ini/ini"
)

type ConfigList struct {
	Port     string
	Serpapi  string
	Database string
	Gemini   string
	Prompt   string
	Google   bool
	Yandex   bool
	Bing     bool
	Alternative bool
}

var Config ConfigList

func LoadConfig(path string) ConfigList {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	Config = ConfigList{
		Port:     cfg.Section("api").Key("port").MustString(":3000"),
		Serpapi:  cfg.Section("serpapi").Key("key").String(),
		Database: cfg.Section("db").Key("dsn").String(),
		Gemini:   cfg.Section("ai").Key("gemini").String(),
		Prompt:   cfg.Section("ai").Key("prompt").String(),
		Google:   cfg.Section("serpapi").Key("google").MustBool(true),
		Yandex:   cfg.Section("serpapi").Key("yandex").MustBool(true),
		Bing:     cfg.Section("serpapi").Key("bing").MustBool(true),
		Alternative: cfg.Section("serpapi").Key("alternative").MustBool(false),
	}

	return Config
}
