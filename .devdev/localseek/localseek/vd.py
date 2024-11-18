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
    list = store.similarity_search_with_score(keyword, k=7)

    for item in list:
        doc = item[0]
        path = doc.metadata['path']
        score = item[1] # 小さいほど近い
        print(f'path: {path}')
        print(f'score: {score}')
        print('---------------')
