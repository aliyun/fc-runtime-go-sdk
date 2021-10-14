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
