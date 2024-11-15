import os
from langchain_huggingface import HuggingFaceEmbeddings
from langchain_redis import RedisConfig, RedisVectorStore
import pathlib
 
def connect() -> RedisVectorStore:
    embeddings = HuggingFaceEmbeddings(model_name='sentence-transformers/all-mpnet-base-v2')
    config = RedisConfig(
        index_name='files',
        redis_url='redis://localhost:6379',
        metadata_schema=[
            {'name': 'path', 'type': 'tag'},
        ],
    )
    return RedisVectorStore(embeddings, config=config)

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
    res = store.similarity_search(keyword, k=7)

    for doc in res:
        content = doc.page_content[:20]
        path = doc.metadata['path']
        print(f'{content}... in {path}')
        print('---------------')
