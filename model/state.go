/**
 * @Author: dsreshiram@gmail.com
 * @Date: 2022/8/27 上午 11:14
 */

package model

import "time"

type State struct {
	EntityId    string    `json:"entity_id"`
	State       string    `json:"state"`
	Attributes  any       `json:"attributes"`
	LastChanged time.Time `json:"last_changed"`
	LastUpdated time.Time `json:"last_updated"`
	Context     `json:"context"`
}

type Context struct {
	ID       string  `json:"id"`
	ParentID *string `json:"parent_id"`
	UserID   *string `json:"user_id"`
}
