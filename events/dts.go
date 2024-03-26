package events

//
// // Dts 定义了数据库表变更事件的详细信息集合
// // 每个元素记录了一次具体变更的前后图像、模式信息等
// type Dts []struct {
// 	// Data 包含了关于数据库表变更的详细信息
// 	Data struct {
// 		Id             *interface{} `json:"id"` // Id 为数据的唯一标识
// 		TopicPartition struct {     // TopicPartition 包含了主题分区的信息，如哈希值、分区号和主题名
// 			Hash      *int    `json:"hash"`
// 			Partition *int    `json:"partition"`
// 			Topic     *string `json:"topic"`
// 		} `json:"topicPartition"`
// 		Offset          *int    `json:"offset"`          // Offset 为数据的偏移量
// 		SourceTimestamp *int    `json:"sourceTimestamp"` // SourceTimestamp 为源时间戳
// 		OperationType   *string `json:"operationType"`   // OperationType 指明了操作类型
// 		// Schema 包含了数据模式的详细定义
// 		Schema struct {
// 			// RecordFields 定义了记录字段的详细信息
// 			RecordFields []struct {
// 				FieldName      *string `json:"fieldName"`
// 				RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 				IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 				IsUniqueKey    bool    `json:"isUniqueKey"`
// 				FieldPosition  *int    `json:"fieldPosition"`
// 			} `json:"recordFields"`
// 			// NameIndex 为名称索引
// 			NameIndex struct {
// 				Id struct {
// 					FieldName      *string `json:"fieldName"`
// 					RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 					IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 					IsUniqueKey    bool    `json:"isUniqueKey"`
// 					FieldPosition  *int    `json:"fieldPosition"`
// 				} `json:"id"`
// 				Topic struct {
// 					FieldName      *string `json:"fieldName"`
// 					RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 					IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 					IsUniqueKey    bool    `json:"isUniqueKey"`
// 					FieldPosition  *int    `json:"fieldPosition"`
// 				} `json:"topic"`
// 			} `json:"nameIndex"`
// 			SchemaId         *string  `json:"schemaId"`     // SchemaId 为模式的唯一标识
// 			DatabaseName     *string  `json:"databaseName"` // DatabaseName 指明了数据库名称
// 			TableName        *string  `json:"tableName"`    // TableName 指明了表名称
// 			PrimaryIndexInfo struct { // PrimaryIndexInfo 包含了主索引的详细信息
// 				IndexType   *string `json:"indexType"`
// 				IndexFields []struct {
// 					FieldName      *string `json:"fieldName"`
// 					RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 					IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 					IsUniqueKey    bool    `json:"isUniqueKey"`
// 					FieldPosition  *int    `json:"fieldPosition"`
// 				} `json:"indexFields"`
// 				Cardinality        *int `json:"cardinality"`
// 				Nullable           bool `json:"nullable"`
// 				IsFirstUniqueIndex bool `json:"isFirstUniqueIndex"`
// 			} `json:"primaryIndexInfo"`
//
// 			UniqueIndexInfo  *[]interface{} `json:"uniqueIndexInfo"`  // UniqueIndexInfo 包含了唯一索引的详细信息
// 			ForeignIndexInfo *[]interface{} `json:"foreignIndexInfo"` // ForeignIndexInfo 包含了外键索引的详细信息
// 			NormalIndexInfo  *[]interface{} `json:"normalIndexInfo"`  // NormalIndexInfo 包含了普通索引的详细信息
// 			// DatabaseInfo 包含了数据库的类型和版本信息
// 			DatabaseInfo struct {
// 				DatabaseType *string `json:"databaseType"`
// 				Version      *string `json:"version"`
// 			} `json:"databaseInfo"`
// 			TotalRows *int `json:"totalRows"` // TotalRows 为表中的总行数
// 		} `json:"schema"`
// 		BeforeImage struct { // BeforeImage 包含了变更前的数据图像
// 			RecordSchema struct { // RecordSchema 包含了记录模式的详细定义
// 				RecordFields []struct {
// 					FieldName      *string `json:"fieldName"`
// 					RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 					IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 					IsUniqueKey    bool    `json:"isUniqueKey"`
// 					FieldPosition  *int    `json:"fieldPosition"`
// 				} `json:"recordFields"`
// 				NameIndex struct {
// 					Id struct {
// 						FieldName      *string `json:"fieldName"`
// 						RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 						IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 						IsUniqueKey    bool    `json:"isUniqueKey"`
// 						FieldPosition  *int    `json:"fieldPosition"`
// 					} `json:"id"`
// 					Topic struct {
// 						FieldName      *string `json:"fieldName"`
// 						RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 						IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 						IsUniqueKey    bool    `json:"isUniqueKey"`
// 						FieldPosition  *int    `json:"fieldPosition"`
// 					} `json:"topic"`
// 				} `json:"nameIndex"`
// 				SchemaId         *string `json:"schemaId"`
// 				DatabaseName     *string `json:"databaseName"`
// 				TableName        *string `json:"tableName"`
// 				PrimaryIndexInfo struct {
// 					IndexType   *string `json:"indexType"`
// 					IndexFields []struct {
// 						FieldName      *string `json:"fieldName"`
// 						RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 						IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 						IsUniqueKey    bool    `json:"isUniqueKey"`
// 						FieldPosition  *int    `json:"fieldPosition"`
// 					} `json:"indexFields"`
// 					Cardinality        *int `json:"cardinality"`
// 					Nullable           bool `json:"nullable"`
// 					IsFirstUniqueIndex bool `json:"isFirstUniqueIndex"`
// 				} `json:"primaryIndexInfo"`
// 				UniqueIndexInfo  *[]interface{} `json:"uniqueIndexInfo"`
// 				ForeignIndexInfo *[]interface{} `json:"foreignIndexInfo"`
// 				NormalIndexInfo  *[]interface{} `json:"normalIndexInfo"`
// 				DatabaseInfo     struct {
// 					DatabaseType *string `json:"databaseType"`
// 					Version      *string `json:"version"`
// 				} `json:"databaseInfo"`
// 				TotalRows *int `json:"totalRows"`
// 			} `json:"recordSchema"`
//
// 			Values []struct { // Values 包含了变更前的实际数据值
// 				Data    *interface{} `json:"data"`
// 				Charset *string      `json:"charset,omitempty"`
// 			} `json:"values"`
// 			Size *int `json:"size"` // Size 为变更前数据的大小
// 		} `json:"beforeImage"`
// 		AfterImage struct { // AfterImage 包含了变更后的数据图像
// 			RecordSchema struct { // RecordSchema 包含了记录模式的详细定义
// 				RecordFields []struct {
// 					FieldName      *string `json:"fieldName"`
// 					RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 					IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 					IsUniqueKey    bool    `json:"isUniqueKey"`
// 					FieldPosition  *int    `json:"fieldPosition"`
// 				} `json:"recordFields"`
// 				NameIndex struct {
// 					Id struct {
// 						FieldName      *string `json:"fieldName"`
// 						RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 						IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 						IsUniqueKey    bool    `json:"isUniqueKey"`
// 						FieldPosition  *int    `json:"fieldPosition"`
// 					} `json:"id"`
// 					Topic struct {
// 						FieldName      *string `json:"fieldName"`
// 						RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 						IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 						IsUniqueKey    bool    `json:"isUniqueKey"`
// 						FieldPosition  *int    `json:"fieldPosition"`
// 					} `json:"topic"`
// 				} `json:"nameIndex"`
// 				SchemaId         *string `json:"schemaId"`
// 				DatabaseName     *string `json:"databaseName"`
// 				TableName        *string `json:"tableName"`
// 				PrimaryIndexInfo struct {
// 					IndexType   *string `json:"indexType"`
// 					IndexFields []struct {
// 						FieldName      *string `json:"fieldName"`
// 						RawDataTypeNum *int    `json:"rawDataTypeNum"`
// 						IsPrimaryKey   bool    `json:"isPrimaryKey"`
// 						IsUniqueKey    bool    `json:"isUniqueKey"`
// 						FieldPosition  *int    `json:"fieldPosition"`
// 					} `json:"indexFields"`
// 					Cardinality        *int `json:"cardinality"`
// 					Nullable           bool `json:"nullable"`
// 					IsFirstUniqueIndex bool `json:"isFirstUniqueIndex"`
// 				} `json:"primaryIndexInfo"`
// 				UniqueIndexInfo  *[]interface{} `json:"uniqueIndexInfo"`
// 				ForeignIndexInfo *[]interface{} `json:"foreignIndexInfo"`
// 				NormalIndexInfo  *[]interface{} `json:"normalIndexInfo"`
// 				DatabaseInfo     struct {
// 					DatabaseType *string `json:"databaseType"`
// 					Version      *string `json:"version"`
// 				} `json:"databaseInfo"`
// 				TotalRows *int `json:"totalRows"`
// 			} `json:"recordSchema"`
//
// 			Values []struct { // Values 包含了变更后的实际数据值
// 				Data    *interface{} `json:"data"`
// 				Charset *string      `json:"charset,omitempty"`
// 			} `json:"values"`
// 			Size *int `json:"size"` // Size 为变更后数据的大小
// 		} `json:"afterImage"`
// 	} `json:"data"`
//
// 	Id              *string   `json:"id"`              // Id 为事件的唯一标识
// 	Source          *string   `json:"source"`          // Source 指明了事件的来源
// 	SpecVersion     *string   `json:"specversion"`     // SpecVersion 指明了事件的规范版本
// 	Type            *string   `json:"type"`            // Type 指明了事件的类型
// 	DataContentType *string   `json:"datacontenttype"` // DataContentType 指明了数据的内容类型
// 	Time            time.Time `json:"time"`            // Time 指明了事件发生的时间
// 	Subject         *string   `json:"subject"`         // Subject 指明了事件的主题
// }
