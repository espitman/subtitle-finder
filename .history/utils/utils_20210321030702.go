package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/artdarek/go-unzip"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
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
	// input, _ := ioutil.ReadFile(path)
	// output := make([]byte, len(input))

	// converter, _ := iconv.NewConverter(sourceEncoding, "utf-8")
	// output, _ := converter.ConvertString(".���� � � Ӂ����� ��� 1945")
	// fmt.Println(output)
	// d1 := []byte(output)
	// _ = ioutil.WriteFile(outpath, d1, 0644)
	// // converter can then be closed explicitly
	// // this will also happen when garbage collected
	// converter.Close()
	// converter, _ := iconv.NewConverter(sourceEncoding, "utf-8")
	// _, _, _ = converter.Convert(input, output)
	// fmt.Println(string(output))
	// _ = ioutil.WriteFile(outpath, output, 0644)

	charset := chilkat.NewCharset()

	charset.SetFromCharset("utf-8")
	charset.SetToCharset("ANSI")

	// We could alternatively be more specific and say "Windows-1252".
	// The term "ANSI" means -- whatever character encoding is defined as the ANSI
	// encoding for the computer.  In Poland, for example, it would be the single-byte-per-char
	// used to represnt Eastern European language chars, which is Windows-1250.
	dec := charmap.Windows1250.NewDecoder()
	out, _ := dec.Bytes(enc)
	return string(out)
}
