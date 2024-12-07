# GitHub Actions self hosted runner with CodeBuild

- GitHub Actions のセルフホステッドランナーを CodeBuild で動かした
- CodeConnections と統合しており、GitHub 接続経由で動いた
- 実際に Actions を動かすと、
  1. CodeBuild へ通知 (webhook)
  2. CodeBuild が GitHub Actions Runner をセットアップ
     たぶんインターネットからダウンロードしている
     see https://docs.github.com/ja/actions/hosting-your-own-runners/managing-self-hosted-runners/adding-self-hosted-runners
  3. Workflow を実行
- おもしろいのは 2. のログは CodeBuild に吐き出されて、3. のログは GitHub に吐き出されること。
  - だからプラットフォームエンジニアリングの観点では、インフラとアプリが分離されており、導入しやすい
  - 一方でトラブルシューティングは面倒
- CodeBuild が buildspec で GitHub Actions をサポートしたのは、たぶんセルフホステッドランナーへの対応も関係しているな。内部的にセルフホステッドランナーを使ってそう。

## 手順
1. CodeConnections で GitHub へ接続 (GitHub App をインストール)
2. CodeBuild Project を作成
3. CodeBuild Project の編集ページから「デフォルトソース認証情報」の編集ページへ飛び、1. で作成したアプリを選択し保存
   この設定だけ、まだ terraform でいじれない
4. CodeBuild Project Webhook を作成
   GitHub Repository で webhook の設定がされる

これで完了

## ユースケース
### 1. データを日本国内に置きたいとき
GitHub Actions は北米で実行されるので、日本で実行させたいときはいいかも。
ただ、Workflow の実行ログは GitHub に送信されちゃうので、その意味では、あまり意味がないかも。
Workflow を作成するとき、どうしても printf とかすると思うんだよなあ。

### 2. VPC の内部で実行させたいとき
メインターゲットはこれだと思う。DB への接続だとか VPC 内部で実行が必要なタスクはいくつかあり、この需要は確実にある。
ただし前述の通りログは GitHub に送信されちゃうし、その意味ではどうなのかな。ちょっと懸念点だと思う。

### 3. ECR にホストしている Docker イメージを使用したいとき
このケースはある気がする。例えば監査目的とかで、ログを全部保存しておきたい際は、ECR にそういうセットアップ済みの Docker イメージを登録しておいて、CodeBuild Project で使えばいい。普通の GitHub Actions だとなかなか難しいと思う。GitHub API をポーリングすればできるのかな

## Links
- https://qiita.com/k-kojima-yumemi/items/573bda88d0fb607b3224
