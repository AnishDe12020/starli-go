package utils

import (
	"os"
	"strings"
)

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func LastElementOfSliceString(s []string) string {
	return s[len(s)-1]
}

func GetCurrentDirName() (string, error) {
	dir, err := GetCurrentDirPath()

	if err != nil {
		return "", err
	}

	dirName := LastElementOfSliceString(strings.Split(dir, "/"))

	return dirName, nil
}

func GetCurrentDirPath() (string, error) {
	return os.Getwd()
}
