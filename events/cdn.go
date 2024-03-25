package events

import "time"

// Cdn 结构体定义了CDN相关的日志和事件信息。
type Cdn struct {
	// LogFileCreated 包含日志文件创建事件的详细信息。
	LogFileCreated struct {
		Events []struct {
			EventName    *string   `json:"eventName"`    // 事件名称
			EventSource  *string   `json:"eventSource"`  // 事件来源
			Region       *string   `json:"region"`       // 区域
			EventVersion *string   `json:"eventVersion"` // 事件版本
			EventTime    time.Time `json:"eventTime"`    // 事件时间
			TraceId      *string   `json:"traceId"`      // 追踪ID
			UserIdentity struct {
				AliUid *string `json:"aliUid"` // 用户ID
			} `json:"userIdentity"`
			Resource struct {
				Domain *string `json:"domain"` // 域名
			} `json:"resource"`
			EventParameter struct {
				Domain    *string `json:"domain"`    // 域名
				EndTime   *int    `json:"endTime"`   // 结束时间
				FileSize  *int    `json:"fileSize"`  // 文件大小
				FilePath  *string `json:"filePath"`  // 文件路径
				StartTime *int    `json:"startTime"` // 开始时间
			} `json:"eventParameter"`
		} `json:"events"`
	}
	// CachedObjects 包含缓存对象事件的详细信息。
	CachedObjects struct {
		Events []struct {
			EventName    *string   `json:"eventName"`    // 事件名称
			EventVersion *string   `json:"eventVersion"` // 事件版本
			EventSource  *string   `json:"eventSource"`  // 事件来源
			Region       *string   `json:"region"`       // 区域
			EventTime    time.Time `json:"eventTime"`    // 事件时间
			TraceId      *string   `json:"traceId"`      // 追踪ID
			Resource     struct {
				Domain *string `json:"domain"` // 域名
			} `json:"resource"`
			EventParameter struct {
				ObjectPath   []string `json:"objectPath"`   // 对象路径
				CreateTime   *int     `json:"createTime"`   // 创建时间
				Domain       *string  `json:"domain"`       // 域名
				CompleteTime *int     `json:"completeTime"` // 完成时间
				ObjectType   *string  `json:"objectType"`   // 对象类型
				TaskId       *int     `json:"taskId"`       // 任务ID
			} `json:"eventParameter"`
			UserIdentity struct {
				AliUid *string `json:"aliUid"` // 用户ID
			} `json:"userIdentity"`
		} `json:"events"`
	}
	// CdnDomain 包含CDN域名相关事件的详细信息。
	CdnDomain struct {
		Events []struct {
			EventName    *string    `json:"eventName"`    // 事件名称
			EventVersion *string    `json:"eventVersion"` // 事件版本
			EventSource  *string    `json:"eventSource"`  // 事件来源
			Region       *string    `json:"region"`       // 区域
			EventTime    *time.Time `json:"eventTime"`    // 事件时间
			TraceId      *string    `json:"traceId"`      // 追踪ID
			Resource     struct {
				Domain *string `json:"domain"` // 域名
			} `json:"resource"`
			EventParameter struct {
				Domain *string `json:"domain"` // 域名
				Status *string `json:"status"` // 状态
			} `json:"eventParameter"`
			UserIdentity struct {
				AliUid *string `json:"aliUid"` // 用户ID
			} `json:"userIdentity"`
		} `json:"events"`
	}
}
