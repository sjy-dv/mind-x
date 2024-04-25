package transport

import (
	"encoding/json"
	"log"
	"runtime/debug"

	"github.com/gorilla/websocket"
)

/*
	1. Two Request Lines
	2. One Line -> Training
	3. Two Line -> Based On Training Data. Using For Rags
*/

func (wss *WebSocketServer) Pipe(ws *websocket.Conn) {

	defer func(ws *websocket.Conn) {
		ws.Close()
		if r := recover(); r != nil {
			log.Printf("recoverPanic %v:%v", debug.Stack(), r)
		}
	}(ws)

	for {
		mT, buf, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(ErrUnexpectedClose, err)
				break
			}
			log.Println(ErrUndefinedReadMessage, err)
			break
		}

		dialogQueue := make(chan DialQueue, 1)
		if mT == websocket.TextMessage {
			go func() {
				var msg userText
				err := json.Unmarshal(buf, &msg)
				if err != nil {
					if err := wss.sendErrorMessage(ws, ErrJsonUnmarshal, err); err != nil {
						dialogQueue <- DialQueue{
							bk: true,
						}
						return
					}
				}

				if msg.MessageType == "Training" {

				} else {

				}
				dialogQueue <- DialQueue{
					bk: false,
				}
			}()
		} else {
			break
		}
	}
}
