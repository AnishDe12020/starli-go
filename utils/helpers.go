package utils

import "os"

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CurrentDirName() (string, error) {
	return os.Getwd()
}
