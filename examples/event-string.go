package main

import (
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.Start(HandleRequest)
}

// event: "hello world"
func HandleRequest(ctx context.Context, event string) (string, error) {
	fmt.Printf("event: %v\n", event)
	fmt.Println("hello world! 你好，世界!")
	return "hello world! 你好，世界!", nil
}
