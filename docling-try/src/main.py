from docling.document_converter import DocumentConverter

source = "https://arxiv.org/pdf/2408.09869"

converter = DocumentConverter()
result = converter.convert(source) # download something here.

markdown = result.document.export_to_markdown()

with open('result.md', 'w') as f:
    f.write(markdown)
