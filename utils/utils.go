package utils

import (
	"fmt"
	"regexp"
)

var excludedFolders = regexp.MustCompile("^.*/(bin|obj)(/.*)?$")

// IsTreatableFolder checks if a folder should be treated
func IsTreatableFolder(folder string) bool {
	return !excludedFolders.Match([]byte(folder))
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Errorf("error in must: %w", err))
	}
	return value
}
