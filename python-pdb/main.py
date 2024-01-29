import json
from lib import apply

apply(json)


ret = json.dumps(['foo', {'bar': ('baz', None, 1.0, 2)}])
print(ret)
