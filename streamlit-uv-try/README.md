# streamlit-uv
Streamlit へ uv でトライ

## Steamlit
- やっぱ Snowflake が作っているだけあって完成度が高いし直感的。
- PoC 用途には十分に思える
  - 社内向けの小さいアプリケーションは、もうこれでいいんじゃないか。
- Apache License 2.0
- Component の実装は React (Class Component) + Emotion っぽい
  - https://github.com/streamlit/streamlit/blob/develop/frontend/app/src/components/Sidebar/Sidebar.tsx
  - たぶん長期間メンテナンスしてくれてるんだろうなあ。

### Links
- https://zenn.dev/whitphx/articles/streamlit-realtime-cv-app
- https://qiita.com/tamura__246/items/366b5581c03dd74f4508
- https://docs.streamlit.io/develop/api-reference/status
- https://github.com/streamlit/streamlit

## uv
- rye のメンテナンスをしている組織が、別に作っているツール
- poetry の領域がメインのターゲットに見えているが、uv pip というサブコマンドを実装していたり、カバー範囲が広い
- uv init でプロジェクトをセットアップしてくれる
- __init__.py を下手に消すとなんかエラーが出るので、触らないほうがいい
- 日本語情報が少なくてトライに時間がかかった

### Commands
```bash
# this create `streamlit-uv-try` dir which contains `.venv` and `src` dir.
uv init streamlit-uv-try
cd streamlit-uv-try
uv add streamlit
uv sync
uv run example.py
uv run streamlit run src/streamlit_uv_try/main.py
 ```

### Links
- https://github.com/astral-sh/uv
- https://docs.astral.sh/uv/    ... 公式ドキュメント
- https://gihyo.jp/article/2024/03/monthly-python-2403
- https://zenn.dev/turing_motors/articles/0f1a764d14f581
