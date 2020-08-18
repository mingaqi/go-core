package _6_io

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// byte读 理解为字节流
func TestIoByte(t *testing.T) {
	// open file
	file, err := os.Open("./temp.txt")
	if err != nil {
		fmt.Println("occurred err:", err)
		return
	}
	defer file.Close()

	// 读取到切片
	buf := make([]byte, 128)
	for {
		n, err := file.Read(buf)
		// end of file
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("occurred err at read:", err)
		}

		fmt.Printf("读了%v个字节 \n", n)
		fmt.Println(string(buf))
	}

}

// 使用 bufio 来读取, 读取一行
func TestIoChar(t *testing.T) {
	file, err := os.Open("./temp.txt")
	if err != nil {
		fmt.Println("occurred err:", err)
		return
	}
	defer file.Close()

	// 定义buf
	reader := bufio.NewReader(file)
	for {
		// 读取一行
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("occurred err at read:", err)
			return
		}
		// 读取以表示字符结尾
		fmt.Println(line)
	}
}

// 使用ioutil读取文件所有内容
func TestIoUtil(t *testing.T) {
	ret, err := ioutil.ReadFile("./temp.txt")
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("occurred err at read:", err)
	}
	fmt.Printf("%s", ret)
}
