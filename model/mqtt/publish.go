/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/8/27 下午 04:14
 */

package mqtt

import (
	"hass/model"
)

type ServiceData struct {
	Topic           string `json:"topic"`
	PayloadTemplate string `json:"payload_template"`
}

func (ServiceData) Domain() string {
	return "mqtt"
}
func NewPublishService() *model.Service[ServiceData, any] {
	service := new(model.Service[ServiceData, any])
	service.ServiceData = new(ServiceData)
	service.Domain = ServiceData{}.Domain()
	service.Service = "publish"
	return service
}
