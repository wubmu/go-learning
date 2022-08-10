package main

//
//import (
//	"github.com/gorilla/websocket"
//	"go-websocket/impl"
//	"net/http"
//	"time"
//)
//
//// http升级websocket协议的配置
//var (
//	upgrader = websocket.Upgrader{
//		// 允许跨域
//		CheckOrigin: func(r *http.Request) bool {
//			return true
//		},
//	}
//)
//
//func wsHandler(w http.ResponseWriter, r *http.Request) {
//	var (
//		wsConn *websocket.Conn
//		err    error
//		conn   *impl.wsConnection
//
//		data []byte
//	)
//
//	// Upgrade: websocker
//	//if wsConn, err = upgrader.Upgrade(w, r, nil) {
//
//	for {
//		if msgType, data, err = conn.ReadMessage(); err != nil {
//			goto ERR
//		}
//		if err = conn.WriteMessage(websocket.TextMessage, data)
//	}
//	if conn, err = impl.InitConnection(wsConn); err != nil {
//		goto ERR
//	}
//
//	go func() {
//		var (
//			err error
//		)
//		for {
//			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
//				return
//			}
//			time.Sleep(1 * time.Second)
//		}
//	}()
//
//	for {
//		if data, err = conn.ReadMessage(); err != nil {
//			goto ERR
//		}
//		if err = conn.WriteMessage(data); err != nil {
//			goto ERR
//		}
//	}
//
//ERR:
//	//TODO:关闭连接操作
//	conn.Close()
//}
//
//func main() {
//	http.HandlerFunc("/ws", wsHandler)
//	http.ListenAndServe("0.0.0.0:7777", nil)
//}
