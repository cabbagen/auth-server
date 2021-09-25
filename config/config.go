package config

var ApplicationConfig map[string]string = map[string]string {
	//"static": "./public",
	//"templateDir": "./views",
}

var CacheConfig map[string]string = map[string]string {
	"addr": "localhost:6379",
	"password": "",
	"db": "0",
}

var DatabaseConfig map[string]string = map[string]string {
	"dbname": "cb_auth_server",
	"username": "root",
	"password": "artART5201314??",
}

var HttpProxyConfig map[string]string = map[string]string {
	"decoration": "http://localhost:7001",
	"user-server": "http://localhost:8080",
}
