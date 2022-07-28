package events

import "time"

type MnsTopicEvent struct {
	TopicOwner       *string   `json:"TopicOwner"`
	Message          *string   `json:"Message"`
	Subscriber       *string   `json:"Subscriber"`
	PublishTime      time.Time `json:"PublishTime"`
	SubscriptionName *string   `json:"SubscriptionName"`
	MessageMD5       *string   `json:"MessageMD5"`
	TopicName        *string   `json:"TopicName"`
	MessageID        *string   `json:"MessageId"`
}
