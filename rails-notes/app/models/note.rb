class Note < ApplicationRecord
  validates :name, presence: true
  validates :description, length: { maximum: 10 }
end
