class ApplicationController < ActionController::Base
  protect_from_forgery with: :null_session

  rescue_from ActionController::ParameterMissing, with: :handle_parameter_missing
  rescue_from ActiveRecord::RecordNotFound, with: :handle_record_not_found

  # see https://qiita.com/SoarTec-lab/items/5727b264ed8f21527592
  def handle_parameter_missing(e)
    render status: 400, json: { message: e.original_message }
  end

  def handle_record_not_found
    render status: 404, json: { message: "not found" }
  end
end
