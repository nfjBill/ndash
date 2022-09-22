package ndash

import (
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"io/ioutil"
	"os"
	"strings"
)

func FileRead(path string) ([]byte, error) {
	fsOpen, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fsOpen.Close()
	return ioutil.ReadAll(fsOpen)
	//data, err := ioutil.ReadFile(path)
	//if err != nil {
	//	fmt.Println("read file err:", err.Error())
	//}
	//
	//// 打印文件内容
	//return string(data)
}

func FileReadStr(path string) (string, error) {
	fsOpen, err := FileRead(path)

	return string(fsOpen), err
}

func FileWrite(path string, file []byte) error {
	var err error
	i := strings.LastIndex(path, "/")

	if i != -1 {
		dir := path[:i]
		err = Mkdir(dir)
		if err != nil {
			return err
		}
	}
	out, err := os.Create(path)
	defer out.Close()
	if err != nil {
		return err
	}
	_, err = out.Write(file)
	if err != nil {
		return err
	}
	return err
}

func FileUrlExt(filePath string, dot ...bool) string {
	hasDot := false

	if len(dot) == 1 {
		hasDot = dot[0]
	}

	ext := strings.LastIndex(filePath, ".")
	if ext == -1 || len(filePath) <= ext+1 {
		return ""
	}

	if hasDot {
		return "." + filePath[ext+1:]
	} else {
		return filePath[ext+1:]
	}
}

func FileType(buf []byte) (types.Type, error) {
	return filetype.Match(buf)
}

func FileNameMIME(ext string) types.Type {
	return filetype.GetType(ext)
}

func FileSha256(buf []byte) string {
	return Sha256(buf)
}

func FileRemove(path string) error {
	return os.RemoveAll(path)
}
