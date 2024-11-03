class ErrorsController < ApplicationController
  def notfound
    render status: 404, json: { error: "Not Found" }
  end
end
