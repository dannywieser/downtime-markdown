package utils

import (
	"dgw/downtime/config"
	"fmt"
)

const (
	tagPrefix = "#"
	tagSuffix = "#"
)

func BuildTag(tagType string, tagValue string) string {
	return fmt.Sprintf("%s%s/%s/%s#", tagPrefix, config.RootTag, tagType, tagValue)
}
