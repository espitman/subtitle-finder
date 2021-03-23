package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/artdarek/go-unzip"
	"github.com/djimenez/iconv-go"
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

func EncodeToUTF8(path string, outpath string, sourceEncoding string) {
	input, _ := ioutil.ReadFile(path)
	out := make([]byte, len(input))
	// fmt.Println(sourceEncoding)
	// converter, _ := iconv.NewConverter(sourceEncoding, "utf-8")
	// output, _ := converter.ConvertString(".���� � � Ӂ����� ��� 1945")
	// fmt.Println(output)
	// d1 := []byte(output)
	// _ = ioutil.WriteFile(outpath, d1, 0644)
	// // converter can then be closed explicitly
	// // this will also happen when garbage collected
	// converter.Close()
	converter := iconv.NewConverter("utf-8", "windows-1252")
	bytesRead, bytesWritten, error := converter.Convert(input, output)

}