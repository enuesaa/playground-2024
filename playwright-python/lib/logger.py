import inspect
from types import ModuleType
import functools
import asyncio

def log(func, prefix=''):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print(prefix, func.__name__, args)
        try:
            result = func(*args, **kwargs)
        except Exception as e:
            raise e
        return result
    return wrapper

def debugroll(obj):
    for attr in dir(obj):
        if attr.startswith('__') or attr.startswith('_'):
            continue
        val = getattr(obj, attr)
        if callable(val):
            if asyncio.iscoroutinefunction(val):
                continue
            if isinstance(obj, ModuleType):
                setattr(obj, attr, log(val, prefix=f"instance {obj.__name__}", ))
            else:
                setattr(obj, attr, log(val, prefix=f"module {obj.__class__.__name__}"))
