package service

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func InitService() {
	welcome()

	// 读取输入

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\n")

		// 处理输入命令
		handleCommand(line)
	}
}

func welcome() {
	fmt.Println("oi!")
}

func handleCommand(command string) {

	args := strings.Split(command, " ")

	// 命令
	cmd := Command{
		Name: args[0],
	}

	for _, arg := range args[1:] {
		cmd.Args = append(cmd.Args, arg)
	}

	switch cmd.Name {
	case "gsl":
		handleGenerateSliceCommand(cmd)
		// other commands
	}

}

// 有且只需要 两个 int64 参数
func handleGenerateSliceCommand(cmd Command) {

	// 判断是否满足 参数数量: 2
	if len(cmd.Args) < 2 {
		fmt.Println("❌: 参数缺少, 示例: gsl int64 int64")
		return
	}

	prompt := []int64{}

	// 是否可以转 int64
	for _, arg := range cmd.Args[:2] {
		item, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			fmt.Println("❌: 参数非法, 示例: gsl int64 int64")
			return
		}
		prompt = append(prompt, item)
	}

	slice := genNumberSlice(prompt[0], prompt[1])

	fmt.Printf("✅: %v\n", slice)
}

func genNumberSlice(start, end int64) string {

	slice := make([]int64, 0)

	if start <= end {
		for i := start; i <= end; i++ {
			slice = append(slice, i)
		}
	} else {

		for i := start; i >= end; i-- {
			slice = append(slice, i)
		}
	}

	return IntArrayToString(slice)
}

// 使用strings.Join将数组转化为字符串
func IntArrayToString(nums []int64) string {

	var sb strings.Builder

	for i, num := range nums {

		// 使用 strings.Builder 来构建字符串
		sb.WriteString(strconv.Itoa(int(num)))

		// 不包括最后一个数后面加入逗号
		if i < len(nums)-1 {
			sb.WriteString(",")
		}
	}
	// 返回结果
	return "[" + sb.String() + "]"
}
