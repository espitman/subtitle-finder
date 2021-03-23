package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"unicode/utf8"

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

func EncodeToUTF8(path, outpath string, sourceEncoding string) {
	fmt.Println("Converting " + path + " from " + sourceEncoding + " to UTF-8 ...")
	f, err := os.Open(path)
	if err != nil {
		defer f.Close()
		panic(err)
	}
	reader, err := iconv.NewReader(f, sourceEncoding, "utf-8")
	if err != nil {
		panic(err)
	}
	fo, err := os.Create(outpath)
	if err != nil {
		defer fo.Close()
		panic(err)
	}
	io.Copy(fo, reader)
	fmt.Println(err)
}

func ExampleEncodeRune() {
	r := ".���� � � Ӂ����� ��� 1945"
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
	// Output:
	// [228 184 150]
	// 3
}
