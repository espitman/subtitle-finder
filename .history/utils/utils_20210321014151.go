package utils

import (
	"fmt"
	"io"
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

func DetectEncoding(content []byte, contentType string) (e encoding.Encoding, name string, certain bool) {
	return charset.DetermineEncoding(content, contentType)
}

func  GetFileEncoding(string srcFile) {
	// *** Use Default of Encoding.Default (Ansi CodePage)
	Encoding enc = Encoding.Default;

	// *** Detect byte order mark if any - otherwise assume default
	byte[] buffer = new byte[5];
	FileStream file = new FileStream(srcFile, FileMode.Open);
	file.Read(buffer, 0, 5);
	file.Close();

	if (buffer[0] == 0xef && buffer[1] == 0xbb && buffer[2] == 0xbf)
		enc = Encoding.UTF8;
	else if (buffer[0] == 0xfe && buffer[1] == 0xff)
		enc = Encoding.Unicode;
	else if (buffer[0] == 0 && buffer[1] == 0 && buffer[2] == 0xfe && buffer[3] == 0xff)
		enc = Encoding.UTF32;
	else if (buffer[0] == 0x2b && buffer[1] == 0x2f && buffer[2] == 0x76)
		enc = Encoding.UTF7;
	else if (buffer[0] == 0xFE && buffer[1] == 0xFF)      
		// 1201 unicodeFFFE Unicode (Big-Endian)
		enc = Encoding.GetEncoding(1201);      
	else if (buffer[0] == 0xFF && buffer[1] == 0xFE)      
		// 1200 utf-16 Unicode
		enc = Encoding.GetEncoding(1200);


	return enc;
}

func Big5ToUTF8(path, outpath string) {
	fmt.Println("Converting " + path + " from Big5 to UTF-8 ...")

	f, err := os.Open(path)
	if err != nil {
		panic(err)
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
}
