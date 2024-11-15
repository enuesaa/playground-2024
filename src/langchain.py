from langchain_huggingface import HuggingFaceEmbeddings
from langchain_redis import RedisConfig, RedisVectorStore
from sklearn.datasets import fetch_20newsgroups

embeddings = HuggingFaceEmbeddings(model_name='sentence-transformers/all-mpnet-base-v2')
config = RedisConfig(
    index_name="newsgroups",
    redis_url='redis://localhost:6379',
    metadata_schema=[
        {"name": "category", "type": "tag"},
    ],
)
vector_store = RedisVectorStore(embeddings, config=config)


def create():
    newsgroups = fetch_20newsgroups(
        subset="train",
        categories=["alt.atheism", "sci.space"],
        shuffle=True,
        random_state=42,
    )
    texts = newsgroups.data[:250]
    metadata = [
        {"category": newsgroups.target_names[target]} for target in newsgroups.target[:250]
    ]
    print(metadata)
    ids = vector_store.add_texts(texts, metadata)
    print(ids[0:10])

def search():
    results = vector_store.similarity_search("What planet in the solar system has the largest number of moons?", k=2)

    for doc in results:
        print(f"Content: {doc.page_content[:100]}...")
        print(f"Metadata: {doc.metadata}")
