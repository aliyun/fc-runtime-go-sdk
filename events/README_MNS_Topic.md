## MNS Topic 触发器

### Event 格式

mns topic 触发器的 event 格式如下所示：

- 若event格式设置为 JSON

```json
{
  "TopicOwner":"topic account id",
  "Message":"mock mns message",
  "Subscriber":"subscriber account id",
  "PublishTime":1658235558094,
  "SubscriptionName":"test-5bf13c7e",
  "MessageMD5":"652BF0E6297840015247C3xxxxxxx",
  "TopicName":"fc-example",
  "MessageId":"3405CA51807661353B3xxxxxxxx"
}
```
- 若event格式设置为 STREAM

```bash
# 消息正文
'hello topic'
```
### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置 mns topic 触发器之后，每当你向 mns topic 发送任意消息之后，就会触发该函数，并返回消息内容。

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/events"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	var topicEvent events.MnsTopicEvent
	switch mess := event.(type) {
	case string:
		{
			flog.Info("event:", event)
			return fmt.Sprintf("MessageBody:%s", mess), nil
		}
	default:
		result, _ := json.Marshal(mess)
		_ = json.Unmarshal(result, &topicEvent)
		flog.Info("event:", event)
		return fmt.Sprintf("MessageBody:%s", *topicEvent.Message), nil
	}

}

func main() {
	fc.Start(HandleRequest)
}

```

