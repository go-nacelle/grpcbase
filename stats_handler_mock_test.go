// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.
// This file was generated by robots at
// 2019-06-24T19:52:05-05:00
// using the command
// $ go-mockgen -f google.golang.org/grpc/stats -i Handler -o stats_handler_mock_test.go

package grpcbase

import (
	"context"
	stats "google.golang.org/grpc/stats"
	"sync"
)

// MockHandler is a mock implementation of the Handler interface (from the
// package google.golang.org/grpc/stats) used for unit testing.
type MockHandler struct {
	// HandleConnFunc is an instance of a mock function object controlling
	// the behavior of the method HandleConn.
	HandleConnFunc *HandlerHandleConnFunc
	// HandleRPCFunc is an instance of a mock function object controlling
	// the behavior of the method HandleRPC.
	HandleRPCFunc *HandlerHandleRPCFunc
	// TagConnFunc is an instance of a mock function object controlling the
	// behavior of the method TagConn.
	TagConnFunc *HandlerTagConnFunc
	// TagRPCFunc is an instance of a mock function object controlling the
	// behavior of the method TagRPC.
	TagRPCFunc *HandlerTagRPCFunc
}

// NewMockHandler creates a new mock of the Handler interface. All methods
// return zero values for all results, unless overwritten.
func NewMockHandler() *MockHandler {
	return &MockHandler{
		HandleConnFunc: &HandlerHandleConnFunc{
			defaultHook: func(context.Context, stats.ConnStats) {
				return
			},
		},
		HandleRPCFunc: &HandlerHandleRPCFunc{
			defaultHook: func(context.Context, stats.RPCStats) {
				return
			},
		},
		TagConnFunc: &HandlerTagConnFunc{
			defaultHook: func(context.Context, *stats.ConnTagInfo) context.Context {
				return nil
			},
		},
		TagRPCFunc: &HandlerTagRPCFunc{
			defaultHook: func(context.Context, *stats.RPCTagInfo) context.Context {
				return nil
			},
		},
	}
}

// NewMockHandlerFrom creates a new mock of the MockHandler interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockHandlerFrom(i stats.Handler) *MockHandler {
	return &MockHandler{
		HandleConnFunc: &HandlerHandleConnFunc{
			defaultHook: i.HandleConn,
		},
		HandleRPCFunc: &HandlerHandleRPCFunc{
			defaultHook: i.HandleRPC,
		},
		TagConnFunc: &HandlerTagConnFunc{
			defaultHook: i.TagConn,
		},
		TagRPCFunc: &HandlerTagRPCFunc{
			defaultHook: i.TagRPC,
		},
	}
}

// HandlerHandleConnFunc describes the behavior when the HandleConn method
// of the parent MockHandler instance is invoked.
type HandlerHandleConnFunc struct {
	defaultHook func(context.Context, stats.ConnStats)
	hooks       []func(context.Context, stats.ConnStats)
	history     []HandlerHandleConnFuncCall
	mutex       sync.Mutex
}

// HandleConn delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockHandler) HandleConn(v0 context.Context, v1 stats.ConnStats) {
	m.HandleConnFunc.nextHook()(v0, v1)
	m.HandleConnFunc.appendCall(HandlerHandleConnFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the HandleConn method of
// the parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerHandleConnFunc) SetDefaultHook(hook func(context.Context, stats.ConnStats)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// HandleConn method of the parent MockHandler instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *HandlerHandleConnFunc) PushHook(hook func(context.Context, stats.ConnStats)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerHandleConnFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(context.Context, stats.ConnStats) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerHandleConnFunc) PushReturn() {
	f.PushHook(func(context.Context, stats.ConnStats) {
		return
	})
}

func (f *HandlerHandleConnFunc) nextHook() func(context.Context, stats.ConnStats) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerHandleConnFunc) appendCall(r0 HandlerHandleConnFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerHandleConnFuncCall objects
// describing the invocations of this function.
func (f *HandlerHandleConnFunc) History() []HandlerHandleConnFuncCall {
	f.mutex.Lock()
	history := make([]HandlerHandleConnFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerHandleConnFuncCall is an object that describes an invocation of
// method HandleConn on an instance of MockHandler.
type HandlerHandleConnFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 stats.ConnStats
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerHandleConnFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerHandleConnFuncCall) Results() []interface{} {
	return []interface{}{}
}

// HandlerHandleRPCFunc describes the behavior when the HandleRPC method of
// the parent MockHandler instance is invoked.
type HandlerHandleRPCFunc struct {
	defaultHook func(context.Context, stats.RPCStats)
	hooks       []func(context.Context, stats.RPCStats)
	history     []HandlerHandleRPCFuncCall
	mutex       sync.Mutex
}

// HandleRPC delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandler) HandleRPC(v0 context.Context, v1 stats.RPCStats) {
	m.HandleRPCFunc.nextHook()(v0, v1)
	m.HandleRPCFunc.appendCall(HandlerHandleRPCFuncCall{v0, v1})
	return
}

// SetDefaultHook sets function that is called when the HandleRPC method of
// the parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerHandleRPCFunc) SetDefaultHook(hook func(context.Context, stats.RPCStats)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// HandleRPC method of the parent MockHandler instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *HandlerHandleRPCFunc) PushHook(hook func(context.Context, stats.RPCStats)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerHandleRPCFunc) SetDefaultReturn() {
	f.SetDefaultHook(func(context.Context, stats.RPCStats) {
		return
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerHandleRPCFunc) PushReturn() {
	f.PushHook(func(context.Context, stats.RPCStats) {
		return
	})
}

func (f *HandlerHandleRPCFunc) nextHook() func(context.Context, stats.RPCStats) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerHandleRPCFunc) appendCall(r0 HandlerHandleRPCFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerHandleRPCFuncCall objects describing
// the invocations of this function.
func (f *HandlerHandleRPCFunc) History() []HandlerHandleRPCFuncCall {
	f.mutex.Lock()
	history := make([]HandlerHandleRPCFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerHandleRPCFuncCall is an object that describes an invocation of
// method HandleRPC on an instance of MockHandler.
type HandlerHandleRPCFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 stats.RPCStats
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerHandleRPCFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerHandleRPCFuncCall) Results() []interface{} {
	return []interface{}{}
}

// HandlerTagConnFunc describes the behavior when the TagConn method of the
// parent MockHandler instance is invoked.
type HandlerTagConnFunc struct {
	defaultHook func(context.Context, *stats.ConnTagInfo) context.Context
	hooks       []func(context.Context, *stats.ConnTagInfo) context.Context
	history     []HandlerTagConnFuncCall
	mutex       sync.Mutex
}

// TagConn delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandler) TagConn(v0 context.Context, v1 *stats.ConnTagInfo) context.Context {
	r0 := m.TagConnFunc.nextHook()(v0, v1)
	m.TagConnFunc.appendCall(HandlerTagConnFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the TagConn method of
// the parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerTagConnFunc) SetDefaultHook(hook func(context.Context, *stats.ConnTagInfo) context.Context) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// TagConn method of the parent MockHandler instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *HandlerTagConnFunc) PushHook(hook func(context.Context, *stats.ConnTagInfo) context.Context) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerTagConnFunc) SetDefaultReturn(r0 context.Context) {
	f.SetDefaultHook(func(context.Context, *stats.ConnTagInfo) context.Context {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerTagConnFunc) PushReturn(r0 context.Context) {
	f.PushHook(func(context.Context, *stats.ConnTagInfo) context.Context {
		return r0
	})
}

func (f *HandlerTagConnFunc) nextHook() func(context.Context, *stats.ConnTagInfo) context.Context {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerTagConnFunc) appendCall(r0 HandlerTagConnFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerTagConnFuncCall objects describing
// the invocations of this function.
func (f *HandlerTagConnFunc) History() []HandlerTagConnFuncCall {
	f.mutex.Lock()
	history := make([]HandlerTagConnFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerTagConnFuncCall is an object that describes an invocation of
// method TagConn on an instance of MockHandler.
type HandlerTagConnFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *stats.ConnTagInfo
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 context.Context
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerTagConnFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerTagConnFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// HandlerTagRPCFunc describes the behavior when the TagRPC method of the
// parent MockHandler instance is invoked.
type HandlerTagRPCFunc struct {
	defaultHook func(context.Context, *stats.RPCTagInfo) context.Context
	hooks       []func(context.Context, *stats.RPCTagInfo) context.Context
	history     []HandlerTagRPCFuncCall
	mutex       sync.Mutex
}

// TagRPC delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockHandler) TagRPC(v0 context.Context, v1 *stats.RPCTagInfo) context.Context {
	r0 := m.TagRPCFunc.nextHook()(v0, v1)
	m.TagRPCFunc.appendCall(HandlerTagRPCFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the TagRPC method of the
// parent MockHandler instance is invoked and the hook queue is empty.
func (f *HandlerTagRPCFunc) SetDefaultHook(hook func(context.Context, *stats.RPCTagInfo) context.Context) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// TagRPC method of the parent MockHandler instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *HandlerTagRPCFunc) PushHook(hook func(context.Context, *stats.RPCTagInfo) context.Context) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *HandlerTagRPCFunc) SetDefaultReturn(r0 context.Context) {
	f.SetDefaultHook(func(context.Context, *stats.RPCTagInfo) context.Context {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *HandlerTagRPCFunc) PushReturn(r0 context.Context) {
	f.PushHook(func(context.Context, *stats.RPCTagInfo) context.Context {
		return r0
	})
}

func (f *HandlerTagRPCFunc) nextHook() func(context.Context, *stats.RPCTagInfo) context.Context {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *HandlerTagRPCFunc) appendCall(r0 HandlerTagRPCFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of HandlerTagRPCFuncCall objects describing
// the invocations of this function.
func (f *HandlerTagRPCFunc) History() []HandlerTagRPCFuncCall {
	f.mutex.Lock()
	history := make([]HandlerTagRPCFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// HandlerTagRPCFuncCall is an object that describes an invocation of method
// TagRPC on an instance of MockHandler.
type HandlerTagRPCFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *stats.RPCTagInfo
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 context.Context
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c HandlerTagRPCFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c HandlerTagRPCFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}