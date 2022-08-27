package light

import (
	"hass/model"
)

type ServiceData struct {
	ColorName  string `json:"color_name,omitempty"`
	Brightness string `json:"brightness,omitempty"`
}

func (ServiceData) Domain() string {
	return "light"
}
func NewSevice() *model.Service[ServiceData, model.Target] {
	out := new(model.Service[ServiceData, model.Target])
	out.ServiceData = new(ServiceData)
	out.Target = new(model.Target)
	out.Domain = ServiceData{}.Domain()
	return out
}
