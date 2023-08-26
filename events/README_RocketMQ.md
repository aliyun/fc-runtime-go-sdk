# RocketMQ 触发器

### Event 格式

```json
{
    "id":"94ebc15f-f0db-4bbe-acce-56fb72fb****",
    "source":"RocketMQ-Function-rocketmq-trigger",
    "specversion":"1.0",
    "type":"mq:Topic:SendMessage",
    "datacontenttype":"application/json; charset=utf-8",
    "subject":"acs:mq:cn-hangzhou:164901546557****:MQ_INST_164901546557****_BXhFHryi%TopicName",
    "time":"2021-04-08T06:01:20.766Z",
    "aliyunaccountid":"164901546557****",
    "aliyunpublishtime":"2021-10-15T02:05:16.791Z",
    "aliyunoriginalaccountid":"164901546557****",
    "aliyuneventbusname":"RocketMQ-Function-rocketmq-trigger",
    "aliyunregionid":"cn-chengdu",
    "aliyunpublishaddr":"42.120.XX.XX",
    "data":{
        "topic":"TopicName",
        "systemProperties":{
            "MIN_OFFSET":"0",
            "TRACE_ON":"true",
            "MAX_OFFSET":"8",
            "MSG_REGION":"cn-hangzhou",
            "KEYS":"systemProperties.KEYS",
            "CONSUME_START_TIME":1628577790396,
            "TAGS":"systemProperties.TAGS",
            "INSTANCE_ID":"MQ_INST_164901546557****_BXhFHryi"
        },
        "userProperties":{

        },
        "body":"TEST"
    }
}
```

### 事件流模型的 event 格式如下所示

```json
[
    {
    "id":"94ebc15f-f0db-4bbe-acce-56fb72fb****",
    "source":"RocketMQ-Function-rocketmq-trigger",
    "specversion":"1.0",
    "type":"mq:Topic:SendMessage",
    "datacontenttype":"application/json; charset=utf-8",
    "subject":"acs:mq:cn-hangzhou:164901546557****:MQ_INST_164901546557****_BXhFHryi%TopicName",
    "time":"2021-04-08T06:01:20.766Z",
    "aliyunaccountid":"164901546557****",
    "aliyunpublishtime":"2021-10-15T02:05:16.791Z",
    "aliyunoriginalaccountid":"164901546557****",
    "aliyuneventbusname":"RocketMQ-Function-rocketmq-trigger",
    "aliyunregionid":"cn-chengdu",
    "aliyunpublishaddr":"42.120.XX.XX",
    "data":{
        "topic":"TopicName",
        "systemProperties":{
            "MIN_OFFSET":"0",
            "TRACE_ON":"true",
            "MAX_OFFSET":"8",
            "MSG_REGION":"cn-hangzhou",
            "KEYS":"systemProperties.KEYS",
            "CONSUME_START_TIME":1628577790396,
            "TAGS":"systemProperties.TAGS",
            "INSTANCE_ID":"MQ_INST_164901546557****_BXhFHryi"
        },
        "userProperties":{

        },
        "body":"TEST"
    }
    },
    {
    "id":"94ebc15f-f0db-4bbe-acce-56fb72fb****",
    "source":"RocketMQ-Function-rocketmq-trigger",
    "specversion":"1.0",
    "type":"mq:Topic:SendMessage",
    "datacontenttype":"application/json; charset=utf-8",
    "subject":"acs:mq:cn-hangzhou:164901546557****:MQ_INST_164901546557****_BXhFHryi%TopicName",
    "time":"2021-04-08T06:01:20.766Z",
    "aliyunaccountid":"164901546557****",
    "aliyunpublishtime":"2021-10-15T02:05:16.791Z",
    "aliyunoriginalaccountid":"164901546557****",
    "aliyuneventbusname":"RocketMQ-Function-rocketmq-trigger",
    "aliyunregionid":"cn-chengdu",
    "aliyunpublishaddr":"42.120.XX.XX",
    "data":{
        "topic":"TopicName",
        "systemProperties":{
            "MIN_OFFSET":"0",
            "TRACE_ON":"true",
            "MAX_OFFSET":"8",
            "MSG_REGION":"cn-hangzhou",
            "KEYS":"systemProperties.KEYS",
            "CONSUME_START_TIME":1628577790396,
            "TAGS":"systemProperties.TAGS",
            "INSTANCE_ID":"MQ_INST_164901546557****_BXhFHryi"
        },
        "userProperties":{

        },
        "body":"TEST"
    }
    }
]
```

### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置 RocketMQ 触发器之后，每当你向 RocketMQ 发送任意消息之后，就会触发该函数，并返回消息内容。

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

func HandleRequest(ctx context.Context, event events.RocketMQEvent) (string, error) {
	fcCtx, _ := fccontext.FromContext(ctx)
	flog := fcCtx.GetLogger()
	mes, _ := json.Marshal(event)
	flog.Info("event:", string(mes))
	return fmt.Sprintf("Body:%s", event.Data.Body), nil
}

func main() {
	fc.Start(HandleRequest)
}
```

