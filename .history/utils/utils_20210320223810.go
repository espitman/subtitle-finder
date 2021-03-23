package utils

import (
	"fmt"
	"os"

	"github.com/artdarek/go-unzip"
)

func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
}

func Unzip() {
	uz := unzip.New("file.zip", "directory/")
	err := uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}
