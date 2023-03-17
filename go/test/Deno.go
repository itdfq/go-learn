package test

import (
	"strings"
	"time"
)

func Splict(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
		time.Sleep(1000)
	}
	result = append(result, s)
	return
}

// Splict2 函数接受两个字符串作为参数，返回一个字符串切片
func Splict2(input, sep string) []string {
	var result []string                // 定义一个空切片用来存储结果
	index := strings.Index(input, sep) // 查找sep在input中的位置
	for index >= 0 {                   // 如果找到了sep
		result = append(result, input[:index]) // 将input中sep之前的部分追加到结果中
		input = input[index+len(sep):]         // 将input更新为sep之后的部分
		index = strings.Index(input, sep)      // 再次查找sep在input中的位置
	}
	result = append(result, input) // 将最后剩余的input追加到结果中
	return result                  // 返回结果
}
