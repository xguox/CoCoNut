class CreateUsers < ActiveRecord::Migration[5.1]
  def change
    create_table :users do |t|
      t.string   "email", index: true
      t.string   "username", index: true
      t.string   "password_digest"
      t.timestamps

      t.datetime "deleted_at"
    end
  end
end
