class HealthController < ApplicationController
  def show
    Rails.logger.info "hello this is health controller."
    render_success
  end
end
