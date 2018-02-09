package config

import (
"os"
)

const (
	SITE_NAME string = "LocTalk"
	DEFAULT_LIMIT  int = 10
	MAX_LIMIT      int = 1000
	MAX_POST_CHARS int = 1000
)
func init() {
	mode := os.Getenv("MARTINI_ENV")

	switch mode {
	case "production":
		SiteUrl = "http://mylink.ru"
		AbsolutePath = "/path/to/project/"
	default:
		SiteUrl = "http://127.0.0.1"
		AbsolutePath = "/path/to/project/"
	}
}
