// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved
// Copyright 2021 Alibaba Group Holding Limited. All Rights Reserved.

package fc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc/messages"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
)

const (
	msPerS  = int64(time.Second / time.Millisecond)
	nsPerMS = int64(time.Millisecond / time.Nanosecond)
)

type handlerWrapper struct {
	handler  interface{}
	funcType functionType
}

// startRuntimeAPILoop will return an error if handling a particular invoke resulted in a non-recoverable error
// func startRuntimeAPILoop(ctx context.Context, api string, handler interface{}, funcType functionType) error {
func startRuntimeAPILoop(ctx context.Context, api string, baseHandler handlerWrapper, lifeCycleHandlers []handlerWrapper) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("%v", r)
		}
	}()
	client := newRuntimeAPIClient(api)
	function := NewFunction(baseHandler.handler, baseHandler.funcType).withContext(ctx)
	function.RegistryLifeCycleHandler(lifeCycleHandlers)
	for {
		invoke, err := client.next()
		if err != nil {
			logPrintf("failed to get invoke request due to %v", err)
			continue
		}
		err = handleInvoke(invoke, function)
		if err != nil {
			logPrintf("failed to invoke function due to %v", err)
		}
	}
}

// handleInvoke returns an error if the function panics, or some other non-recoverable error occurred
func handleInvoke(invoke *invoke, function *Function) error {
	functionRequest, err := convertInvokeRequest(invoke)
	if err != nil {
		return fmt.Errorf("unexpected error occurred when parsing the invoke: %v", err)
	}

	functionResponse := &messages.InvokeResponse{}
	ivkErr := function.Invoke(functionRequest, functionResponse, convertInvokeFunctionType(invoke))
	if functionResponse.Error != nil {
		payload := safeMarshal(functionResponse.Error)
		if err := invoke.failure(payload, contentTypeJSON); err != nil {
			return fmt.Errorf("unexpected error occurred when sending the function error to the API: %v", err)
		}
		if functionResponse.Error.ShouldExit {
			return fmt.Errorf("calling the handler function resulted in a panic")
		}
		return ivkErr
	}
	if ivkErr != nil {
		return ivkErr
	}

	if err := invoke.success(functionResponse.Payload, contentTypeJSON); err != nil {
		return fmt.Errorf("unexpected error occurred when sending the function functionResponse to the API: %v", err)
	}

	return nil
}

func convertInvokeFunctionType(invoke *invoke) functionType {
	funcType, err := strconv.ParseInt(invoke.headers.Get(headerFunctionType), 10, 64)
	if err != nil {
		return handleFunction
	}
	switch funcType {
	case int64(initializerFunction):
		return initializerFunction
	case int64(preFreezeFunction):
		return preFreezeFunction
	case int64(preStopFunction):
		return preStopFunction
	default:
		return handleFunction
	}

}

// convertInvokeRequest converts an invoke from the Runtime API, and unpacks it to be compatible with the shape of a `lambda.Function` InvokeRequest.
func convertInvokeRequest(invoke *invoke) (*messages.InvokeRequest, error) {
	deadlineEpochMS, err := strconv.ParseInt(invoke.headers.Get(headerDeadlineMS), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse contents of header: %s", headerDeadlineMS)
	}
	deadlineS := deadlineEpochMS / msPerS
	deadlineNS := (deadlineEpochMS % msPerS) * nsPerMS

	functionTimeoutSec, err := strconv.Atoi(invoke.headers.Get(headerFunctionTimeout))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contents of header: %s", headerFunctionTimeout)
	}

	retryCount := 0
	if retryCountStr := invoke.headers.Get(headerRetryCount); retryCountStr != "" {
		retryCount, err = strconv.Atoi(retryCountStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contents of header: %s", headerFunctionTimeout)
		}
	}

	spanBaggages := make(map[string]string)
	if base64SpanBaggages := invoke.headers.Get(headerOpenTracingSpanBaggages); base64SpanBaggages != "" {
		spanBaggagesByte, err := base64.StdEncoding.DecodeString(base64SpanBaggages)
		if err != nil {
			return nil, fmt.Errorf("failed to parse contents of header %s: %s", headerOpenTracingSpanContext, base64SpanBaggages)
		}
		if err := json.Unmarshal(spanBaggagesByte, &spanBaggages); err != nil {
			return nil, fmt.Errorf("failed to parse contents of header %s: %s", headerOpenTracingSpanContext, base64SpanBaggages)
		}
	}

	res := &messages.InvokeRequest{
		RequestId: invoke.id,
		Deadline: messages.InvokeRequest_Timestamp{
			Seconds: deadlineS,
			Nanos:   deadlineNS,
		},
		Payload: invoke.payload,
		Context: fccontext.FcContext{
			RequestID: invoke.id,
			Credentials: fccontext.Credentials{
				AccessKeyId:     invoke.headers.Get(headerAccessKeyId),
				AccessKeySecret: invoke.headers.Get(headerAccessKeySecret),
				SecurityToken:   invoke.headers.Get(headerSecurityToken),
			},
			Function: fccontext.Function{
				Name:    invoke.headers.Get(headerFunctionName),
				Handler: invoke.headers.Get(headerFunctionHandler),
				Memory:  invoke.headers.Get(headerFunctionMemory),
				Timeout: functionTimeoutSec,
			},
			Service: fccontext.Service{
				Name:       invoke.headers.Get(headerServiceName),
				LogProject: invoke.headers.Get(headerServiceLogproject),
				LogStore:   invoke.headers.Get(headerServiceLogstore),
				Qualifier:  invoke.headers.Get(headerQualifier),
				VersionId:  invoke.headers.Get(headerVersionId),
			},
			Tracing: fccontext.Tracing{
				OpenTracingSpanContext:  invoke.headers.Get(headerOpenTracingSpanContext),
				OpenTracingSpanBaggages: spanBaggages,
				JaegerEndpoint:          invoke.headers.Get(headerJaegerEndpoint),
			},
			Region:     invoke.headers.Get(headerRegion),
			AccountId:  invoke.headers.Get(headerAccountId),
			RetryCount: retryCount,
		},
	}

	if httpParams := invoke.headers.Get(headerHttpParams); httpParams != "" {
		res.HttpParams = &httpParams
	}

	return res, nil
}

func safeMarshal(v interface{}) []byte {
	payload, err := json.Marshal(v)
	if err != nil {
		v := &messages.InvokeResponse_Error{
			Type:    "Runtime.SerializationError",
			Message: err.Error(),
		}
		payload, err := json.Marshal(v)
		if err != nil {
			panic(err) // never reach
		}
		return payload
	}
	return payload
}
