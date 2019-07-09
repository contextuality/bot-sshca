package libca

import (
	"os/user"
	"path/filepath"
	"strings"
)

// Expand out a path that starts with a tilde to be an absolute path
func ExpandPathWithTilde(path string) string {
	usr, _ := user.Current()
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(usr.HomeDir, path[2:])
	}
	return path
}