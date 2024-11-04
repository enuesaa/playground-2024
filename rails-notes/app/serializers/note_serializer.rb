# serializer
#
# see https://github.com/rails-api/active_model_serializers/issues/1788

class NoteSerializer < ActiveModel::Serializer
  attributes :id, :name, :description, :memo

  def memo
    "this is additional memo"
  end
end
