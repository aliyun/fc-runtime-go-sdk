package events

// TableStore 结构体定义了表格存储相关的事件信息。
// 它包含了版本信息和记录集。
type TableStore struct {
	Version string     `json:"Version"` // Version 表示表格存储的版本信息。
	Records []struct { // Records 是一个记录集合，包含了多条记录。
		Type string   `json:"Type"` // Type 表示记录的类型。
		Info struct { // Info 包含了记录的详细信息，例如时间戳。
			Timestamp int `json:"Timestamp"` // Timestamp 表示记录的创建或更新时间戳。
		} `json:"Info"`
		PrimaryKey []struct { // PrimaryKey 包含了记录的主键信息，包括列名和值。
			ColumnName string      `json:"ColumnName"` // ColumnName 表示主键列的名称。
			Value      interface{} `json:"Value"`      // Value 表示主键列的值。
		} `json:"PrimaryKey"`
		Columns []struct { // Columns 包含了记录的所有列信息，包括列名、值和时间戳。
			Type       string      `json:"Type"`       // Type 表示列的类型。
			ColumnName string      `json:"ColumnName"` // ColumnName 表示列的名称。
			Value      interface{} `json:"Value"`      // Value 表示列的值。
			Timestamp  int         `json:"Timestamp"`  // Timestamp 表示列值的修改时间戳。
		} `json:"Columns"`
	} `json:"Records"`
}
