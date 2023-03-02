package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Conn *connection
)

type connection struct {
	Conn    *websocket.Conn
	outChan chan []byte
}

func newConnection(ws *websocket.Conn) *connection {
	return &connection{
		Conn: ws,
	}
}
func (conn *connection) start() {
	go conn.readLoop()
}
func (conn *connection) readLoop() {
	for {
		_, msgData, err := conn.Conn.ReadMessage()
		if err != nil {
			log.Println("read ws err:", err)
			return
		}
		var msg string
		if err := json.Unmarshal(msgData, &msg); err != nil {
			fmt.Println("read ws unmarshal err:", err, msg)
		}
		fmt.Println("===========>", msg)
	}
}
func (conn *connection) SendMsg(data any) {
	fmt.Println("send ws:", data)
	err := conn.Conn.WriteJSON(data)
	if err != nil {
		log.Println("send ws err:", err, data)
	}
}
func WsHandler(c *gin.Context) {
	fmt.Println("init ws")
	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws err:", err)
		return
	}
	Conn = newConnection(ws)
	Conn.start()
}
