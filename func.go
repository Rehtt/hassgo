package hass

import (
	"encoding/json"
)

type Func struct {
	ws *WS
}

const (
	// 订阅类型
	sub_event = iota + 1
	sub_trigger
)

func (f *Func) GetConfig() (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":   id,
			"type": "get_config",
		})
		return
	})
	return
}

// SubscribeEvents 订阅事件
func (f *Func) SubscribeEvents(eventType string, plugin Plugin) (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":   id,
			"type": "subscribe_events",
		})
		return data
	},
		// 订阅
		func() (subType uint8, p Plugin) {
			return sub_event, plugin
		})
	return
}

// SubscribeTrigger 订阅触发器
func (f *Func) SubscribeTrigger(trigger any, plugin Plugin) (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":      id,
			"type":    "subscribe_trigger",
			"trigger": trigger,
		})
		return data
	}, func() (subType uint8, p Plugin) {
		return sub_trigger, plugin
	})
	return
}

// Unsubscribe 取消订阅
func (f *Func) Unsubscribe(subscription uint64) (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":           id,
			"type":         "unsubscribe_events",
			"subscription": subscription,
		})
		return
	})
	return
}

// CallService 调用服务
func (f *Func) CallService(v any) (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(v)
		var tmp map[string]any
		json.Unmarshal(data, &tmp)
		tmp["id"] = id
		tmp["type"] = "call_service"
		data, _ = json.Marshal(tmp)
		return data
	})
	return
}

func (f *Func) GetService() (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":   id,
			"type": "get_services",
		})
		return
	})
	return
}

func (f *Func) GetStates() (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":   id,
			"type": "get_states",
		})
		return
	})
	return
}

// Ping ping
func (f *Func) Ping(v ...any) (out []byte, err error) {
	out, err = f.ws.sendHandle(func(id uint64) (data []byte) {
		data, _ = json.Marshal(map[string]any{
			"id":   id,
			"type": "ping",
		})
		return data
	})
	if err != nil {
		return nil, err
	}
	if len(v) > 0 {
		err = json.Unmarshal(out, v[0])
	}
	return
}
