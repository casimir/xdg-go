package xdg

import "os"

var home = os.Getenv("HOME")

func defaultDataHome() string {
	return home + "/Library"
}

func defaultConfigHome() string {
	return home + "/Library/Preferences"
}

func defaultCacheHome() string {
	return home + "/Library/Caches"
}

func defaultDataDirs() []string {
	return []string{"/Library"}
}

func defaultConfigDirs() []string {
	return []string{"/Library/Preferences", "/Library/Application Support"}
}
