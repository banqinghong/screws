package bfile

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
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

// 下载文件
func DownloadFromRemote(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Append2File(content, filename string) {
	openType := os.O_WRONLY|os.O_APPEND
	if ! checkFileIsExist(filename) {
		openType = os.O_WRONLY|os.O_CREATE
	}
	file, err := os.OpenFile(filename, openType, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		log.Printf("write content[%s] failed: %s", content, err)
		return
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		log.Printf("flush content[%s] failed: %s", content, err)
		return
	}
}

func Write2NewFile(content, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(content)
	if err != nil {
		log.Printf("write content[%s] failed: %s", content, err)
		return
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		log.Printf("flush content[%s] failed: %s", content, err)
		return
	}
}
