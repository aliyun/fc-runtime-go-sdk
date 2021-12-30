# Golang 入门

## 背景信息

函数计算目前支持以下 golang 1.x 版本，推荐使用 go1.8 及以上版本。



#### golang sdk 和工具

函数计算目前提供以下 golang sdk 和工具
- [FC SDK for Go](https://github.com/aliyun/FC-Runtime-Go-SDK/tree/master/fc): golang runtime 编程模型的具体实现，函数计算平台使用这个包运行您的handler
- [github.com/aliyun/FC-Runtime-Go-SDK/fccontext](https://github.com/aliyun/FC-Runtime-Go-SDK/tree/master/fccontext): 访问context信息的辅助库
- [github.com/aliyun/FC-Runtime-Go-SDK/examples](https://github.com/aliyun/FC-Runtime-Go-SDK/tree/master/examples): 使用golang runtime的简单示例

## 编译部署

Golang 语言不支持在线编辑，仅支持 .zip 方式，可以直接上传也可以指定oss路径。



#### 1. 在 Mac 或 Linux 下

1. 下载函数计算 golang sdk 库：

```
github.com/aliyun/fc-runtime-go-sdk
```

2. 编译并打包

```bash
GOOS=linux CGO_ENABLED=0 go build main.go
zip function.zip main
```

设置 GOOS=linux，确保编译后的可执行文件与函数计算平台的 Go 运行系统环境兼容，尤其是在非 Linux 环境中编译时。
在 Linux 下，建议使用纯静态编译，设置 CGO_ENABLED=0，确保可执行文件不依赖任何外部依赖库（如libc库），避免出现编译环境和 Go 运行时环境依赖库兼容问题。
```bash
GOOS=linux CGO_ENABLED=0 go build main.go
```

在 M1 Mac （或其他ARM架构的机器）下，还需要设置 GOARCH=amd64，实现跨平台交叉编译。
```bash
GOOS=linux GOARCH=amd64 go build main.go
```

如果你的包里有多个文件

```
GOOS=linux go build main
```
## 函数入口
Golang 是编译型语言，需要在本地编译后直接上传可执行的二进制文件，在函数入口配置中，不同于 Python，NodeJS的 `[文件名].[函数名]` 格式，Golang 语言的函数入口可直接设置为 `[文件名]`。
该文件名是只编译后的二进制文件名称，当函数被调用时，函数计算平台会直接执行函数入口配置的文件名。
比如，使用GOOS=linux CGO_ENABLED=0 go build main.go 编译出来的文件是 main， 那么在函数入口的配置里就可以填入 main 。

> 注意，这里 main 文件必须在zip包的顶层，可以使用 unzip -l your.zip 命令查看压缩包的目录结构。
> 如果想要调用压缩包中的 code/hello, 可以将函数入口设置为 code/hello

# 事件函数


在 Golang 的代码中，需要引入官方的 sdk 库 `github.com/aliyun/fc-runtime-go-sdk/fc`，并且，需要实现 `handler` 函数和 `main` 函数。

```golang
// hello_world.go
package main

import (
	"fmt"
    "context"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
    return fmt.Sprintf("你好，%s!", event.Name), nil
}

func main() {
	fc.Start(HandleRequest)
}
```
传入的 event 参数，一个包含 name 属性的 json 字符串

```json
{
  "name": "世界"
}
```


示例解析:
- **package main**: 在 Golang 语言中， Go 应用程序都包含一个名为 `main` 的包
- **import**: 需要引用函数计算依赖的包，主要包括
  - **github.com/aliyun/fc-runtime-go-sdk/fc**: 函数计算 Golang 语言的核心库
  - **context**：函数计算 Golang 语言的 Context 对象
- **func HandleRequest(ctx context.Context, event map[string]interface{}) (string, error)** :  这个是程序的入口函数，里面包含将要执行的代码，参数含义如下：
  - **ctx context.Context:** 提供了函数在调用时的运行信息，可以在 Handler 页面找到详细信息
  - **event MyEvent**: 调用函数时传入的数据，可以是多种类型，具体支持格式可以见 Handler 页面
  - **string, error**: 返回两个数据，字符串和错误信息
  - **return fmt.Sprintf("hello world! 你好，%s !", event.Name), nil**: 简单的返回 Hello 信息，其中包含传入的 event。 nil 表示没有错误发生。
- func main(): Golang 函数代码的入口，这个是必备的。

通过添加代码 `fc.Start(HandleRequest)`，你的程序就可以在阿里云的函数计算平台运行了。

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

    res, _ := json.Marshal(fctx)
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

```golang
package main

import (
	"context"
	"fmt"
	"net/http"
    "io/ioutil"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Header().Add("Content-Type", "text/plain")
    w.Write([]byte(err.Error()))
    return nil
  }
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("你好，%s!\n", body)))
	return nil
}

func main() {
	fc.StartHttp(HandleHttpRequest)
}
```

示例解析（有更新）：

- **package main**: 在 Golang 语言中， Go 应用程序都包含一个名为 `main` 的包
- **import**: 需要引用函数计算依赖的包，主要包括
  - **github.com/aliyun/fc-runtime-go-sdk/fc**: 函数计算 Golang 语言的核心库
  - **context**：函数计算 Golang 语言的 Context 对象
  - **net/http**： HTTP 函数中需要用到的 http 包中的 Request 和 ResponseWriter 接口
- **HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req \*http.Request) error**:  这个是程序的入口函数，里面包含将要执行的代码，参数含义如下：
  - **ctx context.Context:** 提供了函数在调用时的运行信息，可以在 Handler 页面找到详细信息
  - **w http.ResponseWriter**: HTTP 函数的响应（responses）接口，可以设置状态行(status code)，消息报头(headers) 和 响应正文(body)，具体支持格式可以见 xxx 页面（根据实际情况填写）
  - **req \*http.Request**: HTTP 函数的请求（request) 接口，包含请求行（request line），请求头部（header）和请求数据（body），具体方法见xxx（根据实际情况填写）
  - **w.WriteHeader(http.StatusOK)**: 填入响应的 HTTP 状态码
  - **w.Header().Add("Content-Type", "text/plain")**： 填入响应的消息报头(headers)
  - **w.Write([]byte(fmt.Sprintf("你好，%s!\n", body)))**： 填入响应的消息体（body）
  - **return nil**: 简单的错误信息，nil 表示没有错误发生，如果设置了错误信息，则认为是函数错误，具体方法见 Error（根据实际情况填写）
- func main(): Golang 函数代码的入口，这个是必备的。

是通过 `fc.StartHttp(HandleHttpRequest)`，你的程序就可以在阿里云的函数计算平台运行了。

注意： http 函数和事件函数调用的方法不同，事件函数： `fc.Start()`， http 函数： `fc.StartHttp` 。

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
    w.WriteHeader(http.StatusOK)
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