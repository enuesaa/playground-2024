import inspect
from types import ModuleType
import functools

def log(func, prefix='', printer=lambda text: print(text)):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        result = func(*args, **kwargs)
        text = f"{prefix}| {func.__name__}| {args}| {result}"
        printer(text)
        return result
    return wrapper

def debugroll(val, printer=lambda text: print(text)):
    if isinstance(val, ModuleType):
        for attr in dir(val):
            if attr.startswith('__') or attr.startswith('_'):
                continue
            attrVal = getattr(val, attr)
            if callable(attrVal):
                setattr(val, attr, log(attrVal, prefix=f"module {val.__name__}", printer=printer))
    elif inspect.isclass(val):
        if callable(val):
            for attr in dir(val):
                if attr.startswith('__') or attr.startswith('_'):
                    continue
                attrVal = getattr(val, attr)
                if callable(attrVal):
                    setattr(val, attr, log(attrVal, prefix=f"class {val.__class__.__name__}", printer=printer))
    elif isinstance(val, object):
        if callable(val):
            for attr in dir(val):
                if attr.startswith('__') or attr.startswith('_'):
                    continue
                attrVal = getattr(val, attr)
                if callable(attrVal):
                    setattr(val, attr, log(attrVal, prefix=f"instance {val.__class__.__name__}", printer=printer))
    else:
        print('not matched')
