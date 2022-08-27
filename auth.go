package hass

import (
	"bytes"
	"encoding/json"
	"log"
)

type authRequest struct {
	Type      string `json:"type"`
	HaVersion string `json:"ha_version"`
	Message   string `json:"message"`
}
type authResponse struct {
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
}

func (w *WS) login(value *bytes.Buffer) {
	var req authRequest
	json.Unmarshal(value.Bytes(), &req)
	switch req.Type {
	case "auth_required":
		json.NewEncoder(w.conn).Encode(authResponse{
			Type:        "auth",
			AccessToken: w.token,
		})
	case "auth_ok":
		log.Println("login success")
		w.isLogin = true
		// 登录成功后
		for _, plugin := range w.Plugins {
			go plugin.Init(&Func{w})
		}
	case "auth_invalid":
		log.Println(req.Message)
		w.cancel()
	}
}
