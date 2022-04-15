package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetUserCacheDir() string {
	userCahceDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println(Error("Failed to get user cache dir"))
		return ""
	}

	return userCahceDir
}

func GetStarliCacheDir() string {
	return filepath.Join(GetUserCacheDir(), "starli")
}

func GetStarliSpecsCacheDir() string {
	return filepath.Join(GetStarliCacheDir(), "specs")
}

func GetStarliSpecsEtagFile() string {
	return filepath.Join(GetStarliCacheDir(), "specs.etag")
}
