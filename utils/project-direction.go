package utils

import "strings"

func GetProjectDir(name string) string {
	split := strings.Split(name, "/")
	if len(split) > 1 {
		return split[len(split)-1]
	}
	return name
}
