package utils

import (
	"dgw/downtime/config"
	"fmt"
	"strings"
)

const (
	tagPrefix = "#"
	tagSuffix = "#"
)

func BuildTag(segments ...string) string {
	var sb strings.Builder
	sb.WriteString(tagPrefix)
	sb.WriteString(config.RootTag)
	for _, segment := range segments {
		sb.WriteString(fmt.Sprintf("/%s", segment))
	}
	sb.WriteString(tagSuffix)
	return sb.String()
}

func DefaultTag() string {
	return BuildTag(config.DefaultTag)
}
