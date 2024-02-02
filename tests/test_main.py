from debugroll import Roller, MockPrinter
import json

def test_example():
    roller = Roller()
    roller.printer = MockPrinter()
    roller(json)
    json.loads('[0]')

    assert roller.printer.out == "tests/test_main.py:8 |     json.loads('[0]') | ('[0]',)| [0]"
