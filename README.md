# Golang 入门

Golang runtime 目前在内测阶段

## 背景信息

函数计算目前支持以下 golang 1.x 版本，推荐使用 go1.8 及以上版本。



#### golang sdk 和工具

- golang runtime sdk : github.com/aliyun/fc-runtime-go-sdk

## 编译部署

Golang 语言不支持在线编辑，仅支持 .zip 方式，可以直接上传也可以指定oss路径。



#### 1. 在 Mac 或 Linux 下

下载函数计算 golang sdk 库：

```
github.com/aliyun/fc-runtime-go-sdk
```

编译并打包

```
GOOS=linux go build main.go
zip function.zip main
```

如果你的包里有多个文件

```
GOOS=linux go build main
```

在 Linux 下，可能需要设置 `**CGO_ENABLED=0**`

```
GOOS=linux CGO_ENABLED=0 go build main.go
```

#### 2. 在 Windows 下

... 待完善



# 事件函数

Golang 是编译型语言，需要在本地编译后直接上传可执行的二进制文件，在函数入口配置中，不同于 Python，NodeJS的 `[文件名].[函数名]` 格式，Golang 语言的函数入口可直接设置为 `[文件名]` ，该文件名是只编译后的二进制文件名称，当函数被调用时，函数计算平台会直接执行函数入口配置的文件名。



在 Golang 的代码中，需要引入官方的 sdk 库 `github.com/aliyun/fc-runtime-go-sdk/fc`，并且，需要实现 `handler` 函数和 `main` 函数。

```
// hello_world.go
package main

import (
    "fmt"
    
    "github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
    fc.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event string) (string, error) {
    fmt.Println("hello world")
    return "hello world", nil
}
```

## Handler

**func HandleRequest(ctx context.Context, event string) (string, error)**

handler 函数里包含了要执行的代码，主要包括：

- ctx contrext.Context: 提供了运行时的信息，主要有 `RequestID`, `Function`,`Service`等，具体信息见：github.com/aliyun/fc-runtime-go-sdk/fccontext。必须放到第一个参数位置。
- event string：事件信息，可以支持 Golang 基本类型，以及 `struct` 类型。只能放到第二个参数位置

- (string, error)：返回两个值，字符串和错误信息，其中第一个值支持 Golang 基本类型，以及 `struct` 类型



Handler 支持的格式如下：

```
 func ()
 func () error
 func (TIn) error
 func () (TOut, error)
 func (TIn) (TOut, error)
 func (context.Context) error
 func (context.Context, TIn) error
 func (context.Context) (TOut, error)
 func (context.Context, TIn) (TOut, error)
```

其中 TIn 和 TOut 与 `encoding/json` 标准库兼容。



## Context



context 提供了以下参数



#### 变量

- RequestID
- Credentials - 证书信息，包含 access key, access key secret 和 security token

- Function - 函数的配置信息
- Service - 服务的配置信息

- Region - 地域
- AccountId - 账号信息

详细信息见：github.com/aliyun/fc-runtime-go-sdk/fccontext



#### 方法

- Deadline - 返回函数执行的超时时间，格式为 Unix 时间戳，单位是毫秒。



### 使用示例

首先，函数的handler需要包含context参数，我们会把上述的变量信息插入到context的Value中。

然后，需要 `import github.com/aliyun/fc-runtime-go-sdk/fccontext`, 通过 `fccontext.FromContext`方法获取 `fccontext`。

```
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/aliyun/fc-runtime-go-sdk/fc"
    "github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

func main() {
    fc.Start(echoContext)
}

func echoContext(ctx context.Context) (string, error) {
    fctx, _ := fccontext.FromContext(ctx)
    log.Println(fctx.AccountId)
    log.Printf("%#v\n", fctx)
   
    return string(res), nil
}
```

下面的示例展示了如何使用 `deadline` 获取函数剩余执行时间。

```
package main

import (
	"context"
	"fmt"
	"log"
	"time"
    
    "github.com/aliyun/fc-runtime-go-sdk/fc"
)

func LongRunningHandler(ctx context.Context) (string, error) {
	deadline, _ := ctx.Deadline()
	fmt.Printf("now: %s\ndeadline: %s\n", time.Now().String(), deadline.String())
	deadline = deadline.Add(-100 * time.Millisecond)
	timeoutChannel := time.After(time.Until(deadline))

	for {

		select {

		case <-timeoutChannel:
			return "Finished before timing out.", nil

		default:
			log.Print("hello!")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	fc.Start(LongRunningHandler)
}
```



# HTTP 函数

## HTTP 函数定义

Golang 的 HTTP 函数定义参考 Golang 标准库 http 中的 [Handler interface](https://pkg.go.dev/net/http#Handler)，并新增一个参数 context。



```
function(ctx context.Context, w http.ResponseWriter, req *http.Request) error
```

函数定义中包含以下三部分内容：

- context - 和事件函数中的 context 一样
- http.Request - 请求结构体 

- http.ResponseWriter - 响应接口 



### 请求结构体

http.Request 是Golang 标准库 http 中的定义，目前支持的参数和方法

| 参数          | 类型          | 描述                  |
| ------------- | ------------- | --------------------- |
| Method        | string        | HTTP 方法             |
| URL           | *url.URL      | 请求地址信息          |
| Header        | http.Header   | HTTP 请求头部的键值对 |
| Body          | io.ReadCloser | 请求结构体            |
| ContentLength | int64         | 请求结构体数据长度    |



### 响应接口

实现了 http.ResponseWriter 声明的三个方法

```
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}
```

- WriteHeader(statusCode int) - 设置状态码
- Header() Header - 获取并设置响应头信息

- Write([]byte) (int, error) 设置响应体



### 限制说明

和其他语言一样，例如nodejs：https://help.aliyun.com/document_detail/74757.html?spm=a2c4g.11186623.6.577.38c877f1mFUxxt

![img](https://intranetproxy.alipay.com/skylark/lark/0/2021/png/18456742/1628491959069-4d4825d5-7ac4-4bfe-a936-3c2f86cd27dd.png)

## 使用示例

```
package main

import (
    "context"
    "net/http"

    "github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
    fc.StartHttp(HandleHttpRequest)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
    w.WriteHeader(http.Ok)
    w.Header().Add("Content-Type", "text/plain")
    w.Write([]byte("hello, world!\n"))
    return nil
}
```





# LifeCycle 函数



## Initializer 函数

Initializer函数是实例的初始化函数，保证同一实例**成功且仅成功执行一次**。本文介绍 golang Initializer 函数的结构和特点。

备注： 成功且仅成功一次的含义，是存在执行两次的情况的，比如第一次失败，第二次成功，但只会重试一次。



函数定义，只有一个 context 参数，使用方法和事件函数一样

```
function(ctx context.Context)
```



使用 Initializer 需要两步

1. 在代码中通过 `fc.RegistryInitializerFunction(Init)` 注册一个 Initializer 函数。
2. 在函数配置中打开 Initializer function 开关



示例

```
package main

import (
    "context"
    "log"

    "github.com/aliyun/fc-runtime-go-sdk/fc"
)

var (
    count int = 1
)

func Init(ctx context.Context) {
    count += 1000
}

func main() {
    fc.RegisterInitializerFunction(Init)
    fc.Start(HandleRequest)
}

func HandleRequest() (int, error) {
    count += 1
    log.Println("count: ", count)
    return count, nil
}
```