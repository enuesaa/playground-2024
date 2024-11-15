import os
from langchain_huggingface import HuggingFaceEmbeddings
from langchain_redis import RedisConfig, RedisVectorStore

embeddings = HuggingFaceEmbeddings(model_name='sentence-transformers/all-mpnet-base-v2')
config = RedisConfig(
    index_name="files",
    redis_url='redis://localhost:6379',
    metadata_schema=[
        {"name": "filename", "type": "tag"},
        {"name": "filepath", "type": "tag"},
    ],
)
vector_store = RedisVectorStore(embeddings, config=config)
basepath = "../cpbuf/pkg"

def create():
    print('a')
    for root, _, files in os.walk(basepath):
        print('b')
        for file in files:
            file_path = os.path.join(root, file)
            print(file_path)
            
            with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
                text = f.read()

            metadata = {
                "filename": file,
                "filepath": file_path
            }
            ids = vector_store.add_texts([text], [metadata])
            print(ids)


def search():
    results = vector_store.similarity_search("WorkDir", k=7)

    for doc in results:
        print('---------------')
        print(f"Content: {doc.page_content[:20]}...")
        print(f"Metadata: {doc.metadata}")
