class ApplicationController < ActionController::Base
  protect_from_forgery with: :null_session

  include RescueConcern

  def render_404
    render json: { error: "Not Found" }, status: :not_found
  end
end
