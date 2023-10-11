package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 欢迎信息
	welcome()

	// 读取输入
	reader := bufio.NewReader(os.Stdin)
	for {
		// 读取一行输入
		text, _ := reader.ReadString('\n')

		// 处理输入命令
		handleCommand(text)
	}
}

func welcome() {
	fmt.Println("欢迎使用命令行程序!")
}

func handleCommand(command string) {
	// 去除换行符
	command = strings.TrimRight(command, "\n")

	// 处理 gar 命令
	if command == "gsl" {
		handleGenerateSliceCommand("123")
		return
	}

	// 其他命令处理...
}

func handleGenerateSliceCommand(num any) {
	fmt.Printf("num: %v\n", num)
}
