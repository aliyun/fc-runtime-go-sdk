package events

//
// // RabbitMQ 定义了RabbitMQ消息的结构体列表
// type RabbitMQ []struct {
// 	Id                      *string   `json:"id"`                      // 消息ID
// 	Source                  *string   `json:"source"`                  // 消息来源
// 	Specversion             *string   `json:"specversion"`             // 规范版本
// 	Type                    *string   `json:"type"`                    // 消息类型
// 	Datacontenttype         *string   `json:"datacontenttype"`         // 数据内容类型
// 	Subject                 *string   `json:"subject"`                 // 主题
// 	Time                    time.Time `json:"time"`                    // 消息时间
// 	Aliyunaccountid         *string   `json:"aliyunaccountid"`         // 阿里云账号ID
// 	Aliyunpublishtime       time.Time `json:"aliyunpublishtime"`       // 阿里云发布时间
// 	Aliyunoriginalaccountid *string   `json:"aliyunoriginalaccountid"` // 原始阿里云账号ID
// 	Aliyuneventbusname      *string   `json:"aliyuneventbusname"`      // 阿里云事件总线名称
// 	Aliyunregionid          *string   `json:"aliyunregionid"`          // 阿里云地域ID
// 	Aliyunpublishaddr       *string   `json:"aliyunpublishaddr"`       // 阿里云发布地址
// 	Data                    struct {  // 消息数据
// 		Envelope struct { // 邮件封套信息
// 			DeliveryTag *int    `json:"deliveryTag"` // 传递标签
// 			Exchange    *string `json:"exchange"`    // 交换器
// 			Redeliver   bool    `json:"redeliver"`   // 是否重新传递
// 			RoutingKey  *string `json:"routingKey"`  // 路由键
// 		} `json:"envelope"`
// 		Body struct { // 消息体
// 			Hello *string `json:"Hello"` // 你好
// 		} `json:"body"`
// 		Props struct { // 消息属性
// 			ContentEncoding *string `json:"contentEncoding"` // 内容编码
// 			MessageId       *string `json:"messageId"`       // 消息ID
// 		} `json:"props"`
// 	} `json:"data"`
// }
