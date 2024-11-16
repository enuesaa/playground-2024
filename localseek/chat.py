from langchain_ollama import ChatOllama
from langchain_core.runnables import RunnablePassthrough
from langchain_core.prompts import ChatPromptTemplate
from localseek.vd import connect

# see https://python.langchain.com/docs/tutorials/local_rag/#qa-with-retrieval
def chat():
    store = connect()
    retriever = store.as_retriever()

    prompt = ChatPromptTemplate.from_template("""
あなたは質問応答タスクのアシスタントです。以下の取得したコンテキストを使用して質問に答えてください。答えがわからない場合は「わかりません」とだけ答えてください。最大3文で簡潔に答えてください。

<context>
{context}
</context>

Answer the following question:

{question}
""")

    llm = ChatOllama(model='llama3.2')

    qa = (
        {"context": retriever, "question": RunnablePassthrough()}
        | prompt
        | llm
    )
    res = qa.invoke("os に依存しそうなパッケージを教えて")
    print(res)
