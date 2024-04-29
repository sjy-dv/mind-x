import protobuf
import grpc
import uuid
import time
import numpy as np
import protobuf.dataset_pb2
import protobuf.dataset_pb2_grpc
import protobuf.search_pb2
import protobuf.search_pb2_grpc

import torch
import torchvision.models as models
import torchvision.transforms as transforms
from PIL import Image


model = models.resnet18(pretrained=True)
model.eval()

preprocess = transforms.Compose([
    transforms.Resize(256),
    transforms.CenterCrop(224),
    transforms.Grayscale(num_output_channels=3),
    transforms.ToTensor(),
    transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
])

def get_image_embedding(image_path):
    image = Image.open(image_path)
    image = preprocess(image).unsqueeze(0)  
    with torch.no_grad():
        embedding = model(image)
    return embedding.cpu().numpy().flatten().astype(np.float32)



def main():
    server_addr = 'localhost:50003'
    
    with grpc.insecure_channel(server_addr) as channel:
        data_manager = protobuf.dataset_pb2_grpc.DataManagerStub(channel)
        dataset_manager = protobuf.dataset_pb2_grpc.DatasetManagerStub(channel)
        search = protobuf.search_pb2_grpc.SearchStub(channel)
        
        # create dataset
        dataset = dataset_manager.Create(protobuf.dataset_pb2.Dataset(
            dimension=1000,
            partition_count=1,
            replication_factor=3
        ))
        dataset_id = dataset.id
        print("Dataset created with ID:", dataset_id)
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=get_image_embedding("./dataset/good.png"),
            metadata={"status": "good"}
        ))
        data_manager.Insert(protobuf.dataset_pb2.InsertRequest(
            dataset_id=dataset_id,
            id= uuid.uuid4().bytes,
            value=get_image_embedding("./dataset/bad.png"),
            metadata={"status": "bad"}
        ))
       
        query = protobuf.search_pb2.SearchRequest(
            dataset_id=dataset_id,
            query=get_image_embedding("./test_data/bad.png"),
            k=1
        )
        results = search.Search(query)
        try:
            for result in results:
                print("Search result:", result.id, result.metadata, result.score)
        except grpc.RpcError as e:
            print("Search failed:", e)
        query = protobuf.search_pb2.SearchRequest(
            dataset_id=dataset_id,
            query=get_image_embedding("./test_data/good.png"),
            k=1
        )
        results = search.Search(query)
        try:
            for result in results:
                print("Search result:", result.id, result.metadata, result.score)
        except grpc.RpcError as e:
            print("Search failed:", e)
if __name__ == '__main__':
    main()