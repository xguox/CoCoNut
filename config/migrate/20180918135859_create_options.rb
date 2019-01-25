class CreateOptions < ActiveRecord::Migration[5.1]
  def change
    create_table :options do |t|
      t.string   :name
      t.integer  :product_id, index: true
      t.integer  :position
      t.string     :vals
      t.timestamps
    end
  end
end
