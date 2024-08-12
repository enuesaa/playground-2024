class NotesController < ApplicationController
  def index
    # render status: :ok, json: {"a" => "b"}

    @notes = Note.all

    # see https://github.com/rails-api/active_model_serializers/issues/1788
    render status: :ok, json: @notes, each_serializer: NoteSerializer
  end
end
