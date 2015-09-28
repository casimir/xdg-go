package xdg

import "os"

func defaultDataHome() string {
	return os.Getenv("HOME") + "/.local/share"
}

func defaultConfigHome() string {
	return os.Getenv("HOME") + "/.config"
}

func defaultCacheHome() string {
	return os.Getenv("HOME") + "/.cache"
}

func defaultDataDirs() []string {
	// The specification gives a default value with trailing slashes but only
	// for this value. Seems odd enough to take the liberty of removing them.
	return []string{"/usr/local/share", "/usr/share"}
}

func defaultConfigDirs() []string {
	return []string{"/etc/xdg"}
}
