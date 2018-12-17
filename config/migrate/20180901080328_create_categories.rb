class CreateCategories < ActiveRecord::Migration[5.1]
  def change
    create_table :categories do |t|
      t.string   "name", index: true
      t.string   "slug", index: true
      t.string   "ancestry", index: true
      t.integer  "ancestry_depth", index: true

      t.timestamps
      t.datetime :deleted_at
    end
  end
end
