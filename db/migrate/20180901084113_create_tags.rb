class CreateTags < ActiveRecord::Migration[5.1]
  def change
    create_table :tags do |t|
      t.string   "name"
      t.timestamps
      t.integer  "taggings_count", default: 0
    end
  end
end
