class HealthController < ApplicationController
  def show
    Rails.logger.info "hello this is health controller."

    render status: 200, json: { health: "ok" }
  end
end
