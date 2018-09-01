class CreateProducts < ActiveRecord::Migration[5.1]
  def change
    create_table :products do |t|
      t.string   "title"
      t.text     "body_html"
      t.datetime "published_at"
      t.string   "vendor"
      t.text     "keywords"
      t.decimal  "price",        precision: 8, scale: 2
      t.string   "slug"
      t.integer  "stock_qty",                            default: 0
      t.integer  "status",                               default: 0
      t.boolean  "hot_sale",                             default: false
      t.boolean  "new_arrival",                          default: true
      t.string   "cover"

      t.timestamps
      t.integer :category_id, index: true
      t.datetime :deleted_at, index: true
    end
  end
end
