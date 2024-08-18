# rails-notes

## Links
- https://guides.rubyonrails.org/getting_started.html
- https://railsguides.jp/getting_started.html

## Commands
### How to create rails app at first.
```bash
# install ruby
brew install rbenv ruby-build
rbenv init # also, need update ~/.zshrc

# install rails to machine
gem install rails

# create app
rails new . -n notes

# start server
bin/rails server

# gen code
# modify config/routes.rb here
bin/rails generate controller Notes index --skip-routes
bin/rails generate model Notes name:string description:text
bin/rails db:migrate:status
bin/rails db:migrate
```

### Local Development with Docker
```bash
docker compose build
docker compose up -d
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
- add docker-compose.yaml
- bin/rails generate controller で view を作らない方法
