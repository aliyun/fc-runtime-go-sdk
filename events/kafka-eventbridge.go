package events

type KafkaEventBridgeEvent struct {
	Data            *KafkaData `json:"data"`
	Id              *string    `json:"id"`
	Source          *string    `json:"source"`
	SpecVersion     *string    `json:"specversion"`
	Type            *string    `json:"type"`
	DataContentType *string    `json:"datacontenttype"`
	Time            *string    `json:"time"`
	Subject         *string    `json:"subject"`
	AliyunAccountId *string    `json:"aliyunaccountid"`
}

type KafkaData struct {
	Topic     *string          `json:"topic"`
	Partition *int             `json:"partition"`
	Offset    *int             `json:"offset"`
	Timestamp *int             `json:"timestamp"`
	Headers   *KafkaDataHeader `json:"headers"`
	Value     *string          `json:"value"`
}

type KafkaDataHeader struct {
	Headers    []*string `json:"headers"`
	IsReadOnly *bool     `json:"isReadOnly"`
}
