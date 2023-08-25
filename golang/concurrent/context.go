package concurrent

import (
	"sync"
	"time"
)

/**
context 主要用在goroutine之间传递取消信号、超时信号、截止时间以及一些共享的值

Context	        接口	定义了 Context 接口的四个方法
emptyCtx		结构体	实现了 Context 接口，它其实是个空的 context
CancelFunc		函数	取消函数
canceler		接口	context 取消接口，定义了两个方法
cancelCtx		结构体	可以被取消
timerCtx		结构体	超时会被取消
valueCtx	    结构体	可以存储 k-v 对
Background  	函数	返回一个空的 context，常作为根 context
WithCancel	    函数	基于父 context，生成一个可以取消的 context
newCancelCtx	函数	创建一个可取消的 context
propagateCancel	函数	向下传递 context 节点间的取消关系
parentCancelCtx	函数	找到第一个可取消的父节点
removeChild	    函数	去掉父节点的孩子节点
init	        函数	包初始化
WithDeadline	函数	创建一个有 deadline 的 context
WithTimeout   	函数	创建一个有 timeout 的 context
WithValue	    函数	创建一个存储 k-v 对的 context
*/

// A Context carries a deadline, a cancellation signal, and other values across
// API boundaries.
//
type Context interface {
	Deadline() (deadline time.Time, ok bool) //返回context取消时间 如果ok==false则该context没有取消时间
	Done() <-chan struct{}                   //https://blog.csdn.net/weixin_42216109/article/details/123694275  cancel机制取消上下文  搭配context.Cancel 取消当前上下文
	// if Done is not yet closed ,Err return nil .
	// if Done is closed , Err returns a non-nil error message explaining why
	Err() error
	Value(key interface{}) interface{}
}

//An emptyContext is never cancel , has no values , and has no deadline. Is is not struct{}, since vars of this type must have distinct addresses.
type emptyContext int

func (*emptyContext) Deadline() (deadline time.Time, ok bool) {
	return
}
func (*emptyContext) Done() <-chan struct{} {
	return nil
}
func (*emptyContext) Err() error {
	return nil
}
func (*emptyContext) Value(key interface{}) interface{} {
	return nil
}

var background = new(emptyContext)

//A canceler is a context type that can be  canceled directly . The implementations is *cancelContext and *timeContext
type canceler interface {
	cancel(removeFromParent bool, err error)
	Done() <-chan struct{}
}

// A cancelContext can be canceled . When canceled, it is also cancels any children context
// cancelContext is implement canceler
type cancelContext struct {
	Context
	//todo want to know  which struct implement deadline method
	mu       sync.Mutex            //protects following fields
	done     chan struct{}         //create lazily , closed by first cancel call
	children map[canceler]struct{} //set to nil by the first cancel call
	err      error                 //set to non-nil by the first cancel call
}

// closedchan is a reusable closed channel. 是一个可重复复用的channel
var closedchan = make(chan struct{})

func init() {
	close(closedchan)
}

func (ctx *cancelContext) Done() <-chan struct{} {
	ctx.mu.Lock()
	if ctx.done == nil {
		ctx.done = make(chan struct{})
	}
	d := ctx.done
	ctx.mu.Unlock()
	return d //这里边为什么不直接返回 ctx.done 呢
}
func (ctx *cancelContext) Value(key interface{}) interface{} {
	return ctx.Context.Value(key) //找到context接口实现函数呗
}
func (ctx *cancelContext) Err() error {
	ctx.mu.Lock()
	err := ctx.err
	ctx.mu.Unlock()
	return err
}

//cancel closes ctx.done , cancels each of c's children ,
//and if removeFromParent is true, remove c from its parent's children
func (ctx *cancelContext) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	ctx.mu.Lock()
	if ctx.err != nil {
		ctx.mu.Unlock()
		return //already cancel
	}
	ctx.err = err
	if ctx.done == nil {
		ctx.done = closedchan
	} else {
		close(ctx.done)
	}

	for child := range ctx.children {
		//NOTE:获取父锁 但同时也获取 孩子锁
		child.cancel(false, err)
	}
	ctx.children = nil
	ctx.mu.Unlock()

	if removeFromParent {
		//removeChild(ctx.Context , ctx )
	}
}

//A CancelFunc
type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {

	return nil, nil
}

func Background() Context {
	return background
}
