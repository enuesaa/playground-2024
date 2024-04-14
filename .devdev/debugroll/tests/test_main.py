from debugroll import Roller, MockPrinter
import json

def test_example():
    roller = Roller()
    roller.printer = MockPrinter()
    roller(json)
    json.loads('[0]')

    assert roller.printer.out == """tests/test_main.py:8
  code      |     json.loads('[0]')
  arguments | ('[0]',)
  return    | [0]"""
