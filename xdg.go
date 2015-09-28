package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

// DataHome returns the base directory where user-specific data should be
// written.
func DataHome() string {
	path := os.Getenv("XDG_DATA_HOME")
	if len(path) > 0 {
		return path
	}
	return defaultDataHome()
}

// ConfigHome returns the base directory where user-specific configuration
// should be written.
func ConfigHome() string {
	path := os.Getenv("XDG_CONFIG_HOME")
	if len(path) > 0 {
		return path
	}
	return defaultConfigHome()
}

// CacheHome returns the base directory where user-specific (non-essential) data
// should be written.
func CacheHome() string {
	path := os.Getenv("XDG_CACHE_HOME")
	if len(path) > 0 {
		return path
	}
	return defaultCacheHome()
}

// DataDirs returns the ordered list of base directories where application data
// should be searched.
func DataDirs() []string {
	dirs := os.Getenv("XDG_DATA_DIRS")
	if len(dirs) > 0 {
		return strings.Split(dirs, ":")
	}
	return defaultDataDirs()
}

// ConfigDirs returns the ordered list of base directories where configuration
// should be searched.
func ConfigDirs() []string {
	dirs := os.Getenv("XDG_CONFIG_DIRS")
	if len(dirs) > 0 {
		return strings.Split(dirs, ":")
	}
	return defaultConfigDirs()
}

// App is an aggregation of XDG information for a named application. Useful to
// propagate app configuration.
type App struct {
	Name string
}

// DataPath determines the full path of a data file.
func (a App) DataPath(file string) string {
	return filepath.Join(DataHome(), a.Name, file)
}

// ConfigPath determines the full path of a data file.
func (a App) ConfigPath(file string) string {
	return filepath.Join(ConfigHome(), a.Name, file)
}

// CachePath determines the full path of a cached file.
func (a App) CachePath(file string) string {
	return filepath.Join(CacheHome(), a.Name, file)
}
