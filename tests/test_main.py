from debugroll import debugroll
import json

def test_example():
    out = {'value': ''}
    def printermock(text: str):
        out['value'] = text
    debugroll(json, printer=printermock)
    json.loads('[0]')

    assert out['value'] == "module json| test_main.py:9 | loads| ('[0]',)| [0]"
