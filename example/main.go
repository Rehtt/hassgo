/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/8/27 下午 05:26
 */

package main

import (
	"github.com/Rehtt/hass"
	"log"
)

var (
	origin = ""
	url    = ""
	token  = ""
)

func main() {
	ws, err := hass.Connect(url, origin, token)
	if err != nil {
		log.Panicln(err)
	}
	ws.Plugins.Register(&Door{})
	ws.Run()
}
