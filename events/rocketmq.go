package events

import "time"

type RocketMQEvent struct {
	ID                      *string       `json:"id"`
	Source                  *string       `json:"source"`
	SpecVersion             *string       `json:"specversion"`
	Type                    *string       `json:"type"`
	DataContentType         *string       `json:"datacontenttype"`
	Subject                 *string       `json:"subject"`
	Time                    *time.Time    `json:"time"`
	AliyunAccountId         *string       `json:"aliyunaccountid"`
	AliyunPublishTime       *time.Time    `json:"aliyunpublishtime"`
	AliyunOriginalAccountId *string       `json:"aliyunoriginalaccountid"`
	AliyunEventBusName      *string       `json:"aliyuneventbusname"`
	AliyunRegionId          *string       `json:"aliyunregionid"`
	AliyunPublishAddr       *string       `json:"aliyunpublishaddr"`
	Data                    *RocketMQData `json:"data"`
}

type RocketMQData struct {
	Topic            string            `json:"topic"`
	SystemProperties map[string]string `json:"systemProperties"`
	UserProperties   map[string]string `json:"userProperties"`
	Body             string            `json:"body"`
}
