package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FileRemove(logFile string) error {
	_, err := os.Stat(logFile)
	if err == nil {
		return os.Remove(logFile)
	}
	return nil
}

func CheckFile(filePath string) {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		path := filepath.Dir(filePath)

		MkDir(path)
	case os.IsPermission(err):
		log.Fatalf("permission:&v", err)
	}
}

func SaveBase64File(b64f string, suffix string) (filename string, err error) {

	// todo 如果不输入后缀,则检查文件类型
	if suffix == "" {

	}

	sf, err := base64.StdEncoding.DecodeString(b64f)
	if err != nil {
		return "", err
	}

	if len(sf) > 1024*1024*5 {
		return "", errors.New("文件大小不超过5mb")
	}

	filename = fmt.Sprintf("%d.%s", RandInt(1000, 9999), suffix)
	err = ioutil.WriteFile("public/uploads/"+filename, sf, 0666)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func MkDir(filePath string) {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+filePath, os.ModePerm)
	if err != nil {
		panic(any(err))
	}
}
