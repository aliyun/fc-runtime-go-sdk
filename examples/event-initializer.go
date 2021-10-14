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
