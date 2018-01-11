package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	"net/url"
	"time"
	"longsocket"
)
func myServerHandle(ws *websocket.Conn){
	req := ws.Request()
	fmt.Println(req)
	u,err :=url.Parse(req.Header.Get("Origin"))
	fmt.Println("the request is",u)
	if err !=nil{
		fmt.Println("there is an err in server")
		ws.Close()
		return
	}
	user :=u.Query().Get("user")
	password := u.Query().Get("password")
	fmt.Println(user,password)
	myScoket := longsocket.NewConn("","","",true,256*1024)
	myScoket.SetSocket(ws)
	defer  myScoket.Close()
	go myScoket.WriteLoop()
	go myScoket.ReadLoop()
	myScoket.Read(dealMsg)
}
func dealMsg(msg []byte , l *longsocket.Longsocket) error{
	fmt.Println(string(msg),"dealMsg")
	return nil
}

func main() {
	http.Handle("/server",websocket.Handler(myServerHandle))
	srv := &http.Server{
		Addr: ":12345",
		Handler:nil,
		ReadTimeout: time.Duration(5) *time.Minute,
		WriteTimeout: time.Duration(5) *time.Minute,
		MaxHeaderBytes: 1 <<30,
	}
	err :=srv.ListenAndServe()
	if err != nil{
		fmt.Println(err)
	}
}
