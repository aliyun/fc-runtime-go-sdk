## OSS 触发器

### Event 格式

oss 触发器的 event 格式如下所示：

```json
{
    "events": [
      {
        "eventName": "ObjectCreated:PutObject",
        "eventSource": "acs:oss",
        "eventTime": "2022-06-10T06:21:35.000Z",
        "eventVersion": "1.0",
        "oss": {
          "bucket": {
            "arn": "acs:oss:cn-shenzhen:1581223139******:bucketname",
            "name": "testbucket",
            "ownerIdentity": "1581224182******",
            "virtualBucket": ""
          },
          "object": {
            "deltaSize": 29203,
            "eTag": "7D73E1052AF595535DE90994F4C3E6D5",
            "key": "image/a.jpg",
            "objectMeta": {
              "mimeType": "image/jpeg"
            },
            "size": 29203
          },
          "ossSchemaVersion": "1.0",
          "ruleId": "7c5a5ccfb06114a4d1ea8867e3916eb543d84e7c"
        },
        "region": "cn-shenzhen",
        "requestParameters": {
          "sourceIPAddress": "106.23.**.**"
        },
        "responseElements": {
          "requestId": "62A2E2EFD0DEFE3930847A6E"
        },
        "userIdentity": {
          "principalId": "1581224182******"
        }
      }
    ]
  }
```

### 使用示例

下面展示了一个简单的 FC demo，当你为该函数配置 oss 触发器之后，每当你在 oss 对象存储进行相关操作之后，oss 系统就会捕捉到对应事件，并将事件信息编码为 JSON 字符串，触发该函数并作为函数的输入，函数反馈给你 bucket 的名称及操作对象的大小信息。

```go
package main

import (
	"context"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/events"
)

func HandleRequest(ctx context.Context, event events.OssEvent) (string, error) {
	fmt.Printf("hello,the name of your bucket is %s", *event.Events[0].Oss.Bucket.Name)
	return fmt.Sprintf("hello,The size of the object you are manipulating is %dB",*event.Events[0].Oss.Object.Size), nil
}

func main() {
	fc.Start(HandleRequest)
}
```

