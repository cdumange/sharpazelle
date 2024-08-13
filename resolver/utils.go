package resolver

import "regexp"

var excludedFolders = regexp.MustCompile("^.*/(bin|obj)(/.*)?$")

func isTreatableFolder(folder string) bool {
	return !excludedFolders.Match([]byte(folder))
}
