package config

import (
	"os"
)

var Config ConfigT
const (
	DB_CONNECTION_STRING_CONST string = "mylink:123@/mylink?charset=utf8"
	SHORT_LINK_LEN_CONST = 9
)

type ConfigT struct {
	DB_CONNECTION_STRING string
	SHORT_LINK_LEN int
}

func init() {
	mode := os.Getenv("MARTINI_ENV")

	switch mode {
	case "production":
		{}
	default:
		{}
	}

}
