ActiveRecord::Schema[7.1].define(version: 2024_01_14_103921) do
  create_table "notes", force: :cascade do |t|
    t.string "name"
    t.text "description"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end
end
