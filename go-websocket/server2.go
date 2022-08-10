package main

//
//import (
//	"fmt"
//	"github.com/gorilla/websocket"
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
//func (wsConn *wsConnection) procLoop() {
//	// 启动一个gouroutine发送心跳
//	go func() {
//		for {
//			time.Sleep(2 * time.Second)
//			if err := wsConn.wsWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
//				fmt.Println("heartbeat fail")
//				wsConn.wsClose()
//				break
//			}
//		}
//	}()
//
//	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
//	for {
//		msg, err := wsConn.wsRead()
//		if err != nil {
//			fmt.Println("read fail")
//			break
//		}
//		fmt.Println(string(msg.data))
//		err = wsConn.wsWrite(msg.messageType, msg.data)
//		if err != nil {
//			fmt.Println("write fail")
//			break
//		}
//	}
//}
//
//func wsHandler(w http.ResponseWriter, r *http.Request) {
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
