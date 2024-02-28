package helper

import (
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func MkMd5(str string) string {
	has := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", has)
}

func GetCurrentPath() string {
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	return exPath
}

func GetAppDir() string {
	appDir, err := os.Getwd()
	if err != nil {
		file, _ := exec.LookPath(os.Args[0])
		applicationPath, _ := filepath.Abs(file)
		appDir, _ = filepath.Split(applicationPath)
	}
	return appDir
}
