package events

import "time"

type OssEvent struct {
	Events []OssEventRecord `json:"events"`
}

type OssEventRecord struct {
	EventName         string            `json:"eventName"`
	EventSource       string            `json:"eventSource"`
	EventTime         time.Time         `json:"eventTime"`
	EventVersion      string            `json:"eventVersion"`
	Oss               Oss               `json:"oss"`
	Region            string            `json:"region"`
	RequestParameters RequestParameters `json:"requestParameters"`
	ResponseElements  ResponseElements  `json:"responseElements"`
	UserIdentity      UserIdentity      `json:"userIdentity"`
}

type Oss struct {
	Bucket           Bucket `json:"bucket"`
	Object           Object `json:"object"`
	OssSchemaVersion string `json:"ossSchemaVersion"`
	RuleID           string `json:"ruleId"`
}

type Bucket struct {
	Arn           string `json:"arn"`
	Name          string `json:"name"`
	OwnerIdentity string `json:"ownerIdentity"`
	VirtualBucket string `json:"virtualBucket"`
}

type Object struct {
	DeltaSize  int        `json:"deltaSize"`
	ETag       string     `json:"eTag"`
	Key        string     `json:"key"`
	ObjectMeta ObjectMeta `json:"objectMeta"`
	Size       int        `json:"size"`
}

type ObjectMeta struct {
	MimeType string `json:"mimeType"`
}

type RequestParameters struct {
	SourceIPAddress string `json:"sourceIPAddress"`
}

type ResponseElements struct {
	RequestID string `json:"requestId"`
}

type UserIdentity struct {
	PrincipalID string `json:"principalId"`
}
