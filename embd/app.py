import asyncio
import websockets
import json
from sentence_transformers import SentenceTransformer
model = SentenceTransformer('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')

async def handle_sentence(websocket, path):
    async for message in websocket:
        print(f"Received: {message}")
        data = json.loads(message)['sentence']
        embedding = model.encode(data)
        await websocket.send(embedding[0])

async def main():
    async with websockets.serve(handle_sentence, "localhost", 8765):
        print("Server started")
        await asyncio.Future() 

if __name__ == "__main__":
    asyncio.run(main())
