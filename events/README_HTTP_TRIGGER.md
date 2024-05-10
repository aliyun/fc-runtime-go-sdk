## FC3 HTTP 结构体

### 结构体部分用途解释
输入：

```go
package events

import "time"
// HTTPTriggerEvent 函数传入结构体
type HTTPTriggerEvent struct {
	Version         *string            `json:"version"`         // HTTP 触发器请求事件版本
	RawPath         *string            `json:"rawPath"`         // 未经解析的路径
	Body            *string            `json:"body"`            // 请求体
	IsBase64Encoded *bool              `json:"isBase64Encoded"` // 请求体是否以 Base64 编码
	Headers         *map[string]string `json:"headers"`         // HTTP 请求头
	QueryParameters *map[string]string `json:"queryParameters"` // 查询 Param
	// RequestContext 包含请求的上下文信息
	TriggerContext HTTPTriggerRequestContext `json:"requestContext"`
}
// HTTPTriggerRequestContext 上下文结构体，用于存储http触发器专属上下文
type HTTPTriggerRequestContext struct {
	AccountId    *string `json:"accountId"`    // 账户 ID
	DomainName   *string `json:"domainName"`   // 域名
	DomainPrefix *string `json:"domainPrefix"` // 域名前缀
	// Http 结构体包含 HTTP 请求的详细信息
	Http struct {
		Method    *string `json:"method"`    // 请求方法
		Path      *string `json:"path"`      // 请求路径
		Protocol  *string `json:"protocol"`  // 请求协议
		SourceIp  *string `json:"sourceIp"`  // 请求来源 IP
		UserAgent *string `json:"userAgent"` // 客户端标识，如: Mozilla/5.0
	} `json:"http"`
	RequestId *string    `json:"requestId"` // 请求 ID
	Time      *time.Time `json:"time"`      // 请求时间
	TimeEpoch *string    `json:"timeEpoch"` // 时间戳（epoch 时间）
}
```
输出：

```go
package events
// HTTPTriggerResponse 回复结构体，您必须使用该结构体编码回复并返回，方可获得有效数据
// 主要包括状态码、头部信息、是否以 Base64 编码、响应体四个部分
type HTTPTriggerResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	Body            string             `json:"body"`
}

```
### 使用示例
这里展示了如何读取和输入

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpTriggerRequest(ctx context.Context, event events.HTTPTriggerEvent) (response *events.HTTPTriggerResponse, err error) {
	// 初始化 param 存储
	var paramsValue string
	var ok bool
	if event.QueryParameters == nil {
		*response.StatusCode = http.StatusBadRequest // 设置状态码
		*response.Body = "params is nil"             // 设置响应体
		*response.IsBase64Encoded = false            // 设置响应体是否以 Base64 编码
		Headers := *response.Headers                 // 解引用指针
		Headers["Content-Type"] = "text/plain"       // 设置响应头
		*response.Headers = Headers                  // 引用指针
		return response, nil
	} else {
		paramsValue, _ = (event.QueryParameters)["params"] // 读取 param 样例
	}
	// 遍历获取并打印 param
	for param := range event.QueryParameters {
		log.Printf("param: %#+v", param)
	}
	response = new(events.HTTPTriggerResponse)
	response.StatusCode = http.StatusOK                   // 设置状态码
	response.Body = "Hello world!"                        // 设置响应体
	response.IsBase64Encoded = false                      // 设置响应体是否以 Base64 编码
	response.Headers["Content-Type"] = "application/json" // 设置响应头
	return response, nil
}
func main() {
	fc.Start(HandleHttpTriggerRequest)
}
```