package events

// HTTPTriggerEvent 事件输入
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

// HTTPTriggerRequestContext 用于存储上下文信息
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
	RequestId *string `json:"requestId"` // 请求 ID
	Time      *string `json:"time"`      // 请求时间
	TimeEpoch *string `json:"timeEpoch"` // 时间戳（epoch 时间）
}

// HTTPTriggerResponse 定义了 HTTP 响应的结构体
// 主要包括状态码、头部信息、是否以 Base64 编码、响应体四个部分
type HTTPTriggerResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	Body            string            `json:"body"`
}
