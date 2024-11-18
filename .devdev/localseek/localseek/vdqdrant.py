from langchain_huggingface import HuggingFaceEmbeddings
from langchain_qdrant import QdrantVectorStore
from qdrant_client import QdrantClient
from qdrant_client.http.models import Distance, VectorParams
import pathlib
 
def connect() -> QdrantVectorStore:
    embeddings = HuggingFaceEmbeddings(model_name='sentence-transformers/all-mpnet-base-v2')
    client = QdrantClient('http://localhost:6333')

    if not client.collection_exists('rag'):
        client.create_collection(
            collection_name='rag',
            vectors_config=VectorParams(size=768, distance=Distance.COSINE),
        )


    return QdrantVectorStore(client, 'rag', embedding=embeddings)

def save(basepath: str):
    store = connect()

    p = pathlib.Path(basepath)
    for path in p.glob('**/*'):
        if '.git' in path.parts:
            continue
        if path.is_dir():
            continue

        try:
            with open(path, errors='ignore') as f:
                text = f.read()
        except:
            continue

        metadata = {'path': path.as_posix()}
        ids = store.add_texts([text], [metadata])
        print(f'{path}: {ids[0]}')

def search(keyword: str):
    store = connect()
    list = store.similarity_search_with_score(keyword, k=7)

    for item in list:
        doc = item[0]
        path = doc.metadata['path']
        score = item[1] # 小さいほど近い
        print(f'path: {path}')
        print(f'score: {score}')
        print('---------------')
