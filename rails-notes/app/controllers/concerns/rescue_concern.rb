# module は他言語でいう trait 的なもの
# concern は Laravel で見かける Concern と同義。というかその由来がこれ

module RescueConcern
  extend ActiveSupport::Concern

  # たぶん、これが module のなかに定義しているから、読み込み元で実行するよう included を実行している
  included do
    rescue_from ActionController::ParameterMissing, with: :handle_parameter_missing
    rescue_from ActiveRecord::RecordNotFound, with: :handle_record_not_found
  end

  private
    def handle_parameter_missing(e)
      render status: 400, json: { message: e.original_message }
    end

    def handle_record_not_found
      render status: 404, json: { message: "not found" }
    end
end
