package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Server	string		`toml:"server"`
	TopicContext string		`toml:"context"`
	Topics	[]string	`toml:"topics"`
	Qos		int			`toml:"qos"`
	ClientId string		`toml:"clientid"`
	Username string		`toml:"username"`
	Password string		`toml:"password"`
}

var Conf tomlConfig

func init() {
	if _, err := toml.DecodeFile("publisher.toml", &Conf); err != nil {
		log.Fatalln("Error decode publisher.toml... ", err)
		return
	}
}
