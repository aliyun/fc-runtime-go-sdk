package events

//
// // ApiGate 结构体定义了API网关的基本信息
// // 其中包含了请求的路径、HTTP方法、头部信息、查询参数、路径参数、请求体以及请求体是否被Base64编码的标志。
// type ApiGate struct {
// 	Path            *string         `json:"path"`            // 请求的路径
// 	HttpMethod      *string         `json:"httpMethod"`      // 请求的HTTP方法
// 	Headers         *map[string]any `json:"headers"`         // 请求头部信息，键值对形式
// 	QueryParameters *map[string]any `json:"queryParameters"` // 查询参数，键值对形式
// 	PathParameters  *map[string]any `json:"pathParameters"`  // 路径参数，键值对形式
// 	Body            *string         `json:"body"`            // 请求体
// 	IsBase64Encoded bool            `json:"isBase64Encoded"` // 标记请求体是否被Base64编码
// }
