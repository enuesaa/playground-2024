# Note Controller
#
# see https://railsguides.jp/api_app.html

class NotesController < ApplicationController
  def index
    Rails.logger.info "hello this is notes controller."

    @notes = Note.all

    # see https://github.com/rails-api/active_model_serializers/issues/1788
    render status: :ok, json: @notes, each_serializer: NoteSerializer
  end

  def show
    @note = Note.find(params[:id])
    render json: @note, serializer: NoteSerializer
  end

  def create
    @note = Note.new(note_params)
  
    if @note.save
      render status: :ok, json: {success: true}
    else
      render status: :unprocessable_entity, json: @note.errors
    end
  end

  def update
    @note = Note.find(params[:id])
    if @note.update(note_params)
      render status: :ok, json: {success: true}
    else
      render status: :unprocessable_entity, json: @note.errors
    end
  end  

  def destroy
    @note = Note.find(params[:id])
    @note.destroy

    render status: :ok, json: {success: true}
  end

  private
    # see https://stackoverflow.com/questions/38432270
    def note_params
      params.require(:data).permit(:name, :description)
    end
end
