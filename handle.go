package hassgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hassgo/model"
	"log"
)

func (w *WS) handle(value *bytes.Buffer) error {
	var response model.Head
	if err := json.Unmarshal(value.Bytes(), &response); err != nil {
		return err
	}

	if !w.isLogin {
		w.login(value)
		return nil
	}
	ctx := w.getContext(response.ID)
	if ctx == nil {
		err := fmt.Errorf("not find ID:%d", response.ID)
		log.Println(err)
		return err
	}
	ctx.value.Reset()
	ctx.value.ReadFrom(value)

	if response.Type == "event" {
		// todo
		if ctx.plugin == nil {
			return nil
		}
		switch ctx.subType {
		case sub_event:
			go ctx.plugin.SubscribeEventCallBack(&Func{ws: w}, ctx.value.Bytes())
		case sub_trigger:
			go ctx.plugin.SubscribeTriggerCallBack(&Func{ws: w}, ctx.value.Bytes())
		}
	} else {
		ctx.cancel() // 非事件返回一律关闭
	}

	return nil
}
