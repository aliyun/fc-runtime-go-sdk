package events

import "time"

// HttpRequest 事件输入
type HttpRequest struct {
	Version         *string                 `json:"version"`         // HTTP版本
	RawPath         *string                 `json:"rawPath"`         // 未经解析的路径
	Body            *string                 `json:"body"`            // 请求体
	IsBase64Encoded *bool                   `json:"isBase64Encoded"` // 请求体是否以Base64编码
	Headers         *map[string]interface{} `json:"headers"`         // HTTP请求头
	QueryParameters *map[string]interface{} `json:"queryParameters"` // 查询参数
	// RequestContext 包含请求的上下文信息
	RequestContext RequestContext `json:"requestContext"`
}

// RequestContext 用于存储上下文信息
type RequestContext struct {
	AccountId    *string `json:"accountId"`    // 账户ID
	DomainName   *string `json:"domainName"`   // 域名
	DomainPrefix *string `json:"domainPrefix"` // 域名前缀
	// HttpRequest 结构体包含HTTP请求的详细信息
	Http struct {
		Method    *string `json:"method"`    // 请求方法
		Path      *string `json:"path"`      // 请求路径
		Protocol  *string `json:"protocol"`  // 请求协议
		SourceIp  *string `json:"sourceIp"`  // 请求来源IP
		UserAgent *string `json:"userAgent"` // 用户代理
	} `json:"http"`
	RequestId *string    `json:"requestId"` // 请求ID
	Time      *time.Time `json:"time"`      // 请求时间
	TimeEpoch *string    `json:"timeEpoch"` // 时间戳（epoch时间）
}

// HttpRepose 定义了HTTP响应的结构体
// 主要包括状态码、头部信息、是否以Base64编码、响应体四个部分
type HttpRepose struct {
	StatusCode      *int                    `json:"statusCode"`
	Headers         *map[string]interface{} `json:"headers"`
	IsBase64Encoded *bool                   `json:"isBase64Encoded"`
	Body            *string                 `json:"body"`
}
