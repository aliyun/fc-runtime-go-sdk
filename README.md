# 阿里云函数计算 Golang Runtime SDK

## 基本介绍

本项目是阿里云函数计算服务的 Golang 运行时的 SDK，包括使用 Golang 语言开发函数计算程序的依赖库、示例代码和一些工具。

详情请参考阿里云函数计算的官方文档：[阿里云函数计算 Golang 代码开发](https://help.aliyun.com/document_detail/323505.html)

> 注意: Golang 运行时与 Custom-Runtime(Golang)运行时是两个不同的运行时，本项目只适用于 Golang 运行时，不可以在 Custom-Runtime 运行时使用。

本项目目前主要由以下几部分组成
- [FC SDK for Go](https://github.com/aliyun/FC-Runtime-Go-SDK/tree/master/fc): golang runtime 编程模型的具体实现，函数计算平台使用这个包运行您的handler
- [github.com/aliyun/fc-runtime-go-sdk/fccontext](https://github.com/aliyun/fc-runtime-go-sdk/tree/master/fccontext): 访问context信息的辅助库
- [github.com/aliyun/fc-runtime-go-sdk/examples](https://github.com/aliyun/fc-runtime-go-sdk/tree/master/examples): 使用golang runtime的简单示例
- [github.com/aliyun/fc-runtime-go-sdk/events](https://github.com/aliyun/fc-runtime-go-sdk/tree/master/events): 常用触发器的 event 格式和示例代码。
- [github.com/aliyun/fc-runtime-go-sdk/cmd/build-fc-zip](https://github.com/aliyun/fc-runtime-go-sdk/tree/master/cmd/build-fc-zip): 在 windows 环境中的打包工具。（在windows环境使用zip命令打包的二进制文件缺少可执行权限）

## 快速开始
```golang
package main

import (
    "fmt"
    "context"

    "github.com/aliyun/fc-runtime-go-sdk/fc"
)

type StructEvent struct {
    Key string `json:"key"`
}

func HandleRequest(ctx context.Context, event StructEvent) (string, error) {
    return fmt.Sprintf("hello, %s!", event.Key), nil
}

func main() {
    fc.Start(HandleRequest)
}
```

## 编译打包

Golang 语言不支持在线编辑，仅支持 .zip 方式，可以直接上传也可以指定oss路径。

#### 1. 在 Mac 或 Linux 下编译打包

1. 下载函数计算 golang sdk 库：

```
go get github.com/aliyun/fc-runtime-go-sdk/fc
```

2. 编译并打包

```bash
GOOS=linux CGO_ENABLED=0 go build -o main main.go
zip function.zip main
```

设置GOOS=linux，确保编译后的可执行文件与函数计算平台的Go运行系统环境兼容，尤其是在非Linux环境中编译时。

补充说明如下：
- 针对Linux操作系统，建议使用纯静态编译，配置CGO_ENABLED=0，确保可执行文件不依赖任何外部依赖库（如libc库），避免出现编译环境和Go运行时环境依赖库的兼容问题。示例如下：
  ```bash
  GOOS=linux CGO_ENABLED=0 go build -o main main.go
  ```

- 针对M1 macOS（或其他ARM架构的机器），配置GOARCH=amd64，实现跨平台编译，示例如下：
  ```bash
  GOOS=linux GOARCH=amd64 go build -o main main.go
  ```
- 如果你的包里有多个文件
  ```bash
  GOOS=linux go build main
  ```

#### 2. 在 Windows 下编译打包
在windows下使用zip工具打包后，zip包中的二进制程序会缺少可执行权限，您可以使用 build-fc-zip 工具进行打包。
1. 下载 build-fc-zip 工具
```bash
go.exe get -u github.com/aliyun/fc-runtime-go-sdk/cmd/build-fc-zip
```

2. 使用 `build-fc-zip` 工具创建 `.zip` 文件。如果你默认安装了 Go，工具默认会在 `%USERPROFILE%\Go\bin.`

```bash
set GOOS=linux
go build -o main main.go
%USERPROFILE%\Go\bin\build-fc-zip.exe -output main.zip main
```

## 部署
部署方法请参考官方文档 [编译部署代码包](https://help.aliyun.com/document_detail/418490.html) 。