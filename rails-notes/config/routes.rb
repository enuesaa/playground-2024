Rails.application.routes.draw do
  # see https://github.com/rails/rails/blob/main/railties/lib/rails/health_controller.rb#L38
  root "rails/health#show"

  # get "notes", to: "notes#index" 
  # post "notes", to: "notes#create"
  # get "notes/:id", to: "notes#show" #, as: :note  # 名前付きルート. note_path や note_url という関数を生成するらしい
  # put "notes/:id", to: "notes#update"
  # delete "notes/:id", to: "notes#destroy"
  resources :notes, except: [:new, :edit]
end
