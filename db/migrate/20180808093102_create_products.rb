class CreateProducts < ActiveRecord::Migration[5.1]
  def change
    create_table :products do |t|
      t.string :name
      t.string :sku
      t.timestamps

      t.datetime :deleted_at
    end
  end
end
