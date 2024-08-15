package events

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

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

func (receiver HTTPTriggerEvent) ReadParameter(name string) (context string, err error) {
	if receiver.QueryParameters == nil {
		return "", errors.New("query parameters is nil")
	}
	var exists bool
	context, exists = (*receiver.QueryParameters)[name]
	if !exists {
		return "", errors.New("target parameter not found")
	}
	return context, nil
}
func (receiver HTTPTriggerEvent) ReadHeader(name string) (context string, err error) {
	if receiver.Headers == nil {
		return "", errors.New("headers is nil")
	}
	var exists bool
	context, exists = (*receiver.Headers)[name]
	if !exists {
		return "", errors.New("target header not found")
	}
	return context, nil
}
func (receiver *HTTPTriggerResponse) SetStatusCode(StatusCode int) error {
	if StatusCode < 100 || StatusCode > 599 {
		return errors.New("invalid status code")
	}
	receiver.StatusCode = StatusCode
	return nil
}
func (receiver *HTTPTriggerResponse) SetHeader(name string, value string) {
	receiver.Headers[name] = value
}
func (receiver *HTTPTriggerResponse) WriteBody(context interface{}) error {
	val := reflect.ValueOf(context)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		receiver.Body = strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		receiver.Body = context.(string)
	case reflect.Struct:
		jsonByte, _ := json.Marshal(context)
		receiver.Body = string(jsonByte)
	case reflect.Invalid:
		return errors.New("invalid type")
	case reflect.Bool:
		receiver.Body = strconv.FormatBool(context.(bool))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		receiver.Body = strconv.FormatUint(val.Uint(), 10)
	case reflect.Uintptr:
	case reflect.Float32, reflect.Float64:
		receiver.Body = strconv.FormatFloat(val.Float(), 'f', -1, 64)
	case reflect.Complex64, reflect.Complex128:
		return errors.New("imaginary numbers are not supported, please assemble real and imaginary parts into one type, e.g., map,array")
	case reflect.Array:
		jsonByte, err := json.Marshal(context)
		if err != nil {
			return err
		}
		receiver.Body = string(jsonByte)
	case reflect.Chan:
		return errors.New("unsupported type")
	case reflect.Func:
		return errors.New("unsupported type")
	case reflect.Interface:
		return errors.New("interface{} are not guaranteed to be encoded")
	case reflect.Map:
		jsonByte, _ := json.Marshal(context)
		receiver.Body = string(jsonByte)
	case reflect.Slice:
		jsonByte, err := json.Marshal(context)
		if err != nil {
			return err
		}
		receiver.Body = string(jsonByte)
	case reflect.UnsafePointer:
		return errors.New("unsafe pointer")
	default:
		return errors.New("unsupported type")
	}
	return nil
}
