class HealthController < ApplicationController
  def show
    render status: 200, json: { health: "ok" }
  end
end
