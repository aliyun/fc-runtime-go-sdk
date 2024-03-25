## FC3 HTTP 结构体

### 结构体部分用途解释
输入：
```go
package events
// HTTPTirgger
type HTTPTriggerEventContext struct {
	Version         *string                 `json:"version"`         // HTTP 触发器请求事件版本
	RawPath         *string                 `json:"rawPath"`         // 未经解析的路径
	Body            *string                 `json:"body"`            // 请求体
	IsBase64Encoded *bool                   `json:"isBase64Encoded"` // 请求体是否以 Base64 编码
	Headers         *map[string]string      `json:"headers"`         // HTTP 请求头，以键值对存储
	QueryParameters *map[string]string      `json:"queryParameters"` // 查询参数，以键值对存储
	// RequestContext 包含请求的上下文信息
	RequestContext HTTPTriggerRequestContext `json:"requestContext"`
}
```
输出：

```go
package events

// HTTPTriggerRequestContext 回复结构体，您必须使用该结构体编码回复并返回，方可获得有效数据
type HTTPTriggerRequestContext struct {
	StatusCode      *int               `json:"statusCode"`      // HTTP 状态码
	Headers         *map[string]string `json:"headers"`         // HTTP 响应头，以键值对存储
	IsBase64Encoded *bool              `json:"isBase64Encoded"` // 响应体是否以 Base64 编码
	Body            *string            `json:"body"`            // 响应体
}
```
### 使用示例

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpTriggerRequest(ctx context.Context, event *events.HTTPTriggerEventContext) (repose *events.HTTPTriggerRequestContext, err error) {
	// 初始化 param 存储
	var paramsValue string
	var ok bool
	if event.QueryParameters == nil {
		repose = new(events.HTTPTriggerRequestContext)  // 创建响应对象
		*repose.StatusCode = http.StatusBadRequest      // 设置状态码
		*repose.Body = "params is nil"                  // 设置响应体
		*repose.IsBase64Encoded = false                 // 设置响应体是否以 Base64 编码
		Headers := *repose.Headers                      // 解引用指针
		Headers["Content-Type"] = "text/plain"          // 设置响应头
		*repose.Headers = Headers                       // 引用指针
		return repose, nil
	} else {
		paramsValue, _ = (*event.QueryParameters)["params"] // 读取 param 样例
	}
	// 遍历获取并打印 param
	for k, v := range *event.QueryParameters {
		fmt.Println(k, v)
	}
	repose = new(events.HTTPTriggerRequestContext)  // 创建响应对象
	*repose.StatusCode = http.StatusOK              // 设置状态码
	*repose.Body = "Hello world!"                   // 设置响应体
	*repose.IsBase64Encoded = false                 // 设置响应体是否以 Base64 编码
	Headers := *repose.Headers                      // 解引用指针
	Headers["Content-Type"] = "application/json"    // 设置响应头
	*repose.Headers = Headers                       // 引用指针
	return repose, nil
}
func main() {
	fc.Start(HandleHttpTriggerRequest)
}
```