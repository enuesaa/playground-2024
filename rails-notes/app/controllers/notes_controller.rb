class NotesController < ApplicationController
  def index
    @notes = Note.all

    # see https://github.com/rails-api/active_model_serializers/issues/1788
    render status: :ok, json: @notes, each_serializer: NoteSerializer
  end

  def show
    @note = Note.find(params[:id])
    render json: @note, serializer: NoteSerializer
  end

  def create
    @note = Note.new(params.require(:note).permit(:name, :description))
  
    if @note.save
      render status: :ok, json: {"success" => true}
    else
      render status: :unprocessable_entity, json: @note.errors
    end
  end

  def update
    @note = Note.find(params[:id])
    if @note.update(params.require(:note).permit(:name, :description))
      render status: :ok, json: {"success" => true}
    else
      render status: :unprocessable_entity, json: @note.errors
    end
  end  

  def destroy
    @note = Note.find(params[:id])
    @note.destroy

    render status: :ok, json: {"success" => true}
  end  
end
