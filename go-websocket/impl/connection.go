package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type wsConnection struct {
	wsSocket  *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	mutex    sync.Mutex
	isClosed bool
}

func InitConnection(wsConn *websocket.Conn) (conn *wsConnection, err error) {
	conn = &wsConnection{
		wsSocket:  wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// 启动读协程
	go conn.readLoop()

	// 启动写协程
	go conn.writeLoop()

	return
}

// ReadMessage API
func (conn *wsConnection) ReadMessage() (data []byte, err error) {
	// 用户调用的时候也可能阻塞，没有数据
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan: // 连接被关闭了
		err = errors.New("connection is closed")
	}

	return
}

func (conn *wsConnection) WriteMessage(data []byte) (err error) {
	conn.outChan <- data
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *wsConnection) Close() {
	// 线程安全，可重入的Close
	err := conn.wsSocket.Close()
	if err != nil {
		log.Println(err)
	}

	// 保证这行代码只执行一次，
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// 内部实现
// 不停的从长连接上读数据
func (conn *wsConnection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		// _消息类型
		if _, data, err = conn.wsSocket.ReadMessage(); err != nil {
			goto ERR
		}
		// 队列满了会阻塞， 等待inChan有空闲的位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			// 当 close chan 被关闭的时候
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *wsConnection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {

		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsSocket.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}

	}

ERR:
	conn.Close()
}
