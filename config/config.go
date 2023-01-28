package config

type importmap struct {
	Name string
	Path string
}

type Config struct {
	PostGresConnectURL string      `toml:"postgres_connect_url"`
	BaseUrl            string      `toml:"base_url"`
	BaseDomain         string      `toml:"base_domain"`
	Importmaps         []importmap `toml:"importmaps"`
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
