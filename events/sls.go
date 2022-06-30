package events

type SlsEvent struct {
	CursorTime *int       `json:"cursorTime"`
	JobName    *string    `json:"jobName"`
	Parameter  *Parameter `json:"parameter"`
	Source     *Source    `json:"source"`
	TaskID     *string    `json:"taskId"`
}

type Parameter interface{}

type Source struct {
	BeginCursor  *string `json:"beginCursor"`
	EndCursor    *string `json:"endCursor"`
	Endpoint     *string `json:"endpoint"`
	LogstoreName *string `json:"logstoreName"`
	ProjectName  *string `json:"projectName"`
	ShardID      *int    `json:"shardId"`
}
