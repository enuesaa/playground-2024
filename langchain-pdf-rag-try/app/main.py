from langchain_community.document_loaders import PDFMinerLoader
from langchain_community.vectorstores import FAISS
from langchain.chains import RetrievalQA
from langchain_huggingface import HuggingFaceEmbeddings
from langchain_huggingface import HuggingFacePipeline

# PDF ファイルから文字列情報を読み取る
loader = PDFMinerLoader('sample2.pdf', concatenate_pages=False)
documents = loader.load()

# テキストは普通に取れている
# PDFファイルをパースしているんだろうと思う
for doc in documents:
    print(doc.page_content)

# see https://note.com/ippei_suzuki_us/n/n11e6574c9657

embeddings = HuggingFaceEmbeddings(model_name='sentence-transformers/all-MiniLM-L6-v2')
vectorstore = FAISS.from_documents(documents, embeddings)

retriever = vectorstore.as_retriever()


llm = HuggingFacePipeline.from_model_id(
    model_id="gpt2",
    task="text-generation",
    pipeline_kwargs={
        "max_length": 110,
        "truncation": True,
    },
)
qa_chain = RetrievalQA.from_chain_type(
    llm=llm,
    retriever=retriever,
    # chain_type="stuff"
)
res = qa_chain.invoke("summary")
print(res)
