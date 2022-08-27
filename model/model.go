/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/8/27 上午 09:00
 */

package model

type Head struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Success bool   `json:"success,omitempty"`
	Result  any    `json:"result,omitempty"`
}
