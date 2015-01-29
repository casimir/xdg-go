package xdg

import "os"

var home = os.Getenv("HOME")

func defaultDataHome() string {
	return home + "/.local/share"
}

func defaultConfigHome() string {
	return home + "/.config"
}

func defaultCacheHome() string {
	return home + "/.cache"
}

func defaultDataDirs() []string {
	// The specification give a default value with trailing slashes but only
	// for this value. Seems odd enough to take the liberty of removing them.
	//return []string{"/usr/local/share/", "/usr/share/"}
	return []string{"/usr/local/share", "/usr/share"}
}

func defaultConfigDirs() []string {
	return []string{"/etc/xdg"}
}
