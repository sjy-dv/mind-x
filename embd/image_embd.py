import torch
import torchvision.models as models
import torchvision.transforms as transforms
from PIL import Image
from torch.nn.functional import cosine_similarity

model = models.resnet18(pretrained=True)
model.eval()

preprocess = transforms.Compose([
    transforms.Resize(256),
    transforms.CenterCrop(224),
    transforms.ToTensor(),
    transforms.Normalize(mean=[0.485, 0.456, 0.406], std=[0.229, 0.224, 0.225]),
])

def get_image_embedding(image_path):
    image = Image.open(image_path)
    image = preprocess(image).unsqueeze(0)  
    with torch.no_grad():
        embedding = model(image)
    return embedding

embedding = get_image_embedding('../imgs/logo.png')
print(embedding)


