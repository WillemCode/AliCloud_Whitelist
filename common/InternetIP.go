package common // 定义一个名为 "common" 的包

import ( // 导入需要的包
	// 引入 JSON 编码/解码包，用于解析 JSON 格式的响应
	"fmt"
	"strings"
	"whitelist/log"

	// 引入 fmt 包，用于格式化输出
	"io/ioutil" // 引入 ioutil 包，用于读取 HTTP 响应的内容
	"net/http"  // 引入 net/http 包，用于发送 HTTP 请求和接收响应
)

// 处理 JSON 响应的结构体
type IPResponse struct {
	IPAddr string `json:"ip_addr"` // 将 JSON 中的 "ip_addr" 字段映射到结构体的 IPAddr 字段
}

// 创建一个获取外网 IP 的函数
func InternetIP() string {
	// 从第一个 URL 获取外网 IP
	responseClient1, errClient1 := http.Get(IpURL1) // 使用 http.Get 发送 HTTP GET 请求，获取第一个 IP 地址
	if errClient1 != nil {                          // 如果发生错误，输出错误并终止程序
		log.Error("获取公网IP地址失败, 请检查网络, %s", errClient1) // 输出错误信息
		log.Error("请按任意键退出...")
		fmt.Scanln()
		panic(errClient1) // 调用 panic 终止程序并输出错误
	}
	defer responseClient1.Body.Close()               // 确保在函数结束时关闭响应体
	body1, _ := ioutil.ReadAll(responseClient1.Body) // 读取响应体内容，存入 body1
	interNetIP1 := string(body1)                     // 将响应内容转换为字符串，存储在 interNetIP1 中

	// 从第一个 URL 获取外网 IP
	responseClient2, errClient2 := http.Get(IpURL2) // 使用 http.Get 发送 HTTP GET 请求，获取第一个 IP 地址
	if errClient2 != nil {                          // 如果发生错误，输出错误并终止程序
		log.Error("获取公网IP地址失败, 请检查网络, %s", errClient2) // 输出错误信息
		log.Error("请按任意键退出...")
		fmt.Scanln()
		panic(errClient2) // 调用 panic 终止程序并输出错误
	}
	defer responseClient2.Body.Close()               // 确保在函数结束时关闭响应体
	body2, _ := ioutil.ReadAll(responseClient2.Body) // 读取响应体内容，存入 body1
	interNetIP2 := string(body2)                     // 将响应内容转换为字符串，存储在 interNetIP1 中

	// // 从第二个 URL 获取外网 IP（该 URL 返回 JSON 格式的数据）
	// responseClient2, errClient2 := http.Get(IpURL2) // 使用 http.Get 发送 HTTP GET 请求，获取第二个 IP 地址
	// if errClient2 != nil {                          // 如果发生错误，输出错误并终止程序
	// 	log.Error("获取公网IP地址失败, 请检查网络, %s", errClient2) // 输出错误信息
	// 	log.Error("请按任意键退出...")
	// 	fmt.Scanln()
	// 	panic(errClient2) // 调用 panic 终止程序并输出错误
	// }
	// defer responseClient2.Body.Close()               // 确保在函数结束时关闭响应体
	// body2, _ := ioutil.ReadAll(responseClient2.Body) // 读取响应体内容，存入 body2

	// // 解析 JSON 响应，提取 IP 地址
	// var ipResp IPResponse                 // 定义一个 IPResponse 类型的变量，用于存储解析后的数据
	// err := json.Unmarshal(body2, &ipResp) // 将 body2 中的 JSON 数据解析到 ipResp 结构体中
	// if err != nil {                       // 如果解析出错，输出错误并终止程序
	// 	log.Error("解析JSON响应时出错") // 输出错误信息
	// 	log.Error("请按任意键退出...")
	// 	fmt.Scanln()
	// 	panic(err) // 调用 panic 终止程序并输出错误
	// }
	// interNetIP2 := ipResp.IPAddr // 获取 JSON 中的 "ip_addr" 字段并存储到 interNetIP2

	interNetIP1 = strings.TrimSpace(interNetIP1) // 清除换行符和空格
	interNetIP2 = strings.TrimSpace(interNetIP2) // 清除换行符和空格

	// 比较两个 IP 地址是否一致
	if interNetIP1 != interNetIP2 { // 如果从两个不同渠道获取的 IP 地址不一致，提示错误
		log.Error("获取到公网IP地址 %s 和 %s , 请关闭代理或检查网络环境", interNetIP1, interNetIP2) // 输出错误信息
		log.Error("请按任意键退出...")
		fmt.Scanln()
		panic("从两个渠道获取的公网IP地址不一致") // 调用 panic 终止程序并输出错误
	}
	return interNetIP1 // 返回第一个 IP 地址作为最终的外网 IP
}
