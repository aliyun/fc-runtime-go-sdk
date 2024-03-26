## FC3 HTTP 结构体

### 结构体部分用途解释
输入：
```go
package example
// HttpRequest 考虑到1.18前不支持map[string]any，所以使用map[string]interface{}代替，可以视为等效。
type HttpRequest struct {
	Version         *string                 `json:"version"`         // HTTP版本
	RawPath         *string                 `json:"rawPath"`         // 未经解析的路径
	Body            *string                 `json:"body"`            // 请求体
	IsBase64Encoded *bool                   `json:"isBase64Encoded"` // 请求体是否以 Base64 编码
	Headers         *map[string]interface{} `json:"headers"`         // HTTP请求头，以键值对存储
	QueryParameters *map[string]interface{} `json:"queryParameters"` // 查询参数，以键值对存储
	// RequestContext 包含请求的上下文信息
	RequestContext RequestContext `json:"requestContext"`
}
```
输出：
```go
package example
// HttpResponse 同上，可以视为等效。
type HttpResponse struct {
	StatusCode      *int                    `json:"statusCode"`     // HTTP 状态码
	Headers         *map[string]interface{} `json:"headers"`        // HTTP 响应头，以键值对存储
	IsBase64Encoded *bool                   `json:"isBase64Encoded"`// 响应体是否以 Base64 编码
	Body            *string                 `json:"body"`           // 响应体
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

func HandleRequest(ctx context.Context, event *events.HttpRequest) (repose *events.HttpRepose, err error) {
	// 读取 params
	var paramsValue any
	var ok bool
	if event.QueryParameters == nil {
		repose = new(events.HttpRepose)             // 创建响应对象
		*repose.StatusCode = http.StatusBadRequest  // 设置状态码
		*repose.Body = "params is nil"              // 设置响应体
		*repose.IsBase64Encoded = false             // 设置响应体是否以 Base64 编码
		Headers := *repose.Headers                  // 解引用指针
		Headers["Content-Type"] = "text/plain"      // 设置响应头
		*repose.Headers = Headers                   // 引用指针
		return repose, nil
	} else {
		paramsValue, ok = (*event.QueryParameters)["params"]
	}
	// 遍历获取 param
	for k, v := range *event.QueryParameters {
		fmt.Println(k, v)
	}
	repose = new(events.HttpRepose)             // 创建响应对象
	*repose.StatusCode = http.StatusOK          // 设置状态码
	*repose.Body = ""                           // 设置响应体
	*repose.IsBase64Encoded = false             // 设置响应体是否以 Base64 编码
	Headers := *repose.Headers                  // 解引用指针
	Headers["Content-Type"] = "application/json"// 设置响应头
	*repose.Headers = Headers                   // 引用指针
	return repose, nil
}
func main() {
	fc.Start(HandleRequest)
}
```
