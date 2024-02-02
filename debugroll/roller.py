import inspect
from types import ModuleType
import functools
from inspect import getframeinfo, stack
from typing import Any
from debugroll.printer import PrinterProtocol, StdoutPrinter

class Roller:
    printer: PrinterProtocol = StdoutPrinter()

    def log(self, func, prefix=''):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            result = func(*args, **kwargs)
            # see https://stackoverflow.com/questions/24438976
            caller = getframeinfo(stack()[1][0])
            filename = caller.filename.split('/')[-1]
            text = f"{prefix}| {filename}:{caller.lineno} | {func.__name__}| {args}| {result}"
            self.printer.print(text)
            return result
        return wrapper
    
    def inject(self, val: Any) -> None:
        if isinstance(val, ModuleType):
            for attr in dir(val):
                if attr.startswith('__') or attr.startswith('_'):
                    continue
                attrVal = getattr(val, attr)
                if callable(attrVal):
                    setattr(val, attr, self.log(attrVal, prefix=f"module {val.__name__}"))
        elif inspect.isclass(val):
            if callable(val):
                for attr in dir(val):
                    if attr.startswith('__') or attr.startswith('_'):
                        continue
                    attrVal = getattr(val, attr)
                    if callable(attrVal):
                        setattr(val, attr, self.log(attrVal, prefix=f"class {val.__class__.__name__}"))
        elif isinstance(val, object):
            if callable(val):
                for attr in dir(val):
                    if attr.startswith('__') or attr.startswith('_'):
                        continue
                    attrVal = getattr(val, attr)
                    if callable(attrVal):
                        setattr(val, attr, self.log(attrVal, prefix=f"instance {val.__class__.__name__}"))
        else:
            print('not matched')

    def __call__(self, val: Any) -> None:
        self.inject(val)
