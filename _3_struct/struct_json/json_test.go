package struct_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// struct转json, 需要字段名首字母大写, 用于外部包来访问  结构体名则不需要
// tag 需要加""
type Company struct {
	Name      string    `json:"name" ini:"name"`
	Address   address   `json:"address"`
	WorkPlace workPlace `json:"workPlace"`
}
type address struct {
	Province string `json:"province"`
	City     string `json:"city"`
}
type workPlace struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

func TestJson(t *testing.T) {
	c := Company{
		"domiii",
		address{"陕西", "西安"},
		workPlace{"广州", "珠海"},
	}

	// 序列化
	jsonStr, err := json.Marshal(c)
	if err == nil {
		fmt.Println(string(jsonStr))
	}

	// 反序列化 unmarshal 只接受指针
	str := `{"name":"domiii","address":{"province":"陕西","city":"西安"},"workPlace":{"province":"广州","city":"珠海"}}`
	var d Company
	json.Unmarshal([]byte(str), &d)
	fmt.Printf("%v \n", d)

}
