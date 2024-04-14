import requests
from debugroll import debugroll

debugroll(requests)

res = requests.get('https://example.com/')
