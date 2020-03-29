package demo

import (
	"net/http"
	"strconv"
	"log"
	"strings"
)

type Handler func(c *Context)

type handlerMap map[string]Handler
//app
type App struct {
	router map[string]handlerMap
}

//创建实例
func Default()*App{
	return &App{
		router: make(map[string]handlerMap),
	}
}

//方法

//GET
func(a *App)GET(uri string,handler Handler){
	a.handle("GET",uri,handler)
}

//POST
func(a *App)POST(uri string,handler Handler){
	a.handle("POST",uri,handler)
}

//统一处理
func(a *App)handle(method string,uri string,handler Handler ){
	handlers,ok :=a.router[method]
	if !ok {
		m := make(handlerMap)
		a.router[method] = m
		handlers = m
	}
	_, ok2:= handlers[uri]
	if ok2 {
		panic("same route")
	}

	handlers[uri] = handler
}

//请求端口
func(a *App)Run(port int){
	portS := strconv.FormatInt(int64(port), 10)
	http.Handle("/", a)
	if err := http.ListenAndServe(":"+portS, nil); err != nil {
		log.Fatal(err.Error())
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	httpMethod := req.Method
	uri := req.RequestURI
	uris := strings.Split(uri, "?")
	if len(uris) < 1 {
		return
	}

	handlers, ok := a.router[httpMethod]
	if !ok {
		log.Println("may by a hacker:", req.RemoteAddr)
		return
	}
	h, ok := handlers[uris[0]]
	if !ok {
		Handler404(w, req)
		return
	}

	c := NewContext(w, req)

	h(&c)

}

func Handler404(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("404 not foundddddddd"))
}
