Rails.application.routes.draw do
  root "notes#index"

  get "notes", to: "notes#index" 
  post "notes", to: "notes#create"
  get "notes/:id", to: "notes#show", as: :note
  put "notes/:id", to: "notes#update"
  delete "notes/:id", to: "notes#destroy"

  # or
  # resources :notes, except: [:new, :edit]

  get "up" => "rails/health#show", as: :rails_health_check
end
