import langchain
from langchain.text_splitter import CharacterTextSplitter
from langchain_community.document_loaders import PDFMinerLoader

loader = PDFMinerLoader('sample.pdf', concatenate_pages=False)
documents = loader.load()

for doc in documents:
    print(doc)
