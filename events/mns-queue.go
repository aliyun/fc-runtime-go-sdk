package events

import "time"

type MnsQueueEvent struct {
	ID                      *string   `json:"id"`
	Source                  *string   `json:"source"`
	Specversion             *string   `json:"specversion"`
	Type                    *string   `json:"type"`
	Datacontenttype         *string   `json:"datacontenttype"`
	Subject                 *string   `json:"subject"`
	Time                    time.Time `json:"time"`
	Aliyunaccountid         *string   `json:"aliyunaccountid"`
	Aliyunpublishtime       time.Time `json:"aliyunpublishtime"`
	Aliyunoriginalaccountid *string   `json:"aliyunoriginalaccountid"`
	Aliyuneventbusname      *string   `json:"aliyuneventbusname"`
	Aliyunregionid          *string   `json:"aliyunregionid"`
	Aliyunpublishaddr       *string   `json:"aliyunpublishaddr"`
	Data                    *Data     `json:"data"`
}
type Data struct {
	RequestID   *string `json:"requestId"`
	MessageID   *string `json:"messageId"`
	MessageBody *string `json:"messageBody"`
}
