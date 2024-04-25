package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime/debug"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"github.com/sjy-dv/mind-x/processor/embeddings"
	"github.com/sjy-dv/mind-x/processor/entity"
	"github.com/sjy-dv/mind-x/processor/protobuf/protocol/v0"
	"github.com/sjy-dv/mind-x/processor/transformers"
	"github.com/tmc/langchaingo/schema"
)

/*
	1. Two Request Lines
	2. One Line -> Training
	3. Two Line -> Based On Training Data. Using For Rags
*/

func (wss *WebSocketServer) mindXpipe(ws *websocket.Conn) {

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
				vector, err := embeddings.TextEmbedding(msg.Message)
				if err != nil {
					if err := wss.sendErrorMessage(ws, ErrWriteMessage, err); err != nil {
						dialogQueue <- DialQueue{
							bk: true,
						}
						return
					}
				}
				if msg.MessageType == "Training" {
					err = wss.mxvd.Insert(context.Background(), &entity.InsertObject{
						ID:     uuid.NewV4().Bytes(),
						Vector: vector,
						Metadata: map[string]string{
							"sentence": msg.Message,
						},
					})
					if err != nil {
						if err := wss.sendErrorMessage(ws, ErrWriteMessage, err); err != nil {
							dialogQueue <- DialQueue{
								bk: true,
							}
							return
						}
					}
					if err := ws.WriteJSON(botText{
						Message: "training is successful. I'm ready for you.",
					}); err != nil {
						if err := wss.sendErrorMessage(ws, ErrWriteMessage, err); err != nil {
							dialogQueue <- DialQueue{
								bk: true,
							}
							return
						}
					}
				} else {
					reply, err := wss.mxvd.SearchManager().Search(context.Background(),
						&protocol.SearchRequest{
							DatasetId: wss.mxvd.PersonalID,
							Query:     vector,
							K:         10,
						})

					if err != nil {
						if err := wss.sendErrorMessage(ws, ErrWriteMessage, err); err != nil {
							dialogQueue <- DialQueue{
								bk: true,
							}
							return
						}
					}
					referDocs := make([]schema.Document, 0, 10)
					for {
						historyMention, err := reply.Recv()
						if err == io.EOF {
							break
						}
						if err != nil {
							break
						}
						fmt.Println(historyMention.GetMetadata()["sentence"])
						referDocs = append(referDocs, schema.Document{
							PageContent: historyMention.GetMetadata()["sentence"],
						})
					}
					replyMsg := transformers.RagResponse(msg.Message, referDocs, wss.llama3)
					if err := ws.WriteJSON(botText{
						Message: replyMsg,
					}); err != nil {
						if err := wss.sendErrorMessage(ws, ErrWriteMessage, err); err != nil {
							dialogQueue <- DialQueue{
								bk: true,
							}
							return
						}
					}
				}
			}()

			next := <-dialogQueue
			if next.bk {
				break
			}
			continue
		} else {
			break
		}
	}
}
