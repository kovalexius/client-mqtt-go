package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Database dbConfig `toml:"mongoDb"`
	Web webConfig		`toml:"web"`
	Mqtt mqttConfig		`toml:"mqtt"`
}

type dbConfig struct {
	DbURL string	`toml:"url"`
	DbPort string 	`toml:"port"`
	DbName string	`toml:"name"`
	DbUser string	`toml:"user"`
	DbPwd string	`toml:"password"`
	
}

type webConfig struct {
	Context string		`toml:"context"`
	Port string			`toml:"port"`
	HtmlStatic string	`toml:"html_static"`
}

type mqttConfig struct {
	Server	string		`toml:"server"`
	Topic	string		`toml:"topic"`
	Qos		int			`toml:"qos"`
	ClientId string		`toml:"clientid"`
	Username string		`toml:"username"`
	Password string		`toml:"password"`
}

var Conf tomlConfig

func init() {
	if _, err := toml.DecodeFile("subscriber.toml", &Conf); err != nil {
		log.Fatalln("Error decode subscriber.toml... ", err)
		return
	}
}
