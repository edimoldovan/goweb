package config

import "os"

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
				Name: "flatpickr", Path: "/public/js/flatpickr.min.js",
			},
			{
				Name: "webcomponent", Path: "/public/js/web-component.js",
			},
		},
	}
	return config
}

func IsDevelopment() bool {
	return os.Getenv("GO_WEB_ENV") == "development"
}
