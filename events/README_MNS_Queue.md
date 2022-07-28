## MNS Queue 触发器

### Event 格式

mns queue 触发器的 event 格式如下所示：

```json
{
  "id":"c2g71017-6f65-fhcf-a814-a396fc8d****",
  "source":"MNS-Function-mnstrigger",
  "specversion":"1.0",
  "type":"mns:Queue:SendMessage",
  "datacontenttype":"application/json; charset=utf-8",
  "subject":"acs:mns:cn-hangzhou:164901546557****:queues/zeus",
  "time":"2021-04-08T06:28:17.093Z",
  "aliyunaccountid":"1649015465574023",
  "aliyunpublishtime":"2021-10-15T07:06:34.028Z",
  "aliyunoriginalaccountid":"164901546557****",
  "aliyuneventbusname":"MNS-Function-mnstrigger",
  "aliyunregionid":"cn-chengdu",
  "aliyunpublishaddr":"42.120.XX.XX",
  "data":{
    "requestId":"606EA3074344430D4C81****",
    "messageId":"C6DB60D1574661357FA227277445****",
    "messageBody":"TEST"
  }
}
```

### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置 mns queue 触发器之后，每当你向 mns queue 发送任意消息之后，就会触发该函数，并返回消息内容。

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

func HandleRequest(ctx context.Context, event events.MnsQueueEvent) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	flog := fctx.GetLogger()
	mes, _ := json.Marshal(event)
	flog.Info("event:", string(mes))
	return fmt.Sprintf("MessageBody:%s", *event.Data.MessageBody), nil
}

func main() {
	fc.Start(HandleRequest)
}
```

