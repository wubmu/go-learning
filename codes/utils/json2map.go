package main

import (
	"encoding/json"
	"fmt"
)

type RequestBody struct {
	req string
}

//以指针的方式传入，但在使用时却可以不用关心
// result 是函数内的临时变量，作为返回值可以直接返回调用层
func (req *RequestBody) json2map() (s map[string]interface{}, err error) {
	var res map[string]interface{}
	if err := json.Unmarshal([]byte(req.req), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	var r RequestBody
	r.req = `{"name": "xym","sex": "male"}`
	if req2map, err := r.json2map(); err == nil {
		fmt.Println(req2map["name"])
		fmt.Println(req2map)
	} else {
		fmt.Println(err)
	}
}
