class CreateVariants < ActiveRecord::Migration[5.1]
  def change
    create_table :variants do |t|
      t.decimal  "price",      precision: 8, scale: 2
      t.string   "sku"
      t.integer  "stock"
      t.integer  "position"
      t.integer  "product_id"
      t.string   "option1"
      t.string   "option2"
      t.string   "option3"

      t.timestamps
      t.datetime :deleted_at
    end
  end
end
