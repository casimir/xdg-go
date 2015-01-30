package xdg

import "os"

func defaultDataHome() string {
	return os.Getenv("HOME") + "/Library"
}

func defaultConfigHome() string {
	return os.Getenv("HOME") + "/Library/Preferences"
}

func defaultCacheHome() string {
	return os.Getenv("HOME") + "/Library/Caches"
}

func defaultDataDirs() []string {
	return []string{"/Library"}
}

func defaultConfigDirs() []string {
	return []string{"/Library/Preferences", "/Library/Application Support"}
}
