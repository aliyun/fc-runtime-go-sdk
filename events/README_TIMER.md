## Kafka Timer 触发器

### Event 格式

定时器触发器的 event 格式如下所示：

```json
{
    "triggerTime": "2022-07-29T10:02:58Z",
    "triggerName": "TestTimer",
    "payload": "TestPayload"
}
```

```bash
# 消息正文
'{ TestPayload }'
```
### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置定时器触发器之后，根据定时规则就会触发该函数，并返回消息内容。

```go
package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	"github.com/aliyun/fc-runtime-go-sdk/events"
)

func HandleRequest(ctx context.Context, event events.TimerEvent) (string, error) {
	fctx, ok := fccontext.FromContext(ctx)
	if !ok {
		return "Get fccontext fail.", nil
	}
	flog := fctx.GetLogger()

	flog.Info("triggerName: ", event.TriggerName)
	flog.Info("triggerTime: ", event.TriggerTime)
	flog.Info("payload:", event.Payload)

	return fmt.Sprintf("Timer Payload: %s", event.Payload), nil
}

func main() {
	fc.Start(HandleRequest)
}
```

