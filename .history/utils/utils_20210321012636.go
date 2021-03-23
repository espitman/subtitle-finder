package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/artdarek/go-unzip"
	"github.com/djimenez/iconv-go"
)

func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}
}

func Unzip(source string, dest string) {
	uz := unzip.New(source, dest)
	err := uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}

func GetDirFiles(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func MoveFile(oldLocation string, newLocation string) {
	_ = os.Rename(oldLocation, newLocation)
}

func big5ToUTF8(path, outpath string) {
	fmt.Println("Converting " + path + " from Big5 to UTF-8 ...")

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader, err := iconv.NewReader(f, "big5", "utf-8")
	if err != nil {
		panic(err)
	}

	fo, err := os.Create(outpath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	io.Copy(fo, reader)
}
