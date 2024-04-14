from typing import Protocol

class PrinterProtocol(Protocol):
    def print(self, text: str) -> None:
        pass

class StdoutPrinter:
    def print(self, text: str) -> None:
        print(text)

class MockPrinter:
    out: str = ''

    def print(self, text: str) -> None:
        self.out += text
