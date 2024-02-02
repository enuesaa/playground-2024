from debugroll import Roller, MockPrinter
import json

def test_example():
    roller = Roller()
    roller.printer = MockPrinter()
    roller(json)
    json.loads('[0]')

    assert roller.printer.out == "module json| test_main.py:8 | loads| ('[0]',)| [0]"
