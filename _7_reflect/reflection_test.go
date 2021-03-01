package _7_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 反射原理
// http://wen.topgoer.com/docs/gozhuanjia/chapter066.1-reflect

// typeOf 和 valueOf
// value对象能获取到type类型, 而type类型不能取到值

// type name 和 type kind
// 反射中类型分为两种: 类型(type) 种类(kind)
// type: 多为自定义类型, 例如下面打印的type.name=order
// kind: 底层类型, kind打印为struct
func TestReflect(t *testing.T) {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	//o:=10
	fmt.Println("--------------ValueOf-----------------------")
	v := reflect.ValueOf(o)
	fmt.Println("获取自定义类型----- ", v.Type().Name())
	fmt.Println("获取原始类型------- ", v.Kind())
	fmt.Println("类型判断----------  reflect.Struct == v.Kind():", reflect.Struct == v.Kind())
	fmt.Println("获取字段值---------", v.Field(0).Int())
	fmt.Println("value获取tag------", v.Type().Field(0).Tag.Get("db"))
	fmt.Println("--------------TypeOf------------------------")
	ty := reflect.TypeOf(o)
	fmt.Println("获取自定义类型----- ", ty.Name())
	fmt.Println("获取原始类型------- ", ty.Kind())
	fmt.Println("获取field Tag----- ", ty.Field(0).Tag.Get("db"))

}

func TestCreateQuery(t *testing.T) {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	CreateQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	CreateQuery(e)
	i := 90
	CreateQuery(i)
}
