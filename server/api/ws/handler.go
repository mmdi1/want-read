package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"want-read/core/message"

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
			log.Println("read ws unmarshal err:", err, msg)
		}
		log.Println("===========>", msg)
	}
}
func (conn *connection) SendMsg(data any) {
	log.Println("send ws:", data)
	bt, err := json.Marshal(data)
	if err != nil {
		message.TipErr(err.Error())
		return
	}
	err = conn.Conn.WriteMessage(1, bt)
	if err != nil {
		message.TipErr(err.Error())
		return
	}
}
func WsHandler(c *gin.Context) {
	log.Println("init ws")
	ws, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ws err:", err)
		return
	}
	Conn = newConnection(ws)
	Conn.start()
}
