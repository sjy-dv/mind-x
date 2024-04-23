import protobuf
import grpc
import uuid
import time
import numpy as np
import protobuf.dataset_pb2
import protobuf.dataset_pb2_grpc
import protobuf.search_pb2
import protobuf.search_pb2_grpc

from sentence_transformers import SentenceTransformer

model = SentenceTransformer('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')
# embeddings = model.encode(sentences)
# print(embeddings[0])

def sentence_embd(s):
    embedding = model.encode([s])[0].astype(np.float32)
    print("Embedding:", embedding)  
    return embedding.tolist()

def main():
    server_addr = 'localhost:50003'
    
    with grpc.insecure_channel(server_addr) as channel:
        data_manager = protobuf.dataset_pb2_grpc.DataManagerStub(channel)
        dataset_manager = protobuf.dataset_pb2_grpc.DatasetManagerStub(channel)
        search = protobuf.search_pb2_grpc.SearchStub(channel)
        
        # create dataset
        dataset = dataset_manager.Create(protobuf.dataset_pb2.Dataset(
            dimension=384,
            partition_count=1,
            replication_factor=3
        ))
        dataset_id = dataset.id
        print("Dataset created with ID:", dataset_id)
        sentence="I eat chicken when I'm hungry."        
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=sentence_embd(sentence),
            metadata={"sentence": sentence}
        ))
        sentence="I eat hamburger when I'm hungry."        
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=sentence_embd(sentence),
            metadata={"sentence": sentence}
        ))
        sentence="I usually also enjoy eating toast."        
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=sentence_embd(sentence),
            metadata={"sentence": sentence}
        ))
        sentence="I like to tidy up my house when I'm bored."        
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=sentence_embd(sentence),
            metadata={"sentence": sentence}
        ))
        sentence="When I usually eat, I have to drink carbonated water with it for better digestion."        
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=sentence_embd(sentence),
            metadata={"sentence": sentence}
        ))
        sentence = "I'm hungry. What should I eat?"
        query = protobuf.search_pb2.SearchRequest(
            dataset_id=dataset_id,
            query=sentence_embd(sentence),
            k=10
        )
        results = search.Search(query)
        try:
            for result in results:
                print("Search result:", result.id, result.metadata, result.score)
        except grpc.RpcError as e:
            print("Search failed:", e)
        sentence="I'm bored. What should I do?"
        query = protobuf.search_pb2.SearchRequest(
            dataset_id=dataset_id,
            query=sentence_embd(sentence),
            k=10
        )
        results = search.Search(query)
        try:
            for result in results:
                print("Search result:", result.id, result.metadata, result.score)
        except grpc.RpcError as e:
            print("Search failed:", e)
if __name__ == '__main__':
    main()