# see https://redis.io/docs/latest/develop/get-started/vector-database/

import numpy as np
import requests
import redis
from redis.commands.search.field import NumericField, TagField, TextField, VectorField
from redis.commands.search.indexDefinition import IndexDefinition, IndexType
from redis.commands.search.query import Query
from sentence_transformers import SentenceTransformer

client = redis.Redis(host='127.0.0.1', decode_responses=True)
embedder = SentenceTransformer('msmarco-distilbert-base-v4')

def fetch_bikes():
    url = 'https://raw.githubusercontent.com/bsbodden/redis_vss_getting_started/main/data/bikes.json'
    response = requests.get(url)
    return response.json()

def create():
    bikes = fetch_bikes()

    with client.pipeline() as pipeline:
        for i, bike in enumerate(bikes, start=1):
            pipeline.json().set(f'bikes:{i:03}', '$', bike)
        pipeline.execute()

    keys = client.keys('bikes:*')
    descriptions = client.json().mget(keys, '$.description')
    descriptions = [item for sublist in descriptions for item in sublist]
    embeddings = embedder.encode(descriptions).astype(np.float32).tolist()
    dimension = len(embeddings[0])

    with client.pipeline() as pipeline:
        for key, embedding in zip(keys, embeddings):
            pipeline.json().set(key, '$.description_embeddings', embedding)
        pipeline.execute()

    schema = (
        TextField('$.model', no_stem=True, as_name='model'),
        TextField('$.brand', no_stem=True, as_name='brand'),
        NumericField('$.price', as_name='price'),
        TagField('$.type', as_name='type'),
        TextField('$.description', as_name='description'),
        VectorField('$.description_embeddings',
            algorithm='FLAT',
            attributes={
                'TYPE': 'FLOAT32',
                'DIM': dimension,
                'DISTANCE_METRIC': 'COSINE',
            },
            as_name='vector',
        ),
    )
    definition = IndexDefinition(prefix=['bikes:'], index_type=IndexType.JSON)
    client.ft('idx:bikes_vss').create_index(fields=schema, definition=definition)

def search():
    query = 'Bike for small kids'
    encoded_query = embedder.encode(query)

    query = Query('(*)=>[KNN 3 @vector $query_vector AS vector_score]') \
        .sort_by('vector_score') \
        .return_fields('vector_score', 'id', 'brand', 'model', 'description') \
        .dialect(2)
    
    results = client.ft('idx:bikes_vss') \
        .search(query, {
            'query_vector': np.array(encoded_query, dtype=np.float32).tobytes(),
        }) \
        .docs

    for doc in results:
        print({
            'score': round(1 - float(doc.vector_score), 2),
            'id': doc.id,
            'brand': doc.brand,
            'description': doc.description,
        })
