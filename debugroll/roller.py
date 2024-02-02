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
    code_context: str
    name: str
    args: Any
    result: Any

class Roller:
    printer: PrinterProtocol = StdoutPrinter()

    # this should be changeable.
    def fmtmessage(self, ctx: LogCtx) -> str:
        # TODO
        code = ctx.code_context.replace('\n', '')
        return f"{ctx.path}:{ctx.lineno} | {code} | {ctx.args}| {ctx.result}"

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
                code_context=caller.code_context[0] if caller.code_context is not None and len(caller.code_context) > 0 else '',
                name=func.__name__,
                args=args,
                result=result,
            )
            self.printer.print(self.fmtmessage(ctx))
            return result
        return wrapper
    
    def inject_module(self, val: ModuleType) -> None:
        for attr in dir(val):
            if attr.startswith('__') or attr.startswith('_'):
                continue
            attrVal = getattr(val, attr)
            if callable(attrVal):
                setattr(val, attr, self.log(attrVal, prefix=val.__name__))

    # is this true?
    def inject_callable(self, val: Callable) -> None:
        for attr in dir(val):
            if attr.startswith('__') or attr.startswith('_'):
                continue
            attrVal = getattr(val, attr)
            if callable(attrVal):
                setattr(val, attr, self.log(attrVal, prefix=val.__class__.__name__))

    def inject(self, val: Any) -> None:
        if isinstance(val, ModuleType):
           self.inject_module(val)
        elif inspect.isclass(val) and callable(val):
            self.inject_callable(val)
        elif isinstance(val, object) and callable(val):
            self.inject_callable(val)
        else:
            print('not matched')

    def __call__(self, val: Any) -> None:
        self.inject(val)
