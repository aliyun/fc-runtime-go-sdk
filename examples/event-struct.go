package main

import (
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type People struct {
	Name string
	Age  int
}

func main() {
	fc.Start(HandleRequest)
}

/*
{
    "Name": "hello world",
    "Age": 100
}
*/
func HandleRequest(event People) (string, error) {
	fmt.Printf("event: %v\n", event)
	fmt.Println("hello world! 你好，世界!")
	return "hello world! 你好，世界!", nil
}
