module RenderConcern
  extend ActiveSupport::Concern

  private

  def render_success
    render status: 200, json: { success: true }
  end
end
