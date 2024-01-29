def log(func):
    def wrapper(*args, **kwargs):
        print('hey', func.__name__, args)
        result = func(*args, **kwargs)
        return result
    return wrapper

def apply(module):
    for attr in dir(module):
        val = getattr(module, attr)
        if callable(val):
            setattr(module, attr, log(val))
