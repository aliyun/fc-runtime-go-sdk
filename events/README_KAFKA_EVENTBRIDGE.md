## Kafka Eventbridge 触发器

### Event 格式

mns topic 触发器的 event 格式如下所示：

- event为带有转义字符的字符串数组，

```json
["{\"data\":{\"topic\":\"HelloTopic\",\"partition\":9,\"offset\":3,\"timestamp\":1659346376797,\"headers\":{\"headers\":[],\"isReadOnly\":false},\"value\":\"b\\u0027{\\\\n    \\\"Test\\\": \\\"TestKafkaEventBridgetrigger\\\"\\\\n}\\u0027\"},\"id\":\"1cb591f9-987e-41d9-b974-0342e9acb90a\",\"source\":\"acs:alikafka\",\"specversion\":\"1.0\",\"type\":\"alikafka:Topic:Message\",\"datacontenttype\":\"application/json; charset\\u003dutf-8\",\"time\":\"2022-08-01T09:32:56.797Z\",\"subject\":\"acs:alikafka:alikafka_pre-cn-7pp2t2jwj001:topic:HelloTopic\",\"aliyunaccountid\":\"111111111111\"}"]
```

```bash
# 消息正文
'{ "Test" : "TestKafkaEventBridgetrigger" }'
```
### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置 Kafka Eventbridge 触发器之后，每当你向 Kafka topic 发送任意消息之后，就会触发该函数，并返回消息内容。

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	"github.com/aliyun/fc-runtime-go-sdk/events"
)

func HandleRequest(ctx context.Context, event []string) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "Get fccontext fail.", nil
	}
	flog := fctx.GetLogger()

	for _, eventString := range event {
		var evt events.KafkaEventBridgeEvent
		err := json.Unmarshal([]byte(eventString), &evt)
		if err != nil {
			return "Unmarshal event fail.", err
		}
		flog.Info("kafka event:", event)

		// The trigger event data is in the `Data` json object from the json array
		flog.Info("kafka topic:", evt.Data.Topic)
		flog.Info("kafka messgae:", evt.Data.Value)
	}

	return fmt.Sprintf("Receive Kafka Trigger Event: %v", event), nil
}
func main() {
	fc.Start(HandleRequest)
}

```

