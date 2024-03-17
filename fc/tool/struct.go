package fc

import "time"

// EventStruct 是一个用于表示事件信息的结构体
type EventStruct struct {
	// Http 结构体包含HTTP请求的相关信息
	Http struct {
		Version         string         `json:"version"`         // HTTP版本
		RawPath         string         `json:"rawPath"`         // 未经解析的路径
		Body            string         `json:"body"`            // 请求体
		IsBase64Encoded bool           `json:"isBase64Encoded"` // 请求体是否以Base64编码
		Headers         map[string]any `json:"headers"`         // HTTP请求头
		QueryParameters map[string]any `json:"queryParameters"` // 查询参数
		// RequestContext 包含请求的上下文信息
		RequestContext struct {
			AccountId    string `json:"accountId"`    // 账户ID
			DomainName   string `json:"domainName"`   // 域名
			DomainPrefix string `json:"domainPrefix"` // 域名前缀
			// Http 结构体包含HTTP请求的详细信息
			Http struct {
				Method    string `json:"method"`    // 请求方法
				Path      string `json:"path"`      // 请求路径
				Protocol  string `json:"protocol"`  // 请求协议
				SourceIp  string `json:"sourceIp"`  // 请求来源IP
				UserAgent string `json:"userAgent"` // 用户代理
			} `json:"http"`
			RequestId string    `json:"requestId"` // 请求ID
			Time      time.Time `json:"time"`      // 请求时间
			TimeEpoch string    `json:"timeEpoch"` // 时间戳（epoch时间）
		} `json:"requestContext"`
	}
	// Timer 结构体
	// Timer 结构用于定义一个定时器，包含以下字段：
	// TriggerTime: 定时器触发的时间，是一个 time.Time 类型，使用 JSON 标记为 "triggerTime"。
	// TriggerName: 定时器的触发名称，是一个 string 类型，使用 JSON 标记为 "triggerName"。
	// Payload: 定时器触发时携带的负载信息，是一个 string 类型，使用 JSON 标记为 "payload"。
	Timer struct {
		TriggerTime time.Time `json:"triggerTime"` // 触发时间
		TriggerName string    `json:"triggerName"` // 触发名称
		Payload     string    `json:"payload"`     // 负载信息
	}
	// Oss 结构体定义了OSS事件的基本结构。
	Oss struct {
		// Events 是一个包含多个事件的切片。每个事件都包含了关于OSS事件的详细信息。
		Events []struct {
			EventName    string    `json:"eventName"`    // EventName 事件名称。
			EventSource  string    `json:"eventSource"`  // EventSource 事件来源。
			EventTime    time.Time `json:"eventTime"`    // EventTime 事件发生的时间。
			EventVersion string    `json:"eventVersion"` // EventVersion 事件版本。
			Oss          struct {  // Oss 包含了关于OSS事件的具体信息。
				Bucket struct { // Bucket 包含了关于发生事件的存储桶的信息。
					Arn           string `json:"arn"`           // Arn 存储桶的ARN（Amazon Resource Name）。
					Name          string `json:"name"`          // Name 存储桶的名称。
					OwnerIdentity string `json:"ownerIdentity"` // OwnerIdentity 存储桶所有者的身份信息。
				} `json:"bucket"`

				Object struct { // Object 包含了关于事件涉及的对象的信息。
					DeltaSize int    `json:"deltaSize"` // DeltaSize 对象的增量大小。
					ETag      string `json:"eTag"`      // ETag 对象的ETag，用于标识对象的内容。
					Key       string `json:"key"`       // Key 对象的键。
					Size      int    `json:"size"`      // Size 对象的大小。
				} `json:"object"`
				OssSchemaVersion string `json:"ossSchemaVersion"` // OssSchemaVersion OSS事件模式的版本。
				RuleId           string `json:"ruleId"`           // RuleId 触发此事件的规则ID。
			} `json:"oss"`
			Region            string   `json:"region"` // Region 事件发生的区域。
			RequestParameters struct { // RequestParameters 包含了请求参数的信息。
				SourceIPAddress string `json:"sourceIPAddress"` // SourceIPAddress 请求的源IP地址。
			} `json:"requestParameters"`

			ResponseElements struct { // ResponseElements 包含了响应元素的信息，如请求ID。
				RequestId string `json:"requestId"` // RequestId 请求的ID。
			} `json:"responseElements"`

			UserIdentity struct { // UserIdentity 包含了发起此事件的用户身份信息。
				PrincipalId string `json:"principalId"` // PrincipalId 用户的主要ID。
			} `json:"userIdentity"`
		} `json:"events"`
	}
	// Sls 结构体用于接收日志服务（SLS）的事件信息。
	Sls struct {
		// Parameter 字段用于存储自定义参数，键值对形式。
		Parameter map[string]any `json:"parameter"`
		// Source 字段定义了日志服务的源数据配置。
		Source struct {
			Endpoint     string `json:"endpoint"`     // 日志服务的访问Endpoint。
			ProjectName  string `json:"projectName"`  // 日志服务的项目名称。
			LogstoreName string `json:"logstoreName"` // 日志服务的日志库名称。
			ShardId      int    `json:"shardId"`      // 日志服务的分区ID。
			BeginCursor  string `json:"beginCursor"`  // 数据读取的起始游标。
			EndCursor    string `json:"endCursor"`    // 数据读取的结束游标。
		} `json:"source"`
		// JobName 定义了任务的名称。
		JobName string `json:"jobName"`
		// TaskId 用于标识具体的任务实例。
		TaskId string `json:"taskId"`
		// CursorTime 表示游标所指向的时间戳。
		CursorTime int `json:"cursorTime"`
	}
	// Cdn 结构体定义了CDN相关的日志和事件信息。
	Cdn struct {
		// LogFileCreated 包含日志文件创建事件的详细信息。
		LogFileCreated struct {
			Events []struct {
				EventName    string    `json:"eventName"`    // 事件名称
				EventSource  string    `json:"eventSource"`  // 事件来源
				Region       string    `json:"region"`       // 区域
				EventVersion string    `json:"eventVersion"` // 事件版本
				EventTime    time.Time `json:"eventTime"`    // 事件时间
				TraceId      string    `json:"traceId"`      // 追踪ID
				UserIdentity struct {
					AliUid string `json:"aliUid"` // 用户ID
				} `json:"userIdentity"`
				Resource struct {
					Domain string `json:"domain"` // 域名
				} `json:"resource"`
				EventParameter struct {
					Domain    string `json:"domain"`    // 域名
					EndTime   int    `json:"endTime"`   // 结束时间
					FileSize  int    `json:"fileSize"`  // 文件大小
					FilePath  string `json:"filePath"`  // 文件路径
					StartTime int    `json:"startTime"` // 开始时间
				} `json:"eventParameter"`
			} `json:"events"`
		}
		// CachedObjects 包含缓存对象事件的详细信息。
		CachedObjects struct {
			Events []struct {
				EventName    string    `json:"eventName"`    // 事件名称
				EventVersion string    `json:"eventVersion"` // 事件版本
				EventSource  string    `json:"eventSource"`  // 事件来源
				Region       string    `json:"region"`       // 区域
				EventTime    time.Time `json:"eventTime"`    // 事件时间
				TraceId      string    `json:"traceId"`      // 追踪ID
				Resource     struct {
					Domain string `json:"domain"` // 域名
				} `json:"resource"`
				EventParameter struct {
					ObjectPath   []string `json:"objectPath"`   // 对象路径
					CreateTime   int      `json:"createTime"`   // 创建时间
					Domain       string   `json:"domain"`       // 域名
					CompleteTime int      `json:"completeTime"` // 完成时间
					ObjectType   string   `json:"objectType"`   // 对象类型
					TaskId       int      `json:"taskId"`       // 任务ID
				} `json:"eventParameter"`
				UserIdentity struct {
					AliUid string `json:"aliUid"` // 用户ID
				} `json:"userIdentity"`
			} `json:"events"`
		}
		// CdnDomain 包含CDN域名相关事件的详细信息。
		CdnDomain struct {
			Events []struct {
				EventName    string    `json:"eventName"`    // 事件名称
				EventVersion string    `json:"eventVersion"` // 事件版本
				EventSource  string    `json:"eventSource"`  // 事件来源
				Region       string    `json:"region"`       // 区域
				EventTime    time.Time `json:"eventTime"`    // 事件时间
				TraceId      string    `json:"traceId"`      // 追踪ID
				Resource     struct {
					Domain string `json:"domain"` // 域名
				} `json:"resource"`
				EventParameter struct {
					Domain string `json:"domain"` // 域名
					Status string `json:"status"` // 状态
				} `json:"eventParameter"`
				UserIdentity struct {
					AliUid string `json:"aliUid"` // 用户ID
				} `json:"userIdentity"`
			} `json:"events"`
		}
	}
	// TableStore 结构体定义了表格存储相关的事件信息。
	// 它包含了版本信息和记录集。
	TableStore struct {
		Version string     `json:"Version"` // Version 表示表格存储的版本信息。
		Records []struct { // Records 是一个记录集合，包含了多条记录。
			Type string   `json:"Type"` // Type 表示记录的类型。
			Info struct { // Info 包含了记录的详细信息，例如时间戳。
				Timestamp int64 `json:"Timestamp"` // Timestamp 表示记录的创建或更新时间戳。
			} `json:"Info"`
			PrimaryKey []struct { // PrimaryKey 包含了记录的主键信息，包括列名和值。
				ColumnName string      `json:"ColumnName"` // ColumnName 表示主键列的名称。
				Value      interface{} `json:"Value"`      // Value 表示主键列的值。
			} `json:"PrimaryKey"`
			Columns []struct { // Columns 包含了记录的所有列信息，包括列名、值和时间戳。
				Type       string      `json:"Type"`       // Type 表示列的类型。
				ColumnName string      `json:"ColumnName"` // ColumnName 表示列的名称。
				Value      interface{} `json:"Value"`      // Value 表示列的值。
				Timestamp  int64       `json:"Timestamp"`  // Timestamp 表示列值的修改时间戳。
			} `json:"Columns"`
		} `json:"Records"`
	}
	// Mns 结构体定义了与消息服务（MNS）相关的主题和队列配置。
	Mns struct {
		Theme struct { // Theme 结构体包含有关主题的消息格式定义。
			Stream struct { // Stream 结构体定义了流式消息的属性。
				NoAttribute string   // NoAttribute 为无属性消息提供字段。
				Attribute   struct { // Attribute 结构体包含消息的属性。
					Body  string   `json:"body"` // Body 存储消息的正文内容。
					Attrs struct { // Attrs 结构体包含消息的额外属性。
						Extend string `json:"Extend"` // Extend 存储扩展属性信息。
					} `json:"attrs"`
				}
			}
			// Json 结构体定义了JSON格式消息的属性。
			Json struct {
				// NoAttribute 结构体为无属性消息提供字段。
				NoAttribute struct {
					TopicOwner       string `json:"TopicOwner"`       // 主题所有者
					Message          string `json:"Message"`          // 消息内容
					Subscriber       string `json:"Subscriber"`       // 订阅者
					PublishTime      int64  `json:"PublishTime"`      // 发布时间
					SubscriptionName string `json:"SubscriptionName"` // 订阅名称
					MessageMD5       string `json:"MessageMD5"`       // 消息MD5
					TopicName        string `json:"TopicName"`        // 主题名称
					MessageId        string `json:"MessageId"`        // 消息ID
				}
				// Attribute 结构体包含消息的属性。
				Attribute struct {
					Key              string `json:"key"`              // 消息键
					TopicOwner       string `json:"TopicOwner"`       // 主题所有者
					Message          string `json:"Message"`          // 消息内容
					Subscriber       string `json:"Subscriber"`       // 订阅者
					PublishTime      int64  `json:"PublishTime"`      // 发布时间
					SubscriptionName string `json:"SubscriptionName"` // 订阅名称
					MessageMD5       string `json:"MessageMD5"`       // 消息MD5
					TopicName        string `json:"TopicName"`        // 主题名称
					MessageId        string `json:"MessageId"`        // 消息ID
				}
			}
		}
		// Queue 结构体定义了队列的消息属性。
		Queue struct {
			Id                      string    `json:"id"`                      // 队列ID
			Source                  string    `json:"source"`                  // 源
			SpecVersion             string    `json:"specversion"`             // 规范版本
			Type                    string    `json:"type"`                    // 类型
			DataContentType         string    `json:"datacontenttype"`         // 数据内容类型
			Subject                 string    `json:"subject"`                 // 主题
			Time                    time.Time `json:"time"`                    // 时间
			AliyunAccountID         string    `json:"aliyunaccountid"`         // 阿里云账号ID
			AliyunPublishTime       time.Time `json:"aliyunpublishtime"`       // 阿里云发布时间
			AliyunOriginalAccountID string    `json:"aliyunoriginalaccountid"` // 原始阿里云账号ID
			AliyunEventbusName      string    `json:"aliyuneventbusname"`      // 阿里云事件总线名称
			AliyunRegionID          string    `json:"aliyunregionid"`          // 阿里云地域ID
			AliyunPublishAddr       string    `json:"aliyunpublishaddr"`       // 阿里云发布地址
			// Data 结构体包含队列消息的数据内容。
			Data struct {
				RequestId   string `json:"requestId"`   // 请求ID
				MessageId   string `json:"messageId"`   // 消息ID
				MessageBody string `json:"messageBody"` // 消息正文
			} `json:"data"`
		}
	}
	// ApiGate 结构体定义了API网关的基本信息
	// 其中包含了请求的路径、HTTP方法、头部信息、查询参数、路径参数、请求体以及请求体是否被Base64编码的标志。
	ApiGate struct {
		Path            string         `json:"path"`            // 请求的路径
		HttpMethod      string         `json:"httpMethod"`      // 请求的HTTP方法
		Headers         map[string]any `json:"headers"`         // 请求头部信息，键值对形式
		QueryParameters map[string]any `json:"queryParameters"` // 查询参数，键值对形式
		PathParameters  map[string]any `json:"pathParameters"`  // 路径参数，键值对形式
		Body            string         `json:"body"`            // 请求体
		IsBase64Encoded bool           `json:"isBase64Encoded"` // 标记请求体是否被Base64编码
	}
	// DataHub 结构体定义了数据枢纽的基本信息
	// 包括事件源、事件名称、事件源ARN、区域信息以及事件记录数组。
	DataHub struct {
		EventSource    string     `json:"eventSource"`    // 事件源
		EventName      string     `json:"eventName"`      // 事件名称
		EventSourceARN string     `json:"eventSourceARN"` // 事件源的ARN（Amazon Resource Name）
		Region         string     `json:"region"`         // 区域
		Records        []struct { // 事件记录的数组
			EventId    string `json:"eventId"`    // 事件ID
			SystemTime int64  `json:"systemTime"` // 系统时间，单位为毫秒
			Data       string `json:"data"`       // 事件数据
		} `json:"records"`
	}
	// RocketMQ 定义了RocketMQ消息的结构体列表
	RocketMQ []struct {
		Id                      string    `json:"id"`                      // 消息ID
		Source                  string    `json:"source"`                  // 消息来源
		SpecVersion             string    `json:"specversion"`             // 规范版本
		Type                    string    `json:"type"`                    // 消息类型
		DataContentType         string    `json:"datacontenttype"`         // 数据内容类型
		Subject                 string    `json:"subject"`                 // 主题
		Time                    time.Time `json:"time"`                    // 消息时间
		AliyunAccountID         string    `json:"aliyunaccountid"`         // 阿里云账号ID
		AliyunPublishTime       time.Time `json:"aliyunpublishtime"`       // 阿里云发布时间
		AliyunOriginalAccountID string    `json:"aliyunoriginalaccountid"` // 原始阿里云账号ID
		AliyunEventbusName      string    `json:"aliyuneventbusname"`      // 阿里云事件总线名称
		AliyunRegionID          string    `json:"aliyunregionid"`          // 阿里云地域ID
		AliyunPublishAddr       string    `json:"aliyunpublishaddr"`       // 阿里云发布地址
		Data                    struct {  // 消息数据
			Topic            string   `json:"topic"` // 主题
			SystemProperties struct { // 系统属性
				MinOffset        string `json:"MIN_OFFSET"`         // 最小偏移量
				TraceOn          string `json:"TRACE_ON"`           // 跟踪状态
				MaxOffset        string `json:"MAX_OFFSET"`         // 最大偏移量
				MSGRegion        string `json:"MSG_REGION"`         // 消息区域
				Keys             string `json:"KEYS"`               // 关键字
				ConsumeStartTime int64  `json:"CONSUME_START_TIME"` // 消费开始时间
				Tags             string `json:"TAGS"`               // 标签
				InstanceID       string `json:"INSTANCE_ID"`        // 实例ID
			} `json:"systemProperties"`
			UserProperties struct { // 用户属性
			} `json:"userProperties"`
			Body string `json:"body"` // 消息体
		} `json:"data"`
	}
	// RabbitMQ 定义了RabbitMQ消息的结构体列表
	RabbitMQ []struct {
		Id                      string    `json:"id"`                      // 消息ID
		Source                  string    `json:"source"`                  // 消息来源
		Specversion             string    `json:"specversion"`             // 规范版本
		Type                    string    `json:"type"`                    // 消息类型
		Datacontenttype         string    `json:"datacontenttype"`         // 数据内容类型
		Subject                 string    `json:"subject"`                 // 主题
		Time                    time.Time `json:"time"`                    // 消息时间
		Aliyunaccountid         string    `json:"aliyunaccountid"`         // 阿里云账号ID
		Aliyunpublishtime       time.Time `json:"aliyunpublishtime"`       // 阿里云发布时间
		Aliyunoriginalaccountid string    `json:"aliyunoriginalaccountid"` // 原始阿里云账号ID
		Aliyuneventbusname      string    `json:"aliyuneventbusname"`      // 阿里云事件总线名称
		Aliyunregionid          string    `json:"aliyunregionid"`          // 阿里云地域ID
		Aliyunpublishaddr       string    `json:"aliyunpublishaddr"`       // 阿里云发布地址
		Data                    struct {  // 消息数据
			Envelope struct { // 邮件封套信息
				DeliveryTag int    `json:"deliveryTag"` // 传递标签
				Exchange    string `json:"exchange"`    // 交换器
				Redeliver   bool   `json:"redeliver"`   // 是否重新传递
				RoutingKey  string `json:"routingKey"`  // 路由键
			} `json:"envelope"`
			Body struct { // 消息体
				Hello string `json:"Hello"` // 你好
			} `json:"body"`
			Props struct { // 消息属性
				ContentEncoding string `json:"contentEncoding"` // 内容编码
				MessageId       string `json:"messageId"`       // 消息ID
			} `json:"props"`
		} `json:"data"`
	}
	// Kafka 定义了一个结构体切片，用于表示Kafka消息的相关信息。
	Kafka []struct {
		Specversion     string    `json:"specversion"`     // Specversion 指示了规范的版本。
		Id              string    `json:"id"`              // Id 表示事件的唯一标识。
		Source          string    `json:"source"`          // Source 标识事件的来源。
		Type            string    `json:"type"`            // Type 指示事件的类型。
		Subject         string    `json:"subject"`         // Subject 是事件主体的标识。
		Datacontenttype string    `json:"datacontenttype"` // Datacontenttype 表明数据的内容类型。
		Time            time.Time `json:"time"`            // Time 记录了事件发生的时间。
		Aliyunaccountid string    `json:"aliyunaccountid"` // Aliyunaccountid 记录阿里云账号的ID。
		Data            struct {
			// Data 包含了事件的具体数据。
			Topic     string `json:"topic"`     // Topic 指示了消息所属的主题。
			Partition int    `json:"partition"` // Partition 表示消息所在的分区。
			Offset    int    `json:"offset"`    // Offset 指明了消息在分区中的偏移量。
			Timestamp int64  `json:"timestamp"` // Timestamp 记录了消息的 时间戳。
			Headers   struct {
				// Headers 包含了消息的头部信息。
				Headers    []interface{} `json:"headers"`    // Headers 存储了头部的键值对。
				IsReadOnly bool          `json:"isReadOnly"` // IsReadOnly 标记头部信息是否为只读。
			} `json:"headers"`
			Key   string `json:"key"`   // Key 指示了消息的键。
			Value string `json:"value"` // Value 存储了消息的值。
		} `json:"data"`
	}
	// Dts 定义了数据库表变更事件的详细信息集合
	// 每个元素记录了一次具体变更的前后图像、模式信息等
	Dts []struct {
		// Data 包含了关于数据库表变更的详细信息
		Data struct {
			Id             interface{} `json:"id"` // Id 为数据的唯一标识
			TopicPartition struct {    // TopicPartition 包含了主题分区的信息，如哈希值、分区号和主题名
				Hash      int    `json:"hash"`
				Partition int    `json:"partition"`
				Topic     string `json:"topic"`
			} `json:"topicPartition"`
			Offset          int    `json:"offset"`          // Offset 为数据的偏移量
			SourceTimestamp int    `json:"sourceTimestamp"` // SourceTimestamp 为源时间戳
			OperationType   string `json:"operationType"`   // OperationType 指明了操作类型
			// Schema 包含了数据模式的详细定义
			Schema struct {
				// RecordFields 定义了记录字段的详细信息
				RecordFields []struct {
					FieldName      string `json:"fieldName"`
					RawDataTypeNum int    `json:"rawDataTypeNum"`
					IsPrimaryKey   bool   `json:"isPrimaryKey"`
					IsUniqueKey    bool   `json:"isUniqueKey"`
					FieldPosition  int    `json:"fieldPosition"`
				} `json:"recordFields"`
				// NameIndex 为名称索引
				NameIndex struct {
					Id struct {
						FieldName      string `json:"fieldName"`
						RawDataTypeNum int    `json:"rawDataTypeNum"`
						IsPrimaryKey   bool   `json:"isPrimaryKey"`
						IsUniqueKey    bool   `json:"isUniqueKey"`
						FieldPosition  int    `json:"fieldPosition"`
					} `json:"id"`
					Topic struct {
						FieldName      string `json:"fieldName"`
						RawDataTypeNum int    `json:"rawDataTypeNum"`
						IsPrimaryKey   bool   `json:"isPrimaryKey"`
						IsUniqueKey    bool   `json:"isUniqueKey"`
						FieldPosition  int    `json:"fieldPosition"`
					} `json:"topic"`
				} `json:"nameIndex"`
				SchemaId         string   `json:"schemaId"`     // SchemaId 为模式的唯一标识
				DatabaseName     string   `json:"databaseName"` // DatabaseName 指明了数据库名称
				TableName        string   `json:"tableName"`    // TableName 指明了表名称
				PrimaryIndexInfo struct { // PrimaryIndexInfo 包含了主索引的详细信息
					IndexType   string `json:"indexType"`
					IndexFields []struct {
						FieldName      string `json:"fieldName"`
						RawDataTypeNum int    `json:"rawDataTypeNum"`
						IsPrimaryKey   bool   `json:"isPrimaryKey"`
						IsUniqueKey    bool   `json:"isUniqueKey"`
						FieldPosition  int    `json:"fieldPosition"`
					} `json:"indexFields"`
					Cardinality        int  `json:"cardinality"`
					Nullable           bool `json:"nullable"`
					IsFirstUniqueIndex bool `json:"isFirstUniqueIndex"`
				} `json:"primaryIndexInfo"`

				UniqueIndexInfo  []interface{} `json:"uniqueIndexInfo"`  // UniqueIndexInfo 包含了唯一索引的详细信息
				ForeignIndexInfo []interface{} `json:"foreignIndexInfo"` // ForeignIndexInfo 包含了外键索引的详细信息
				NormalIndexInfo  []interface{} `json:"normalIndexInfo"`  // NormalIndexInfo 包含了普通索引的详细信息
				// DatabaseInfo 包含了数据库的类型和版本信息
				DatabaseInfo struct {
					DatabaseType string `json:"databaseType"`
					Version      string `json:"version"`
				} `json:"databaseInfo"`
				TotalRows int `json:"totalRows"` // TotalRows 为表中的总行数
			} `json:"schema"`
			BeforeImage struct { // BeforeImage 包含了变更前的数据图像
				RecordSchema struct { // RecordSchema 包含了记录模式的详细定义
					RecordFields []struct {
						FieldName      string `json:"fieldName"`
						RawDataTypeNum int    `json:"rawDataTypeNum"`
						IsPrimaryKey   bool   `json:"isPrimaryKey"`
						IsUniqueKey    bool   `json:"isUniqueKey"`
						FieldPosition  int    `json:"fieldPosition"`
					} `json:"recordFields"`
					NameIndex struct {
						Id struct {
							FieldName      string `json:"fieldName"`
							RawDataTypeNum int    `json:"rawDataTypeNum"`
							IsPrimaryKey   bool   `json:"isPrimaryKey"`
							IsUniqueKey    bool   `json:"isUniqueKey"`
							FieldPosition  int    `json:"fieldPosition"`
						} `json:"id"`
						Topic struct {
							FieldName      string `json:"fieldName"`
							RawDataTypeNum int    `json:"rawDataTypeNum"`
							IsPrimaryKey   bool   `json:"isPrimaryKey"`
							IsUniqueKey    bool   `json:"isUniqueKey"`
							FieldPosition  int    `json:"fieldPosition"`
						} `json:"topic"`
					} `json:"nameIndex"`
					SchemaId         string `json:"schemaId"`
					DatabaseName     string `json:"databaseName"`
					TableName        string `json:"tableName"`
					PrimaryIndexInfo struct {
						IndexType   string `json:"indexType"`
						IndexFields []struct {
							FieldName      string `json:"fieldName"`
							RawDataTypeNum int    `json:"rawDataTypeNum"`
							IsPrimaryKey   bool   `json:"isPrimaryKey"`
							IsUniqueKey    bool   `json:"isUniqueKey"`
							FieldPosition  int    `json:"fieldPosition"`
						} `json:"indexFields"`
						Cardinality        int  `json:"cardinality"`
						Nullable           bool `json:"nullable"`
						IsFirstUniqueIndex bool `json:"isFirstUniqueIndex"`
					} `json:"primaryIndexInfo"`
					UniqueIndexInfo  []interface{} `json:"uniqueIndexInfo"`
					ForeignIndexInfo []interface{} `json:"foreignIndexInfo"`
					NormalIndexInfo  []interface{} `json:"normalIndexInfo"`
					DatabaseInfo     struct {
						DatabaseType string `json:"databaseType"`
						Version      string `json:"version"`
					} `json:"databaseInfo"`
					TotalRows int `json:"totalRows"`
				} `json:"recordSchema"`

				Values []struct { // Values 包含了变更前的实际数据值
					Data    interface{} `json:"data"`
					Charset string      `json:"charset,omitempty"`
				} `json:"values"`
				Size int `json:"size"` // Size 为变更前数据的大小
			} `json:"beforeImage"`
			AfterImage struct { // AfterImage 包含了变更后的数据图像
				RecordSchema struct { // RecordSchema 包含了记录模式的详细定义
					RecordFields []struct {
						FieldName      string `json:"fieldName"`
						RawDataTypeNum int    `json:"rawDataTypeNum"`
						IsPrimaryKey   bool   `json:"isPrimaryKey"`
						IsUniqueKey    bool   `json:"isUniqueKey"`
						FieldPosition  int    `json:"fieldPosition"`
					} `json:"recordFields"`
					NameIndex struct {
						Id struct {
							FieldName      string `json:"fieldName"`
							RawDataTypeNum int    `json:"rawDataTypeNum"`
							IsPrimaryKey   bool   `json:"isPrimaryKey"`
							IsUniqueKey    bool   `json:"isUniqueKey"`
							FieldPosition  int    `json:"fieldPosition"`
						} `json:"id"`
						Topic struct {
							FieldName      string `json:"fieldName"`
							RawDataTypeNum int    `json:"rawDataTypeNum"`
							IsPrimaryKey   bool   `json:"isPrimaryKey"`
							IsUniqueKey    bool   `json:"isUniqueKey"`
							FieldPosition  int    `json:"fieldPosition"`
						} `json:"topic"`
					} `json:"nameIndex"`
					SchemaId         string `json:"schemaId"`
					DatabaseName     string `json:"databaseName"`
					TableName        string `json:"tableName"`
					PrimaryIndexInfo struct {
						IndexType   string `json:"indexType"`
						IndexFields []struct {
							FieldName      string `json:"fieldName"`
							RawDataTypeNum int    `json:"rawDataTypeNum"`
							IsPrimaryKey   bool   `json:"isPrimaryKey"`
							IsUniqueKey    bool   `json:"isUniqueKey"`
							FieldPosition  int    `json:"fieldPosition"`
						} `json:"indexFields"`
						Cardinality        int  `json:"cardinality"`
						Nullable           bool `json:"nullable"`
						IsFirstUniqueIndex bool `json:"isFirstUniqueIndex"`
					} `json:"primaryIndexInfo"`
					UniqueIndexInfo  []interface{} `json:"uniqueIndexInfo"`
					ForeignIndexInfo []interface{} `json:"foreignIndexInfo"`
					NormalIndexInfo  []interface{} `json:"normalIndexInfo"`
					DatabaseInfo     struct {
						DatabaseType string `json:"databaseType"`
						Version      string `json:"version"`
					} `json:"databaseInfo"`
					TotalRows int `json:"totalRows"`
				} `json:"recordSchema"`

				Values []struct { // Values 包含了变更后的实际数据值
					Data    interface{} `json:"data"`
					Charset string      `json:"charset,omitempty"`
				} `json:"values"`
				Size int `json:"size"` // Size 为变更后数据的大小
			} `json:"afterImage"`
		} `json:"data"`

		Id              string    `json:"id"`              // Id 为事件的唯一标识
		Source          string    `json:"source"`          // Source 指明了事件的来源
		SpecVersion     string    `json:"specversion"`     // SpecVersion 指明了事件的规范版本
		Type            string    `json:"type"`            // Type 指明了事件的类型
		DataContentType string    `json:"datacontenttype"` // DataContentType 指明了数据的内容类型
		Time            time.Time `json:"time"`            // Time 指明了事件发生的时间
		Subject         string    `json:"subject"`         // Subject 指明了事件的主题
	}
	// MQTT 定义了处理MQTT消息的结构体
	// 其中包含了消息的属性和消息体
	MQTT []struct {
		Props struct { // Props 为MQTT消息的属性
			FirstTopic  string `json:"firstTopic"`  // 第一个主题
			SecondTopic string `json:"secondTopic"` // 第二个主题
			ClientId    string `json:"clientId"`    // 客户端ID
		} `json:"props"`
		Body string `json:"body"` // 消息体
	}
	// OffCloudRocketMQInstances 定义了处理OffCloud RocketMQ实例信息的结构体
	// 包含消息ID、主题、系统属性和用户属性以及消息体
	OffCloudRocketMQInstances []struct {
		MsgId            string   `json:"msgId"` // 消息ID
		Topic            string   `json:"topic"` // 主题
		SystemProperties struct { // 系统属性
			UniqKey   string `json:"UNIQ_KEY"`   // 唯一键
			Cluster   string `json:"CLUSTER"`    // 集群
			MinOffset string `json:"MIN_OFFSET"` // 最小偏移量
			Tags      string `json:"TAGS"`       // 标签
			MaxOffset string `json:"MAX_OFFSET"` // 最大偏移量
		} `json:"systemProperties"`
		UserProperties struct { // 用户属性，留空
		} `json:"userProperties"`
		Body string `json:"body"` // 消息体
	}
}

// HttpRepose 定义了HTTP响应的结构体
// 主要包括状态码、头部信息、是否以Base64编码、响应体四个部分
type HttpRepose struct {
	StatusCode      int16          `json:"statusCode"`
	Headers         map[string]any `json:"headers"`
	IsBase64Encoded bool           `json:"isBase64Encoded"`
	Body            string         `json:"body"`
}
