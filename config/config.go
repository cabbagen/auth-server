package config

var ApplicationConfig map[string]string = map[string]string {
	//"static": "./public",
	//"templateDir": "./views",
}

var DatabaseConfig map[string]string = map[string]string {
	"dbname": "cb_auth_server",
	"username": "root",
	"password": "artART5201314??",
}

var HttpProxyConfig map[string]string = map[string]string {
	"http": "http://localhost:7001",
}
