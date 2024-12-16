from markitdown import MarkItDown

def main():
    markitdown = MarkItDown()
    result = markitdown.convert("test.docx")
    print(result.text_content)
