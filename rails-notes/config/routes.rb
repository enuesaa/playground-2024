# Routing
#
# see https://railsguides.jp/routing.html

Rails.application.routes.draw do
  root to: "errors#notfound"

  # health
  # rails デフォルトの health controller を使うこともできる
  # see https://github.com/rails/rails/blob/main/railties/lib/rails/health_controller.rb#L38
  get "health", to: "health#show"

  # notes resources.
  # ネストできる. 例: /notes/aaa/tags
  resources :notes, only: [:index, :show, :create, :update, :destroy]

  # 個別に書くこともできる
  # get "notes", to: "notes#index" 
  # post "notes", to: "notes#create"
  # get "notes/:id", to: "notes#show" #, as: :note  # 名前付きルート. note_path や note_url という関数を生成するらしい
  # put "notes/:id", to: "notes#update"
  # delete "notes/:id", to: "notes#destroy"

  match "*path", to: "errors#notfound", via: :all
end
