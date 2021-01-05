package bfile

import (
	"bufio"
	"fmt"
	"os"
)

// 使用buffer读取大文件
func ReadFile(fileName string, bufLength int) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file %s error %s\n", fileName, err.Error())
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("close file %s error %s\n", fileName, err.Error())
		}
	}()
	// 读取文件
	r := bufio.NewReader(f)
	buf := make([]byte, bufLength)
	for {
		_, err := r.Read(buf)
		if err != nil {
			fmt.Printf("read file %s error %s\n", fileName, err.Error())
			break
		}
		fmt.Println(string(buf))
	}
}

// 按行读取文件
func ReadFileLine(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file %s error %s\n", fileName, err.Error())
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Printf("close file %s error %s\n", fileName, err.Error())
		}
	}()
	// 读取文件
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		fmt.Printf("scan file %s error %s\n", fileName, err.Error())
	}
}
