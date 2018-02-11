package config

import (
"os"
)

const (
	DB_CONNECTION_STRING string = "mylink:123@/mylink?charset=utf8"
)

func init() {
	mode := os.Getenv("MARTINI_ENV")

	switch mode {
	case "production":
		{}
	default:
		{}
	}
}
