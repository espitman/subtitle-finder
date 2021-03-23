package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/artdarek/go-unzip"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
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

func DetectEncoding(content []byte) (e encoding.Encoding, name string, certain bool) {
	return charset.DetermineEncoding(content, "")
}

func EncodeToUTF8(path, outpath string, sourceEncoding string) {
	fmt.Println(outpath)
	fmt.Println("000")
	// fmt.Println("Converting " + path + " from " + sourceEncoding + " to UTF-8 ...")
	_, err := os.Open(path)
	fmt.Println("111")
	if err != nil {
		fmt.Println("AAA")
		// panic(err)
		// defer f.Close()
		// reader, err := iconv.NewReader(f, sourceEncoding, "utf-8")
		// fmt.Println(reader)
		// if err != nil {
		// 	fmt.Println("BBB")
		// 	panic(err)
		// }
		// fo, err := os.Create(outpath)
		// if err != nil {
		// 	fmt.Println("CCC")
		// 	panic(err)
		// }
		// defer fo.Close()
		// fmt.Println("DDD")
		// io.Copy(fo, reader)
		// fmt.Println("EEE")
	}
	fmt.Println(err)
}
