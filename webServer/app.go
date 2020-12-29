package webServer

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

func ListApp(appPath string) []string {
	var apps []string
	files, _ := ioutil.ReadDir(appPath)
	for _, f := range files {
		if f.IsDir() {
			apps = append(apps, f.Name())
		}
	}
	return apps
}

//DeleteAppVersion ...
func DeleteAppVersion(path, name, version string) bool {
	appPath := fmt.Sprintf("%s/%s/%s", path, name, version)
	err := os.Remove(appPath)
	if err != nil {
		return false
	}
	return true
}

// IsExists ...
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// AppNameValid ...
func AppNameValid(name string) bool {
	nameByte := []byte(name)

	valid := func(name string) bool {
		for _, n := range name {
			if !(unicode.IsDigit(n) || unicode.IsLetter(n) || string(n) == "_") {
				return false
			}
		}
		return true
	}

	if name != "" && len(nameByte) <= 32 && unicode.IsLetter([]rune(name)[0]) && valid(name) {
		return true
	} else {
		return false
	}
}
