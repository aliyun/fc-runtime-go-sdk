package events

// MQTT 定义了处理MQTT消息的结构体
// 其中包含了消息的属性和消息体
type MQTT []struct {
	Props struct { // Props 为MQTT消息的属性
		FirstTopic  *string `json:"firstTopic"`  // 第一个主题
		SecondTopic *string `json:"secondTopic"` // 第二个主题
		ClientId    *string `json:"clientId"`    // 客户端ID
	} `json:"props"`
	Body *string `json:"body"` // 消息体
}
