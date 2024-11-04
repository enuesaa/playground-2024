# App Controller
#
# see https://railsguides.jp/api_app.html

class ApplicationController < ActionController::Base
  protect_from_forgery with: :null_session

  include RescueConcern, RenderConcern
end
