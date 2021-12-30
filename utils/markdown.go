package utils

import (
	"dgw/downtime/config"
	"fmt"
)

func H2(title string) string {
	return fmt.Sprintf("%s %s\n", config.H2Prefix, title)
}

func H3(title string) string {
	return fmt.Sprintf("%s %s\n", config.H3Prefix, title)
}

func KeyValue(key string, value string) string {
	return fmt.Sprintf("**%s**: %s\n", key, value)
}
