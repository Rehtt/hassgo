/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/8/27 上午 11:34
 */

package model

const (
	TurnON  = ServiceS("turn_on")
	TurnOFF = ServiceS("turn_off")
)

type ServiceS string

type Service[data, target any] struct {
	Service     ServiceS `json:"service"`
	Domain      string   `json:"domain"`
	ServiceData *data    `json:"service_data,omitempty"`
	Target      *target  `json:"target,omitempty"`
}
type Target struct {
	EntityId string `json:"entity_id,omitempty"`
}
