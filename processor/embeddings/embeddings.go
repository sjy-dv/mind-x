package embeddings

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

const defaultURL = "ws://localhost:8765"

var ec *websocket.Conn

func InitEmbeddings() error {
	c, _, err := websocket.DefaultDialer.Dial(defaultURL, nil)
	if err != nil {
		return err
	}
	ec = c
	return nil
}

func TextEmbedding(text string) ([]float32, error) {
	msgBytes, err := json.Marshal(map[string]string{"sentence": text})
	if err != nil {
		return nil, err
	}
	err = ec.WriteMessage(websocket.TextMessage, msgBytes)
	if err != nil {
		return nil, err
	}

	_, msg, err := ec.ReadMessage()
	if err != nil {
		return nil, err
	}
	var embeddings []float32
	err = json.Unmarshal(msg, &embeddings)
	if err != nil {
		return nil, err
	}
	return embeddings, nil
}
