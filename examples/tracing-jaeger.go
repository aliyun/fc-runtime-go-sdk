package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport"
)

type MyEvent struct {
	Name string `json:"name"`
}

func NewJaegerTracer(service, endpoint string) (opentracing.Tracer, io.Closer) {
	sender := transport.NewHTTPTransport(endpoint)
	tracer, closer := jaeger.NewTracer(service,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender))
	return tracer, closer
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fctx, _ := fccontext.FromContext(ctx)
	fmt.Printf("context: %#v\n", fctx)
	fmt.Println("hello world")

	// New tracer
	tracer, closer := NewJaegerTracer("FCTracer", fctx.Tracing.JaegerEndpoint)
	defer closer.Close()

	// retrieve spanContext
	spanContext, err := jaeger.ContextFromString(fctx.Tracing.OpenTracingSpanContext)
	if err != nil {
		return "", fmt.Errorf("OpenTracingSpanContext: %s, error: %v", fctx.Tracing.OpenTracingSpanContext, err)
	}

	// span start/finish
	span := tracer.StartSpan("MyFCSpan", opentracing.ChildOf(spanContext))
	for k, v := range fctx.Tracing.OpenTracingSpanBaggages {
		span.SetBaggageItem(k, v)
	}
	span.SetOperationName("fc-operation")
	span.SetTag("version", "fc-v1")
	time.Sleep(100 * time.Millisecond)
	span.LogFields(
		log.String("event", "soft error"),
		log.String("type", "cache timeout"),
		log.Int("waited.millis", 100))
	time.Sleep(100 * time.Millisecond)
	span.Finish()

	return fmt.Sprintf("hello world! 你好，%s!", event.Name), nil
}

func main() {
	fc.Start(HandleRequest)
}
