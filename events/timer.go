package events

type TimerEvent struct {
	TriggerTime *string `json:"triggerTime"`
	TriggerName *string `json:"triggerName"`
	Payload     *string `json:"payload"`
}
