package hassgo

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type WS struct {
	token string
	conn  *websocket.Conn
	context.Context
	cancel  context.CancelFunc
	isLogin bool
	index   *uint64
	Plugins Plugins
}
type Context struct {
	ws *WS
	context.Context
	value   bytes.Buffer
	id      uint64
	plugin  Plugin
	cancel  context.CancelFunc
	subType uint8 // 订阅类型
}

var (
	contextMap  = make(map[uint64]*Context)
	lock        sync.RWMutex
	contextPool = sync.Pool{
		New: func() any {
			return new(Context)
		},
	}
)

func (w *WS) newContext() (ctx *Context) {
	lock.Lock()
	defer lock.Unlock()
	ctx = contextPool.Get().(*Context)
	ctx.ws = w
	ctx.Context, ctx.cancel = context.WithCancel(w.Context)

	atomic.AddUint64(w.index, 1)
	if atomic.LoadUint64(w.index) == 0 {
		atomic.AddUint64(w.index, 1)
	}
	ctx.id = atomic.LoadUint64(w.index)
	contextMap[ctx.id] = ctx
	return
}
func (w *WS) getContext(id uint64) *Context {
	lock.RLock()
	ctx, ok := contextMap[id]
	lock.RUnlock()
	if ok {
		return ctx
	}

	return nil
}
func (ctx *Context) destroyContext() {
	lock.Lock()
	defer lock.Unlock()
	delete(contextMap, ctx.id)
	contextPool.Put(ctx)
}

func Connect(wsUrl, origin, token string) (ws *WS, err error) {
	ws = new(WS)
	ws.Plugins = make(Plugins, 0, 5)
	ws.conn, err = websocket.Dial(wsUrl, "", origin)
	ws.Context, ws.cancel = context.WithCancel(context.Background())
	ws.index = new(uint64)
	ws.token = token
	return
}

func (w *WS) Run() {
	// ping pong
	go func() {
		f := &Func{ws: w}
		for {
		Sleep:
			time.Sleep(30 * time.Second)
			if !w.isLogin {
				goto Sleep
			}
			var tmp map[string]any

			_, err := f.Ping(&tmp)
			if err != nil || tmp["type"] != "pong" {
				log.Println("连接中断 1")
				w.cancel()
			}
		}
	}()
	for {
		select {
		case <-w.Done():
			return
		default:
			break
		}
		if err := w.readAll(); err != nil && !(w.conn.IsServerConn() || w.conn.IsClientConn()) {
			log.Println("连接中断 2")
			w.cancel()
		}
	}
}

func (w *WS) readAll() error {
	var tmp = make([]byte, 512)
	var value bytes.Buffer
	for {
		n, err := w.conn.Read(tmp)
		if err != nil {
			return err
		}
		value.Write(tmp[:n])
		if err = w.handle(&value); err == nil {
			return nil
		}

	}
}

func (ctx *Context) Value() []byte {
	return ctx.value.Bytes()
}
func (ctx *Context) Write(b []byte) (n int, err error) {
	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("当前Context已关闭,id: %d", ctx.id)
	default:
		return ctx.ws.conn.Write(b)
	}

}
func (ctx *Context) Read(b []byte) (n int, err error) {
	select {
	case <-ctx.Done():
		return ctx.value.Read(b)
	}
}
func (ctx *Context) ReadAll() []byte {
	<-ctx.Done()
	return ctx.value.Bytes()
}

func (ctx *Context) Close() error {
	err := ctx.Err()
	ctx.ws.cancel()
	return err
}

func (ctx *Context) GetContextID() uint64 {
	return ctx.id
}
