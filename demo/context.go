package demo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//上下文
type Context struct {
	req  *http.Request
	w    http.ResponseWriter
	queryParam map[string]string //url参数
	formParam  map[string]string //表单参数
}

//返回类型

//string
func (c *Context) String(s string) {
	_, _ = c.w.Write([]byte(s))
}
//json
func(c *Context) JSON(i interface{}){
	newi,_ :=json.Marshal(i)
	_, _ = c.w.Write(newi)
}
//封装一个新的上下文
func NewContext(rw http.ResponseWriter, r *http.Request)(ctx Context){
	ctx = Context{
		req:        r,
		w:          rw,
		formParam:  make(map[string]string),
	}
	ctx.queryParam = parseQuery(r.RequestURI)
	return
}


func (c *Context) Query(key string) string {
	v := c.queryParam[key]
	return v
}


//解析uri
func parseQuery(uri string) (res map[string]string) {
	res = make(map[string]string)
	uris := strings.Split(uri, "?")
	if len(uris) == 1 {
		return
	}
	param := uris[len(uris)-1]
	pair := strings.Split(param, "&")
	for _, kv := range pair {
		kvPair := strings.Split(kv, "=")
		if len(kvPair) != 2 {
			fmt.Println(kvPair)
			panic("request error")
		}
		res[kvPair[0]] = kvPair[1]
	}
	return
}