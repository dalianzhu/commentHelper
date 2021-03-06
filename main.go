package main

import (
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := os.Args[1]
	if !Exists(path) {
		log.Println("path is not exists")
		return
	}
	if IsDir(path) {
		err := filepath.Walk(path,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, ".pb.go") {
					log.Println(path)
					commentFile(path)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	} else if IsFile(path) {
		log.Println(path)
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, ".pb.go") {
			commentFile(path)
		}
	}
}

func commentFile(path string) {
	textBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return
	}
	text := string(textBytes)
	text = AddCommentToText(text)

	ret := wrapLine(text)
	fmtRet, err := format.Source([]byte(ret))
	if err != nil {
		log.Printf("wrap line error:%v", err)
		return
	}

	err = ioutil.WriteFile(path, fmtRet, 0666)
	if err != nil {
		log.Println(err)
		return
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
