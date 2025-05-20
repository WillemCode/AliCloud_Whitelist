package common // 定义一个名为 "common" 的包

import ( // 导入需要的包
	"encoding/json" // 引入 JSON 编码/解码包，用于处理 JSON 格式的响应
	"whitelist/log"

	// 引入 fmt 包，用于格式化输出
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs" // 引入阿里云 SDK，操作 ECS（弹性计算服务）
)

// 创建一个处理安全组中 IP 的函数，接收两个值：IP 地址和动作 ("authorize" 为增加 IP，"revoke" 为删除安全组中的 IP)
func HandleSecurityGroup(
	ip string,
	action string,
	Name string,
	RegionId string,
	AccessKey string,
	AccessSecret string,
	Policy string,
	PortRange string,
	IpProtocol string,
	SecurityGroupId string,
) error {

	// 创建 ECS 客户端，使用提供的阿里云区域、访问密钥和密钥进行身份验证
	client, err := ecs.NewClientWithAccessKey(RegionId, AccessKey, AccessSecret)
	if err != nil { // 如果客户端创建失败，输出错误信息
		// fmt.Print(err.Error()) // 输出错误信息
		log.Error(err.Error())
	}

	if action == "authorize" { // 如果动作是 "authorize"，表示要增加 IP 到安全组
		// 创建一个新的授权请求实例
		request := ecs.CreateAuthorizeSecurityGroupRequest()
		request.Scheme = "https"                  // 设置请求的协议为 https
		request.Policy = Policy                   // 设置访问策略（如 "Accept"）
		request.NicType = "internet"              // 设置网络类型为 internet（公网）
		request.Priority = "1"                    // 设置优先级
		request.PortRange = PortRange             // 设置端口范围
		request.IpProtocol = IpProtocol           // 设置 IP 协议（如 TCP、UDP）
		request.Description = "加白程序自动维护"          // 设置描述信息，记录此操作是由程序自动进行的，方便追溯
		request.SourceCidrIp = ip                 // 设置允许访问的 IP 地址
		request.SecurityGroupId = SecurityGroupId // 设置要操作的安全组 ID

		// 执行授权请求，向安全组添加该 IP
		response, err := client.AuthorizeSecurityGroup(request)
		if err != nil { // 如果请求出错，返回错误
			return err
		}
		// 将响应数据格式化为 JSON 并输出
		jsonBytes, err := json.MarshalIndent(response, "", "    ")
		if err != nil { // 如果 JSON 格式化出错，返回错误
			return err
		}
		// 输出请求 ID
		log.Info("请求ID: %s", string(jsonBytes))
	} else { // 如果动作是 "revoke"，表示要删除 IP 从安全组中
		// 创建一个新的撤销请求实例
		request := ecs.CreateRevokeSecurityGroupRequest()
		request.Scheme = "https"                  // 设置请求的协议为 https
		request.Policy = Policy                   // 设置访问策略
		request.NicType = "internet"              // 设置网络类型为 internet（公网）
		request.Priority = "1"                    // 设置优先级
		request.PortRange = PortRange             // 设置端口范围
		request.IpProtocol = IpProtocol           // 设置 IP 协议
		request.Description = "加白程序自动维护"          // 设置描述信息
		request.SourceCidrIp = ip                 // 设置需要撤销的 IP 地址
		request.SecurityGroupId = SecurityGroupId // 设置要操作的安全组 ID

		// 执行撤销请求，从安全组中删除该 IP
		response, err := client.RevokeSecurityGroup(request)
		if err != nil { // 如果请求出错，返回错误
			return err
		}
		// 将响应数据格式化为 JSON 并输出
		jsonBytes, err := json.MarshalIndent(response, "", "    ")
		if err != nil { // 如果 JSON 格式化出错，返回错误
			return err
		}
		// 输出请求 ID
		log.Info("请求ID: %s", string(jsonBytes))
	}
	return nil // 返回 nil，表示操作成功
}
