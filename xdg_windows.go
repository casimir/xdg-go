package xdg

import "os"

func defaultDataHome() string {
	return os.Getenv("APPDATA")
}

func defaultConfigHome() string {
	return os.Getenv("APPDATA")
}

func defaultCacheHome() string {
	return os.Getenv("TEMP")
}

func defaultDataDirs() []string {
	return []string{os.Getenv("APPDATA"), os.Getenv("LOCALAPPDATA")}
}

func defaultConfigDirs() []string {
	return []string{os.Getenv("APPDATA"), os.Getenv("LOCALAPPDATA")}
}
