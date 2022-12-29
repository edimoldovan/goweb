package config

import (
	"log"
	"main/utilities"

	"github.com/BurntSushi/toml"
)

type importmap struct {
	Name string
	Path string
}

type tomlConfig struct {
	PostGresConnectURL string      `toml:"postgres_connect_url"`
	BaseUrl            string      `toml:"base_url"`
	BaseDomain         string      `toml:"base_domain"`
	Importmaps         []importmap `toml:"importmaps"`
}

var (
	Config tomlConfig
)

func LoadConfig() {
	f := utilities.GetExecutable() + "/config.toml"

	if _, err := toml.DecodeFile(f, &Config); err != nil {
		log.Fatalln("Reading config failed", err)
	}

	// examples of config use
	// log.Println("PostGres URL:", config.Config.PostGresConnectURL)
}
