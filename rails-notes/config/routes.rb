Rails.application.routes.draw do
  root "notes#index"

  get "notes", to: "notes#index"
  get "up" => "rails/health#show", as: :rails_health_check
end
