from sentence_transformers import SentenceTransformer
sentences = ["This is an example sentence"]

model = SentenceTransformer('sentence-transformers/paraphrase-multilingual-MiniLM-L12-v2')
embeddings = model.encode(sentences)
print(embeddings[0])

# 384 vectors, so Cool