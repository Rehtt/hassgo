package hassgo

func (w *WS) sendHandle(request func(id uint64) (data []byte), option ...func() (subType uint8, p Plugin)) ([]byte, error) {
	ctx := w.newContext()
	if len(option) > 0 {
		ctx.subType, ctx.plugin = option[0]()
	} else {
		defer ctx.destroyContext()
	}
	// 简化
	if _, err := w.conn.Write(request(ctx.id)); err != nil {
		return nil, err
	}
	out := ctx.ReadAll()
	return out, nil
	//if err := ctx.send(request); err != nil {
	//	return nil, err
	//}
	//b, err := ctx.recv(response)
}

//func (ctx *Context) send(request any) (err error) {
//	var tmpM map[string]any
//	var tmp bytes.Buffer
//
//	switch request := request.(type) {
//	case map[string]any:
//		tmpM = request
//	case string:
//		if err = json.Unmarshal([]byte(request), &tmp); err != nil {
//			return
//		}
//	case []byte:
//		if err = json.Unmarshal(request, &tmp); err != nil {
//			return
//		}
//	default:
//		if err = json.NewEncoder(&tmp).Encode(request); err != nil {
//			return
//		}
//		if err = json.NewDecoder(&tmp).Decode(&tmpM); err != nil {
//			return
//		}
//	}
//	tmpM["id"] = ctx.id
//	err = json.NewEncoder(ctx.ws.conn).Encode(tmpM)
//	return
//}
//func (ctx *Context) recv(v any) (b []byte, err error) {
//	b = ctx.ReadAll()
//	if v != nil {
//		err = json.Unmarshal(b, v)
//	}
//	return
//}
