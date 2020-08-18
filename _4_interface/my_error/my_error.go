package my_error

import (
	"encoding/json"
)

// 实现error接口
type err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *err) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func New(code int, msg string) *err {
	return &err{
		code,
		msg,
	}
}
