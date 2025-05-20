package log

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

// 定义日志级别的常量
const (
	INFO  = "[INFO]"
	WARN  = "[WARN]"
	ERROR = "[ERRO]"
)

// 定义颜色常量（使用 ANSI 转义码）
// const (
// 	Reset  = "\033[0m"
// 	Red    = "\033[31m"
// 	Yellow = "\033[33m"
// 	Green  = "\033[32m"
// )

// 获取当前时间的格式化字符串
func getTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 打印 INFO 级别的日志，支持多个参数
func Info(message string, args ...interface{}) {
	// 格式化多个参数为一个字符串
	formattedMessage := fmt.Sprintf(message, args...)
	// fmt.Printf("%s %s %s %s \n", getTimestamp(), INFO, formattedMessage)
	// color.Green.Printf("%s %s %s \n", getTimestamp(), INFO, formattedMessage)
	color.Cyan.Printf("%s %s %s \n", getTimestamp(), INFO, formattedMessage)
}

// 打印 WARN 级别的日志
func Warn(message string, args ...interface{}) {
	// 格式化多个参数为一个字符串
	formattedMessage := fmt.Sprintf(message, args...)
	// fmt.Printf("%s %s %s %s \n", Yellow, getTimestamp(), WARN, formattedMessage)
	color.Yellow.Printf("%s %s %s \n", getTimestamp(), WARN, formattedMessage)
}

// 打印 ERROR 级别的日志
func Error(message string, args ...interface{}) {
	// 格式化多个参数为一个字符串
	formattedMessage := fmt.Sprintf(message, args...)
	// fmt.Printf("%s %s %s %s \n", Red, getTimestamp(), ERROR, formattedMessage)
	color.Red.Printf("%s %s %s \n", getTimestamp(), ERROR, formattedMessage)
}
