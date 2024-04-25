package transport

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sjy-dv/mind-x/processor/mxvd"
	"github.com/tmc/langchaingo/llms/ollama"
)

const (
	ErrRpcConn              string = "gRpcConnectionLost:"
	ErrWebsocketUpgrade     string = "WebSocketUpgradeError:"
	ErrUnexpectedClose      string = "UnexpectedCloseError:"
	ErrUndefinedReadMessage string = "UndefinedReadMessageError:"
	ErrWriteMessage         string = "Write Message Error:"
	ErrUnexpectedType       string = "UnexpectedTypeError:"
	ErrJsonUnmarshal        string = "JsonUnmarshalError"
	ErrCreateRpcStream      string = "Failed to Create Rpc Stream"
	STATUS_ERROR            string = "status_error"
	STATUS_OK               string = "status_ok"
	InvalidVersion          string = "invalid version, closing websocket"
)

type WebSocketServer struct {
	Upgrader websocket.Upgrader
	llama3   *ollama.LLM
	mxvd     *mxvd.MXVDProxy
}

func NewSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  11 << 20,
			WriteBufferSize: 11 << 20,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

type userText struct {
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

type botText struct {
	Message string   `json:"message"`
	Status  wsStatus `json:"status"`
}

type wsStatus struct {
	ErrorMsg    string `json:"error_msg"`
	OriginError error  `json:"origin_error"`
	IsError     bool   `json:"is_error"`
}

type DialQueue struct {
	bk bool
}

func (wss *WebSocketServer) Run(llm *ollama.LLM, vdb *mxvd.MXVDProxy) {
	wss.llama3 = llm
	wss.mxvd = vdb
	http.HandleFunc("/chat", wss.personalpipe)
	server := &http.Server{
		Addr:         ":3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("MIND-X Processor Starting...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (wss *WebSocketServer) personalpipe(w http.ResponseWriter, r *http.Request) {
	ws, err := wss.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	go wss.mindXpipe(ws)
}

func (wss *WebSocketServer) sendErrorMessage(ws *websocket.Conn, errMsg string, originalErr error) error {
	err := ws.WriteJSON(botText{
		Status: wsStatus{
			IsError:     true,
			ErrorMsg:    errMsg,
			OriginError: originalErr,
		},
	})
	if err != nil {
		log.Println(ErrWriteMessage, err)
	}
	return err
}
