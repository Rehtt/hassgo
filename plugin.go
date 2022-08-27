package hass

type Plugins []Plugin

type Plugin interface {
	Init(ctx *Func)
	SubscribeEventCallBack(ctx *Func, data []byte)   // 订阅事件回调
	SubscribeTriggerCallBack(ctx *Func, data []byte) // 订阅触发器回调
}

func (ps *Plugins) Register(p ...Plugin) {
	*ps = append(*ps, p...)
}
