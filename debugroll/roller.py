import inspect
from types import ModuleType
import functools
from inspect import stack
from typing import Any, Callable
from debugroll.printer import PrinterProtocol, StdoutPrinter
import os
from dataclasses import dataclass

@dataclass
class LogCtx:
    prefix: str # should change
    path: str
    lineno: int
    name: str
    args: Any
    result: Any

class Roller:
    printer: PrinterProtocol = StdoutPrinter()

    # this should be changeable.
    def fmtmessage(self, ctx: LogCtx) -> str:
        return f"{ctx.prefix}| {ctx.path}:{ctx.lineno} | {ctx.name}| {ctx.args}| {ctx.result}"

    def log(self, func: Callable, prefix:str='') -> Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            result = func(*args, **kwargs)
            # see https://stackoverflow.com/questions/24438976
            caller = stack()[1]
            workdir = os.getcwd()
            callerpath = os.path.relpath(caller.filename, start=workdir)
            ctx = LogCtx(
                prefix=prefix,
                path=callerpath,
                lineno=caller.lineno,
                name=func.__name__,
                args=args,
                result=result,
            )
            self.printer.print(self.fmtmessage(ctx))
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
