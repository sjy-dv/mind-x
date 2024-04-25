import "./App.css";
import React, { useState, useEffect, useRef } from "react";

function App() {
  const [chatMessage, setChatMessage] = useState("");
  const [learningInput, setLearningInput] = useState("");
  const [chat, setChat] = useState([]);
  const ws = useRef(null);

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:3001/chat");
    ws.current.onmessage = (e) => {
      const response = JSON.parse(e.data);
      console.log("bot reply", response);
      if (response.status && response.status.is_error) {
        setChat((prevChat) => [
          ...prevChat,
          "Error: " + response.status.error_msg,
        ]);
      } else {
        setChat((prevChat) => [...prevChat, response.message]);
      }
    };
    return () => {
      ws.current.close();
    };
  }, []);

  const sendChatMessage = () => {
    if (ws.current && chatMessage !== "") {
      ws.current.send(
        JSON.stringify({ message: chatMessage, message_type: "chat" })
      );
      setChat((prevChat) => [...prevChat, chatMessage]);
      setChatMessage("");
    }
  };

  const handleLearningInput = () => {
    if (ws.current && learningInput !== "") {
      ws.current.send(
        JSON.stringify({ message: learningInput, message_type: "Training" })
      );
      setLearningInput("");
    }
  };

  return (
    <div className="App">
      <ul className="chat-list">
        {chat.map((msg, index) => (
          <li key={index} className="chat-message">
            {msg}
          </li>
        ))}
      </ul>
      <div className="input-container">
        <input
          type="text"
          className="input-field"
          placeholder="Chat message..."
          value={chatMessage}
          onChange={(e) => setChatMessage(e.target.value)}
          onKeyDown={(event) => {
            if (event.key === "Enter") {
              sendChatMessage();
            }
          }}
        />
        <button onClick={sendChatMessage} className="send-button">
          Send
        </button>
      </div>
      <div className="input-container">
        <input
          type="text"
          className="input-field"
          placeholder="Learning input ..."
          value={learningInput}
          onChange={(e) => setLearningInput(e.target.value)}
          onKeyDown={(event) => {
            if (event.key === "Enter") {
              handleLearningInput();
            }
          }}
        />
        <button onClick={handleLearningInput} className="send-button">
          Training
        </button>
      </div>
    </div>
  );
}

export default App;
