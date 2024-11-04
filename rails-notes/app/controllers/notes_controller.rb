# Note Controller
#
# see https://railsguides.jp/api_app.html

class NotesController < ApplicationController
  def index
    @notes = Note.all
    render json: @notes, each_serializer: NoteSerializer
  end

  # レコードが存在しなかったら ActiveRecord::RecordNotFound が投げられ RescueConcern でキャッチしている
  def show
    @note = Note.find(params[:id])
    render json: @note, serializer: NoteSerializer
  end

  def create
    @note = Note.new(reqbody)

    unless @note.save
      return render status: 422, json: @note.errors
    end
    render_success
  end

  def update
    @note = Note.find(params[:id])

    unless @note.update(reqbody)
      return render status: 422, json: @note.errors
    end
    render_success
  end

  def destroy
    @note = Note.find(params[:id])
    
    unless @note.destroy
      return render status: 500, json: {error: "Internal Server Error"}
    end
    render_success
  end

  private

  # see https://stackoverflow.com/questions/38432270
  def reqbody
    params.require(:data).permit(:name, :description)
  end
end
