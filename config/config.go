package config

import (
	"os"
)

type importmap struct {
	Name string
	Path string
}

type Config struct {
	PostGresConnectURL string
	BaseUrl            string
	BaseDomain         string
	Importmaps         []importmap
}

// TODO: rename to ENV and use env vars on here an nowhere else in the project
func EnvConfig() Config {
	config := Config{
		PostGresConnectURL: "some postgres connection url",
		BaseUrl:            "https://some.base.url",
		BaseDomain:         "some.base.url",
		Importmaps: []importmap{
			{
				Name: "solid-js", Path: "https://cdn.skypack.dev/solid-js",
			},
			{
				Name: "solid-js/html", Path: "https://cdn.skypack.dev/solid-js/html",
			},
			{
				Name: "solid-js/web", Path: "https://cdn.skypack.dev/solid-js/web",
			},
		},
	}
	return config
}

func IsDevelopment() bool {
	return os.Getenv("GO_WEB_ENV") == "development"
}
