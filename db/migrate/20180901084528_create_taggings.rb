class CreateTaggings < ActiveRecord::Migration[5.1]
  def change
    create_table :taggings do |t|
      t.integer  "product_id", index: true
      t.integer  "tag_id", index: true
      t.timestamps
    end
  end
end
