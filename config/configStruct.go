package config // 定义一个名为 "config" 的包

import ( // 导入外部依赖包
	_ "embed" // 必须导入 embed 包
	"fmt"
	"os" // 引入 os 包，主要用于操作文件系统，读取配置文件
	"whitelist/common"
	"whitelist/log"

	"gopkg.in/yaml.v2" // 引入 yaml 包，用于解析 YAML 格式的数据
)

// //go:embed config.yaml
// var defaultConfig []byte // 嵌入的配置文件内容

type Account struct { // 定义一个结构体类型 Account，用于存储与账户相关的配置信息
	Name            string `yaml:"name"`             // 在 YAML 中对应的键是 "name"，表示账户的名称
	RegionId        string `yaml:"regionId"`         // 在 YAML 中对应的键是 "regionId"，表示账户的区域 ID
	AccessKey       string `yaml:"access_key"`       // 在 YAML 中对应的键是 "access_key"，表示账户的访问密钥
	AccessSecret    string `yaml:"access_secret"`    // 在 YAML 中对应的键是 "access_secret"，表示账户的访问密钥的秘密
	Policy          string `yaml:"policy"`           // 在 YAML 中对应的键是 "policy"，表示账户的访问策略
	PortRange       string `yaml:"Port_Range"`       // 在 YAML 中对应的键是 "Port_Range"，表示端口范围
	IpProtocol      string `yaml:"Ip_Protocol"`      // 在 YAML 中对应的键是 "Ip_Protocol"，表示 IP 协议
	SecurityGroupId string `yaml:"Security_GroupId"` // 在 YAML 中对应的键是 "Security_GroupId"，表示安全组 ID 或名称
}

type Config struct { // 定义一个结构体类型 Config，用于存储整个配置文件的数据结构
	AliyunAccounts []Account `yaml:"aliyun_accounts"` // 在 YAML 中对应的键是 "aliyun_accounts"，表示阿里云账户信息的数组
}

var ConfigYaml Config // 定义一个全局变量 ConfigYaml，类型为 Config，用于存储解析后的配置信息

// 初始化函数，在程序启动时加载配置文件并解析
func init() { // init 函数在程序启动时自动调用
	// // 优先尝试使用嵌入的配置
	// if len(defaultConfig) > 0 {
	// 	err := yaml.Unmarshal(defaultConfig, &ConfigYaml)
	// 	if err == nil {
	// 		log.Info("成功加载嵌入的配置文件")
	// 		return
	// 	}
	// 	log.Warn("嵌入的配置文件解析失败: %v", err)
	// }

	// 如果嵌入配置不可用，回退到外部文件
	// 读取 config.yaml 配置文件
	yamlFile, err := os.ReadFile(common.ConfigFile) // 使用 os.ReadFile 读取配置文件内容
	if err != nil {                                 // 如果读取文件时出错，处理错误
		// 如果读取文件失败，抛出错误并终止程序
		log.Error("读取配置文件失败 %s", common.ConfigFile)
		log.Error("请按任意键退出...")
		fmt.Scanln()
		panic(err) // 通过 panic 终止程序并输出错误信息
	}
	// 解析 YAML 文件并将内容存入 ConfigYaml 变量
	err = yaml.Unmarshal(yamlFile, &ConfigYaml) // 使用 yaml.Unmarshal 解析 YAML 内容
	if err != nil {                             // 如果解析出错，处理错误
		// 如果解析 YAML 时出错，抛出错误并终止程序
		log.Error("解析配置文件失败 %s", common.ConfigFile)
		log.Error("请按任意键退出...")
		fmt.Scanln()
		panic(err) // 通过 panic 终止程序并输出错误信息
	}
}
