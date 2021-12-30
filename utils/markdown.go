package utils

import (
	"dgw/downtime/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func H1(title string) string {
	return fmt.Sprintf("%s %s\n", config.H1Prefix, title)
}

func H2(title string) string {
	return fmt.Sprintf("%s %s\n", config.H2Prefix, title)
}

func H3(title string) string {
	return fmt.Sprintf("%s %s\n", config.H3Prefix, title)
}

func KeyValue(key string, value string) string {
	return fmt.Sprintf("**%s**: %s\n", key, value)
}

func formatFileName(title string) string {
	filename := title
	filename = strings.Replace(title, " ", "_", -1)
	return fmt.Sprintf("%s/%s.md", config.Outdir, strings.ToLower(filename))
}

func SaveToFile(formatted string, title string) {
	err := os.WriteFile(formatFileName(title), []byte(formatted), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("  🗄 Saved output to `%s`\n", formatFileName(title))
}
