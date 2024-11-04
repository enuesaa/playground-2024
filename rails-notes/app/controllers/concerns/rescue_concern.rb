# module は他言語でいう trait 的なもの
# concern は Laravel で見かける Concern と同義。というかその由来がこれ

module RescueConcern
  extend ActiveSupport::Concern

  # たぶん、これが module のなかに定義しているから、読み込み元で実行するよう included を実行している
  included do
    rescue_from ActionController::ParameterMissing, with: :render400e
    rescue_from ActiveRecord::RecordNotFound, with: :render404
  end
end
