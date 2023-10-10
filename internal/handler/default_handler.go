package handler

import (
	"errors"
	"sync"

	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

func init() {
	extension.SetRejectedExecutionHandler("DefaultValueHandler", GetDefaultValueRejectedExecutionHandlerSingleton)
}

type DefaultValueRejectedExecutionHandler struct {
	defaultResult sync.Map
}

func (mh *DefaultValueRejectedExecutionHandler) RejectedExecution(url *common.URL, invocation protocol.Invocation) protocol.Result {
	key := url.ServiceKey() + "#" + invocation.MethodName()
	result, loaded := mh.defaultResult.Load(key)
	if !loaded {
		// we didn't configure any default value for this invocation
		return &protocol.RPCResult{
			Err: errors.New("请求太频繁"),
		}
	}
	return result.(*protocol.RPCResult)
}

func GetCustomRejectedExecutionHandler() filter.RejectedExecutionHandler {
	return &DefaultValueRejectedExecutionHandler{}
}

var (
	customHandlerOnce     sync.Once
	customHandlerInstance *DefaultValueRejectedExecutionHandler
)

/**
 * the better way is designing the RejectedExecutionHandler as singleton.
 */
func GetDefaultValueRejectedExecutionHandlerSingleton() filter.RejectedExecutionHandler {
	customHandlerOnce.Do(func() {
		customHandlerInstance = &DefaultValueRejectedExecutionHandler{}
	})

	initDefaultValue()

	return customHandlerInstance
}

func initDefaultValue() {
	// setting your default value
}
