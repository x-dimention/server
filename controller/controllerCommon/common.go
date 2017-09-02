package controllerCommon

import (
	"encoding/json"
)

func MakeApiJson(_data interface{}, _err error) []byte {
	type returnJson struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	var code int = 200
	var msg string = ""
	if _err != nil {
		code = 503
	}
	postData := &returnJson{
		Code: code,
		Msg:  msg,
		Data: _data,
	}
	b, err := json.Marshal(postData)
	if err != nil {
		var returnfalse = &returnJson{
			Code: 503,
			Msg:  "",
			Data: "",
		}
		b, _ = json.Marshal(returnfalse)
		return b
	}
	return b
}
