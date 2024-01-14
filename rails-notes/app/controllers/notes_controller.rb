class NotesController < ApplicationController
  def index
    @notes = Note.all

    @c = Aaa.all
  end
end
