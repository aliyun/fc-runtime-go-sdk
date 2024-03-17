package fc

import "time"

type EventStruct struct {
	Http struct {
		Version         string         `json:"version"`
		RawPath         string         `json:"rawPath"`
		Body            string         `json:"body"`
		IsBase64Encoded bool           `json:"isBase64Encoded"`
		Headers         map[string]any `json:"headers"`
		QueryParameters map[string]any `json:"queryParameters"`
		RequestContext  struct {
			AccountId    string `json:"accountId"`
			DomainName   string `json:"domainName"`
			DomainPrefix string `json:"domainPrefix"`
			Http         struct {
				Method    string `json:"method"`
				Path      string `json:"path"`
				Protocol  string `json:"protocol"`
				SourceIp  string `json:"sourceIp"`
				UserAgent string `json:"userAgent"`
			} `json:"http"`
			RequestId string    `json:"requestId"`
			Time      time.Time `json:"time"`
			TimeEpoch string    `json:"timeEpoch"`
		} `json:"requestContext"`
	}
	Timer struct {
		TriggerTime time.Time `json:"triggerTime"`
		TriggerName string    `json:"triggerName"`
		Payload     string    `json:"payload"`
	}
	Oss struct {
		Events []struct {
			EventName    string    `json:"eventName"`
			EventSource  string    `json:"eventSource"`
			EventTime    time.Time `json:"eventTime"`
			EventVersion string    `json:"eventVersion"`
			Oss          struct {
				Bucket struct {
					Arn           string `json:"arn"`
					Name          string `json:"name"`
					OwnerIdentity string `json:"ownerIdentity"`
				} `json:"bucket"`
				Object struct {
					DeltaSize int    `json:"deltaSize"`
					ETag      string `json:"eTag"`
					Key       string `json:"key"`
					Size      int    `json:"size"`
				} `json:"object"`
				OssSchemaVersion string `json:"ossSchemaVersion"`
				RuleId           string `json:"ruleId"`
			} `json:"oss"`
			Region            string `json:"region"`
			RequestParameters struct {
				SourceIPAddress string `json:"sourceIPAddress"`
			} `json:"requestParameters"`
			ResponseElements struct {
				RequestId string `json:"requestId"`
			} `json:"responseElements"`
			UserIdentity struct {
				PrincipalId string `json:"principalId"`
			} `json:"userIdentity"`
		} `json:"events"`
	}
	Sls struct {
		Parameter map[string]any `json:"parameter"`
		Source    struct {
			Endpoint     string `json:"endpoint"`
			ProjectName  string `json:"projectName"`
			LogstoreName string `json:"logstoreName"`
			ShardId      int    `json:"shardId"`
			BeginCursor  string `json:"beginCursor"`
			EndCursor    string `json:"endCursor"`
		} `json:"source"`
		JobName    string `json:"jobName"`
		TaskId     string `json:"taskId"`
		CursorTime int    `json:"cursorTime"`
	}
	Cdn struct {
		LogFileCreated struct {
			Events []struct {
				EventName    string    `json:"eventName"`
				EventSource  string    `json:"eventSource"`
				Region       string    `json:"region"`
				EventVersion string    `json:"eventVersion"`
				EventTime    time.Time `json:"eventTime"`
				TraceId      string    `json:"traceId"`
				UserIdentity struct {
					AliUid string `json:"aliUid"`
				} `json:"userIdentity"`
				Resource struct {
					Domain string `json:"domain"`
				} `json:"resource"`
				EventParameter struct {
					Domain    string `json:"domain"`
					EndTime   int    `json:"endTime"`
					FileSize  int    `json:"fileSize"`
					FilePath  string `json:"filePath"`
					StartTime int    `json:"startTime"`
				} `json:"eventParameter"`
			} `json:"events"`
		}
		CachedObjects struct {
			Events []struct {
				EventName    string    `json:"eventName"`
				EventVersion string    `json:"eventVersion"`
				EventSource  string    `json:"eventSource"`
				Region       string    `json:"region"`
				EventTime    time.Time `json:"eventTime"`
				TraceId      string    `json:"traceId"`
				Resource     struct {
					Domain string `json:"domain"`
				} `json:"resource"`
				EventParameter struct {
					ObjectPath   []string `json:"objectPath"`
					CreateTime   int      `json:"createTime"`
					Domain       string   `json:"domain"`
					CompleteTime int      `json:"completeTime"`
					ObjectType   string   `json:"objectType"`
					TaskId       int      `json:"taskId"`
				} `json:"eventParameter"`
				UserIdentity struct {
					AliUid string `json:"aliUid"`
				} `json:"userIdentity"`
			} `json:"events"`
		}
		CdnDomain struct {
			Events []struct {
				EventName    string    `json:"eventName"`
				EventVersion string    `json:"eventVersion"`
				EventSource  string    `json:"eventSource"`
				Region       string    `json:"region"`
				EventTime    time.Time `json:"eventTime"`
				TraceId      string    `json:"traceId"`
				Resource     struct {
					Domain string `json:"domain"`
				} `json:"resource"`
				EventParameter struct {
					Domain string `json:"domain"`
					Status string `json:"status"`
				} `json:"eventParameter"`
				UserIdentity struct {
					AliUid string `json:"aliUid"`
				} `json:"userIdentity"`
			} `json:"events"`
		}
	}
	TableStore struct {
		Version string `json:"Version"`
		Records []struct {
			Type string `json:"Type"`
			Info struct {
				Timestamp int64 `json:"Timestamp"`
			} `json:"Info"`
			PrimaryKey []struct {
				ColumnName string      `json:"ColumnName"`
				Value      interface{} `json:"Value"`
			} `json:"PrimaryKey"`
			Columns []struct {
				Type       string      `json:"Type"`
				ColumnName string      `json:"ColumnName"`
				Value      interface{} `json:"Value"`
				Timestamp  int64       `json:"Timestamp"`
			} `json:"Columns"`
		} `json:"Records"`
	}
	Mns struct {
		Theme struct {
			Stream struct {
				NoAttribute string
				Attribute   struct {
					Body  string `json:"body"`
					Attrs struct {
						Extend string `json:"Extend"`
					} `json:"attrs"`
				}
			}
			Json struct {
				NoAttribute struct {
					TopicOwner       string `json:"TopicOwner"`
					Message          string `json:"Message"`
					Subscriber       string `json:"Subscriber"`
					PublishTime      int64  `json:"PublishTime"`
					SubscriptionName string `json:"SubscriptionName"`
					MessageMD5       string `json:"MessageMD5"`
					TopicName        string `json:"TopicName"`
					MessageId        string `json:"MessageId"`
				}
				Attribute struct {
					Key              string `json:"key"`
					TopicOwner       string `json:"TopicOwner"`
					Message          string `json:"Message"`
					Subscriber       string `json:"Subscriber"`
					PublishTime      int64  `json:"PublishTime"`
					SubscriptionName string `json:"SubscriptionName"`
					MessageMD5       string `json:"MessageMD5"`
					TopicName        string `json:"TopicName"`
					MessageId        string `json:"MessageId"`
				}
			}
		}
		Queue struct {
			Id                      string    `json:"id"`
			Source                  string    `json:"source"`
			SpecVersion             string    `json:"specversion"`
			Type                    string    `json:"type"`
			DataContentType         string    `json:"datacontenttype"`
			Subject                 string    `json:"subject"`
			Time                    time.Time `json:"time"`
			AliyunAccountID         string    `json:"aliyunaccountid"`
			AliyunPublishTime       time.Time `json:"aliyunpublishtime"`
			AliyunOriginalAccountID string    `json:"aliyunoriginalaccountid"`
			AliyunEventbusName      string    `json:"aliyuneventbusname"`
			AliyunRegionID          string    `json:"aliyunregionid"`
			AliyunPublishAddr       string    `json:"aliyunpublishaddr"`
			Data                    struct {
				RequestId   string `json:"requestId"`
				MessageId   string `json:"messageId"`
				MessageBody string `json:"messageBody"`
			} `json:"data"`
		}
	}
	ApiGate struct {
		Path            string         `json:"path"`
		HttpMethod      string         `json:"httpMethod"`
		Headers         map[string]any `json:"headers"`
		QueryParameters map[string]any `json:"queryParameters"`
		PathParameters  map[string]any `json:"pathParameters"`
		Body            string         `json:"body"`
		IsBase64Encoded bool           `json:"isBase64Encoded"`
	}
	DataHub struct {
		EventSource    string `json:"eventSource"`
		EventName      string `json:"eventName"`
		EventSourceARN string `json:"eventSourceARN"`
		Region         string `json:"region"`
		Records        []struct {
			EventId    string `json:"eventId"`
			SystemTime int64  `json:"systemTime"`
			Data       string `json:"data"`
		} `json:"records"`
	}
	RocketMQ []struct {
		Id                      string    `json:"id"`
		Source                  string    `json:"source"`
		SpecVersion             string    `json:"specversion"`
		Type                    string    `json:"type"`
		DataContentType         string    `json:"datacontenttype"`
		Subject                 string    `json:"subject"`
		Time                    time.Time `json:"time"`
		AliyunAccountID         string    `json:"aliyunaccountid"`
		AliyunPublishTime       time.Time `json:"aliyunpublishtime"`
		AliyunOriginalAccountID string    `json:"aliyunoriginalaccountid"`
		AliyunEventbusName      string    `json:"aliyuneventbusname"`
		AliyunRegionID          string    `json:"aliyunregionid"`
		AliyunPublishAddr       string    `json:"aliyunpublishaddr"`
		Data                    struct {
			Topic            string `json:"topic"`
			SystemProperties struct {
				MinOffset        string `json:"MIN_OFFSET"`
				TraceOn          string `json:"TRACE_ON"`
				MaxOffset        string `json:"MAX_OFFSET"`
				MSGRegion        string `json:"MSG_REGION"`
				Keys             string `json:"KEYS"`
				ConsumeStartTime int64  `json:"CONSUME_START_TIME"`
				Tags             string `json:"TAGS"`
				InstanceID       string `json:"INSTANCE_ID"`
			} `json:"systemProperties"`
			UserProperties struct {
			} `json:"userProperties"`
			Body string `json:"body"`
		} `json:"data"`
	}
	RabbitMQ []struct {
		Id                      string    `json:"id"`
		Source                  string    `json:"source"`
		Specversion             string    `json:"specversion"`
		Type                    string    `json:"type"`
		Datacontenttype         string    `json:"datacontenttype"`
		Subject                 string    `json:"subject"`
		Time                    time.Time `json:"time"`
		Aliyunaccountid         string    `json:"aliyunaccountid"`
		Aliyunpublishtime       time.Time `json:"aliyunpublishtime"`
		Aliyunoriginalaccountid string    `json:"aliyunoriginalaccountid"`
		Aliyuneventbusname      string    `json:"aliyuneventbusname"`
		Aliyunregionid          string    `json:"aliyunregionid"`
		Aliyunpublishaddr       string    `json:"aliyunpublishaddr"`
		Data                    struct {
			Envelope struct {
				DeliveryTag int    `json:"deliveryTag"`
				Exchange    string `json:"exchange"`
				Redeliver   bool   `json:"redeliver"`
				RoutingKey  string `json:"routingKey"`
			} `json:"envelope"`
			Body struct {
				Hello string `json:"Hello"`
			} `json:"body"`
			Props struct {
				ContentEncoding string `json:"contentEncoding"`
				MessageId       string `json:"messageId"`
			} `json:"props"`
		} `json:"data"`
	}
	Kafka []struct {
		Specversion     string    `json:"specversion"`
		Id              string    `json:"id"`
		Source          string    `json:"source"`
		Type            string    `json:"type"`
		Subject         string    `json:"subject"`
		Datacontenttype string    `json:"datacontenttype"`
		Time            time.Time `json:"time"`
		Aliyunaccountid string    `json:"aliyunaccountid"`
		Data            struct {
			Topic     string `json:"topic"`
			Partition int    `json:"partition"`
			Offset    int    `json:"offset"`
			Timestamp int64  `json:"timestamp"`
			Headers   struct {
				Headers    []interface{} `json:"headers"`
				IsReadOnly bool          `json:"isReadOnly"`
			} `json:"headers"`
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"data"`
	}
	Dts []struct {
		Data struct {
			Id             interface{} `json:"id"`
			TopicPartition struct {
				Hash      int    `json:"hash"`
				Partition int    `json:"partition"`
				Topic     string `json:"topic"`
			} `json:"topicPartition"`
			Offset          int    `json:"offset"`
			SourceTimestamp int    `json:"sourceTimestamp"`
			OperationType   string `json:"operationType"`
			Schema          struct {
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
			} `json:"schema"`
			BeforeImage struct {
				RecordSchema struct {
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
				Values []struct {
					Data    interface{} `json:"data"`
					Charset string      `json:"charset,omitempty"`
				} `json:"values"`
				Size int `json:"size"`
			} `json:"beforeImage"`
			AfterImage struct {
				RecordSchema struct {
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
				Values []struct {
					Data    interface{} `json:"data"`
					Charset string      `json:"charset,omitempty"`
				} `json:"values"`
				Size int `json:"size"`
			} `json:"afterImage"`
		} `json:"data"`
		Id              string    `json:"id"`
		Source          string    `json:"source"`
		SpecVersion     string    `json:"specversion"`
		Type            string    `json:"type"`
		DataContentType string    `json:"datacontenttype"`
		Time            time.Time `json:"time"`
		Subject         string    `json:"subject"`
	}
	MQTT []struct {
		Props struct {
			FirstTopic  string `json:"firstTopic"`
			SecondTopic string `json:"secondTopic"`
			ClientId    string `json:"clientId"`
		} `json:"props"`
		Body string `json:"body"`
	}
	OffCloudRocketMQInstances []struct {
		MsgId            string `json:"msgId"`
		Topic            string `json:"topic"`
		SystemProperties struct {
			UniqKey   string `json:"UNIQ_KEY"`
			Cluster   string `json:"CLUSTER"`
			MinOffset string `json:"MIN_OFFSET"`
			Tags      string `json:"TAGS"`
			MaxOffset string `json:"MAX_OFFSET"`
		} `json:"systemProperties"`
		UserProperties struct {
		} `json:"userProperties"`
		Body string `json:"body"`
	}
}

type HttpRepose struct {
	StatusCode      int16          `json:"statusCode"`
	Headers         map[string]any `json:"headers"`
	IsBase64Encoded bool           `json:"isBase64Encoded"`
	Body            string         `json:"body"`
}
