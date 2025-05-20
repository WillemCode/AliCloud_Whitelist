package main

import (
	"fmt"              // 格式化输入输出
	"io/ioutil"        // 文件读写操作
	"os"               // 操作系统功能接口
	"strings"          // 字符串处理
	"time"             // 时间相关操作
	"whitelist/common" // 自定义公共模块
	"whitelist/config" // 自定义配置模块
	"whitelist/log"
)

// 文件存在性检查函数
func fileExists(filename string) bool {
	_, err := os.Stat(filename) // 获取文件信息
	return !os.IsNotExist(err)  // 通过错误类型判断是否存在
}

func main() {

	// 获取当前公网IP地址
	log.Info("程序开始执行")

	currentIP := common.InternetIP()
	log.Info("当前公网IP %s", currentIP) // 打印调试信息

	var oldIP string
	// 检查IP记录文件是否存在
	if fileExists(common.FileName) {
		// 读取文件内容（旧IP记录）
		content, err := ioutil.ReadFile(common.FileName)
		if err != nil {
			log.Error("读取文件错误: %s", err)
			log.Error("请按任意键退出...")
			fmt.Scanln()
			return // 遇到错误直接终止程序
		}

		// 去除前后空白字符并保存旧IP
		oldIP = strings.TrimSpace(string(content))

		// 旧IP处理逻辑（当旧IP存在且与新IP不同时）
		if oldIP != "" && oldIP != currentIP {
			log.Warn("发现记录IP %s", oldIP)
			log.Info("开始清理旧IP的安全组规则...")

			// 遍历所有云账户执行删除操作
			for _, account := range config.ConfigYaml.AliyunAccounts {
				log.Info("当前阿里云账户 [ %s ]", account.Name)

				// 调用安全组操作函数（revoke表示删除）
				err := common.HandleSecurityGroup(
					oldIP,                   // 要删除的IP
					"revoke",                // 操作类型
					account.Name,            // 账户名称
					account.RegionId,        // 地域ID
					account.AccessKey,       // 访问密钥
					account.AccessSecret,    // 访问密钥
					account.Policy,          // 授权策略
					account.PortRange,       // 端口范围
					account.IpProtocol,      // 协议类型
					account.SecurityGroupId, // 安全组ID
				)

				if err != nil {
					log.Error("删除失败: %s", err)
					continue // 单个账户失败不影响后续操作
				}
				log.Warn("成功删除 %s 的访问策略", oldIP)
				time.Sleep(time.Second * 1) // 防止API请求过频
			}
		} else {
			log.Warn("新旧IP相同 %s, 不进行添加操作", currentIP)
			log.Info("请按任意键退出...")
			fmt.Scanln()
			return // 直接结束程序
		}
	}

	// 新IP添加逻辑
	log.Info("开始添加新IP到安全组...")
	for _, account := range config.ConfigYaml.AliyunAccounts {
		log.Info("当前阿里云账户 [ %s ]", account.Name)

		// 调用安全组操作函数（authorize表示添加）
		err := common.HandleSecurityGroup(
			currentIP,               // 当前IP
			"authorize",             // 操作类型
			account.Name,            // 账户名称
			account.RegionId,        // 地域ID
			account.AccessKey,       // 访问密钥
			account.AccessSecret,    // 密钥
			account.Policy,          // 授权策略
			account.PortRange,       // 端口范围
			account.IpProtocol,      // 协议类型
			account.SecurityGroupId, // 安全组ID
		)

		if err != nil {
			log.Error("添加失败: %v", err)
			continue // 跳过当前账户继续执行
		}
		log.Warn("成功添加 %s 的访问策略", currentIP)
		time.Sleep(time.Second * 1) // 请求间隔
	}

	// 文件更新阶段
	log.Info("更新IP记录文件: %s", common.FileName)
	// 写入新IP到文件（覆盖模式）
	if err := ioutil.WriteFile(common.FileName, []byte(currentIP), 0644); err != nil {
		fmt.Println()
		log.Error("文件更新失败: %v", err)
		log.Error("请按任意键退出...")
		fmt.Scanln()
		return // 文件写入失败时终止
	}
	log.Info("程序执行完成") // 最终状态提示
	// 等待用户按任意键退出
	log.Info("请按任意键退出...")
	fmt.Scanln()
}
