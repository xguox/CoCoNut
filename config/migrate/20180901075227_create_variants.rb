class CreateVariants < ActiveRecord::Migration[5.1]
  def change
    create_table :variants do |t|
      t.decimal  "price",      precision: 8, scale: 2, default: 0.0
      t.string   "sku", index: true
      t.integer  "stock"
      t.integer  "position"
      t.integer  "product_id", index: true
      t.string   "option1"
      t.string   "option2"
      t.string   "option3"
      t.boolean  "is_default", default: false
      t.timestamps
      t.datetime :deleted_at
    end
  end
end
