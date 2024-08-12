class NoteSerializer < ActiveModel::Serializer
  attributes :id, :name, :description, :memo

  def memo
    "this is additional memo"
  end
end
