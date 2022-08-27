package main

import (
	"encoding/json"
	"fmt"
	"github.com/Rehtt/hassgo/model"
	"github.com/Rehtt/hassgo/model/light"
	"github.com/Rehtt/hassgo/model/mqtt"
	"log"
	"time"
)

type Door struct {
	DoorLastOn time.Time
}
type Result struct {
	Result []struct {
		EntityId    string    `json:"entity_id"`
		State       string    `json:"state"`
		LastChanged time.Time `json:"last_changed"`
	} `json:"result"`
}

// 事件回调
func (d *Door) SubscribeEventCallBack(ctx *hass.Func, data []byte) {

}

// 触发器回调
func (d *Door) SubscribeTriggerCallBack(ctx *hass.Func, data []byte) {
	// 3秒种后自动切换关门状态
	go func(ctx *hass.Func) {
		time.Sleep(3 * time.Second)
		service := mqtt.NewPublishService()
		service.ServiceData.Topic = "stat/door/big"
		service.ServiceData.PayloadTemplate = "off"
		ctx.CallService(service)
	}(ctx)

	// 获取灯状态信息
	lightState, lightLastCahnge, err := getStateByID(ctx, "light.yeelight_lamp1_9497_light")
	if err != nil {
		log.Println(err)
		return
	}
	// 获取日出日落
	sunState, _, err := getStateByID(ctx, "sun.sun")
	if err != nil {
		log.Println(err)
		return
	}

	// 当距离上次关灯超过5分钟并且日落的情况下
	if lightState == "off" && time.Now().Sub(lightLastCahnge) > time.Minute*5 && sunState == "below_horizon" {
		lightService := light.NewSevice()
		lightService.Service = model.TurnON
		lightService.Target.EntityId = "light.yeelight_lamp1_9497_light"
		ctx.CallService(lightService)
	}

}

// 初始化
func (d *Door) Init(f *hass.Func) {
	// 注册触发器事件，当switch.door从"off"变化到"on"时触发回调
	f.SubscribeTrigger(map[string]any{
		"platform":  "state",
		"entity_id": "switch.door",
		"from":      "off",
		"to":        "on",
	}, d)

}

func getStateByID(f *hass.Func, id string) (state string, lastChange time.Time, err error) {
	data, err := f.GetStates()
	if err != nil {
		return "", time.Time{}, err
	}
	var r Result
	if err = json.Unmarshal(data, &r); err != nil {
		return "", time.Time{}, err
	}
	for _, v := range r.Result {
		if v.EntityId == id {
			return v.State, v.LastChanged, nil
		}
	}
	return "", time.Time{}, fmt.Errorf("not find :%s", id)
}
