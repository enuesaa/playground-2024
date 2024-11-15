# see https://redis.io/docs/latest/develop/get-started/vector-database/

import numpy as np
import requests
import redis
from redis.commands.search.field import (
    NumericField,
    TagField,
    TextField,
    VectorField,
)
from redis.commands.search.indexDefinition import IndexDefinition, IndexType
from redis.commands.search.query import Query
from sentence_transformers import SentenceTransformer

def main():
    client = redis.Redis(host='127.0.0.1', port=6379, decode_responses=True)

    url = "https://raw.githubusercontent.com/bsbodden/redis_vss_getting_started/main/data/bikes.json"
    response = requests.get(url, timeout=10)
    bikes = response.json()

    pipeline = client.pipeline()
    for i, bike in enumerate(bikes, start=1):
        redis_key = f"bikes:{i:03}"
        pipeline.json().set(redis_key, "$", bike)
    res = pipeline.execute()

    embedder = SentenceTransformer('msmarco-distilbert-base-v4')
    keys = client.keys("bikes:*")

    descriptions = client.json().mget(keys, "$.description")
    descriptions = [item for sublist in descriptions for item in sublist]
    embedder = SentenceTransformer("msmarco-distilbert-base-v4")
    embeddings = embedder.encode(descriptions).astype(np.float32).tolist()
    VECTOR_DIMENSION = len(embeddings[0])

    pipeline = client.pipeline()
    for key, embedding in zip(keys, embeddings):
        pipeline.json().set(key, "$.description_embeddings", embedding)
    pipeline.execute()

    schema = (
        TextField("$.model", no_stem=True, as_name="model"),
        TextField("$.brand", no_stem=True, as_name="brand"),
        NumericField("$.price", as_name="price"),
        TagField("$.type", as_name="type"),
        TextField("$.description", as_name="description"),
        VectorField(
            "$.description_embeddings",
            "FLAT",
            {
                "TYPE": "FLOAT32",
                "DIM": VECTOR_DIMENSION,
                "DISTANCE_METRIC": "COSINE",
            },
            as_name="vector",
        ),
    )
    definition = IndexDefinition(prefix=["bikes:"], index_type=IndexType.JSON)
    client.ft("idx:bikes_vss").create_index(fields=schema, definition=definition)

    query = "Bike for small kids"
    encoded_query = embedder.encode(query)

    query = Query('(*)=>[KNN 3 @vector $query_vector AS vector_score]') \
        .sort_by('vector_score') \
        .return_fields('vector_score', 'id', 'brand', 'model', 'description') \
        .dialect(2)
    
    results = client.ft("idx:bikes_vss") \
        .search(
            query,
            {"query_vector": np.array(encoded_query, dtype=np.float32).tobytes()}
        ) \
        .docs

    for doc in results:
        vector_score = round(1 - float(doc.vector_score), 2)
        print({
            "query": query,
            "score": vector_score,
            "id": doc.id,
            "brand": doc.brand,
            "model": doc.model,
            "description": doc.description,
        })
