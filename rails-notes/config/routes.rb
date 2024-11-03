Rails.application.routes.draw do
  root "application#render_404"

  # health
  # see https://github.com/rails/rails/blob/main/railties/lib/rails/health_controller.rb#L38
  get "health", to: "rails/health#show"

  # notes
  resources :notes, except: [:new, :edit]
  # get "notes", to: "notes#index" 
  # post "notes", to: "notes#create"
  # get "notes/:id", to: "notes#show" #, as: :note  # 名前付きルート. note_path や note_url という関数を生成するらしい
  # put "notes/:id", to: "notes#update"
  # delete "notes/:id", to: "notes#destroy"

  match "*path", to: "application#render_404", via: :all
end
