module RenderConcern
  extend ActiveSupport::Concern

  private

  def render200
    render status: 200, json: { success: true }
  end

  def render400e(e)
    render status: 400, json: { success: false, error: e.original_message }
  end

  def render404
    render status: 404, json: { success: false, error: "not found" }
  end
  
  def render500
    render status: 500, json: { success: false, error: "Internal Server Error" }
  end
end
