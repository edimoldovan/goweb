package utilities

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExecutable() string {
	var dirAbsPath string
	dirname, err := os.Executable()
	if err == nil {
		dirAbsPath = filepath.Dir(dirname)
	}
	f := strings.Trim(dirAbsPath, "tmp")
	// fmt.Println("executable path: " + f)
	return f
}
