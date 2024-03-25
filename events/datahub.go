package events

// DataHub 结构体定义了数据枢纽的基本信息
// 包括事件源、事件名称、事件源ARN、区域信息以及事件记录数组。
type DataHub struct {
	EventSource    *string    `json:"eventSource"`    // 事件源
	EventName      *string    `json:"eventName"`      // 事件名称
	EventSourceARN *string    `json:"eventSourceARN"` // 事件源的ARN（Amazon Resource Name）
	Region         *string    `json:"region"`         // 区域
	Records        []struct { // 事件记录的数组
		EventId    *string `json:"eventId"`    // 事件ID
		SystemTime *int    `json:"systemTime"` // 系统时间，单位为毫秒
		Data       *string `json:"data"`       // 事件数据
	} `json:"records"`
}
