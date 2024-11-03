# rails-notes

## Links
- https://guides.rubyonrails.org/getting_started.html
- https://railsguides.jp/getting_started.html

## Commands
### How to create rails app
```bash
# install ruby
brew install rbenv ruby-build
rbenv init # also, need update ~/.zshrc

# install rails to machine
gem install rails

# create app
rails new . -n notes

# migrate. see storage/development.sqlite3
bin/rails db:migrate:status
bin/rails db:migrate

# start server
bin/rails server

# gen code
bin/rails generate controller Notes index --skip-routes
bin/rails generate model Notes name:string description:text
```

### How to setup Local Dev Environment
```bash
bin/rails db:migrate
bin/rails server
```

## Versions
### ruby
3.3.0
### gem
3.5.3
### bundler
2.5.3
### rails
7.1.2 (latest)

## メモ
- vscode に lsp を入れた
- import 的なものがどうなっているのか分からない. グローバルに Note クラス?が存在する?
- view で静的解析みたいなのできないのか. 存在しないプロパティを指定しても vscode でエラー表示にならず戸惑っている
  - rubocop で phpstan & prettier 的なものができるらしい

## TODO
- bin/rails generate controller で view を作らない方法
