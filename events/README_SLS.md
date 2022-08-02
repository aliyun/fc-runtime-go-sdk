## SLS 触发器

### Event 格式

sls 触发器的 event 格式如下所示：

```json
{
  "cursorTime": 1655645277,
  "jobName": "39955366992fd7a0ee85a9e1ecba042a5c737682",
  "parameter": {},   
  "source": {
    "beginCursor": "MTY1NDUzOTMyNjA5GFH1MzYyNf==",
    "endCursor": "MTY1NDUzOTMyNjA5GFH1MzYyOQ==",
    "endpoint": "http://cn-shenzhen-intranet.log.aliyuncs.com",
    "logstoreName": "function-log",
    "projectName": "aliyun-fc-cn-shenzhen-3a8b13243-b031-5b64-abd3-5b54********",
    "shardId": 0
  },
  "taskId": "617efabb-8dc3-41fv-b908-7cbc********"
}
```

### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置 sls 触发器之后，sls 触发器会定时获取日志库里更新的数据并触发函数的执行，该函数反馈给你日志库等信息。

```go
package main

import (
	"context"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/events"
)

func HandleRequest(ctx context.Context, event events.SlsEvent) (string, error) {
	fmt.Printf("hello,the name of your logstoreName is %s", *event.Source.LogstoreName)
	return fmt.Sprintf("hello,the name of your projectName is %s",*event.Source.ProjectName), nil
}

func main() {
	fc.Start(HandleRequest)
}
```

