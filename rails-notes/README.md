# rails-notes
## Links
- https://guides.rubyonrails.org/getting_started.html
- https://railsguides.jp/getting_started.html

## Commands
```bash
brew install rbenv ruby-build
rbenv init # also, need update ~/.zshrc
rails new . -n notes # create app
bin/rails server
# modify config/routes.rb here
bin/rails generate controller Notes index --skip-routes
bin/rails generate model Notes name:string description:text
bin/rails db:migrate:status
bin/rails db:migrate
```

## Versions
### ruby
3.3.0
### gem
3.5.3
### rails
7.1.2 (latest)

## メモ
- vscode に lsp を入れた
- import 的なものがどうなっているのか分からない. グローバルに Note クラス?が存在する?
- view で静的解析みたいなのできないのか. 存在しないプロパティを指定しても vscode でエラー表示にならず戸惑っている
