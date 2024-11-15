import redis
import random
import json
import time
import numpy as np
from redis.commands.search.query import Query

def main():
    client = redis.Redis(host='127.0.0.1', port=6379, decode_responses=True)
 