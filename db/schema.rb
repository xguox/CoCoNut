# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20180905133152) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "categories", force: :cascade do |t|
    t.string "name"
    t.string "slug"
    t.string "ancestry"
    t.integer "ancestry_depth"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.datetime "deleted_at"
    t.index ["ancestry"], name: "index_categories_on_ancestry"
    t.index ["ancestry_depth"], name: "index_categories_on_ancestry_depth"
    t.index ["name"], name: "index_categories_on_name"
    t.index ["slug"], name: "index_categories_on_slug"
  end

  create_table "products", force: :cascade do |t|
    t.string "title"
    t.text "body_html"
    t.datetime "published_at"
    t.string "vendor"
    t.text "keywords"
    t.decimal "price", precision: 8, scale: 2, default: "0.0"
    t.string "slug"
    t.integer "stock_qty", default: 0
    t.integer "status", default: 0
    t.boolean "hot_sale", default: false
    t.boolean "new_arrival", default: true
    t.string "cover"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.integer "category_id"
    t.datetime "deleted_at"
    t.index ["category_id"], name: "index_products_on_category_id"
    t.index ["deleted_at"], name: "index_products_on_deleted_at"
  end

  create_table "taggings", force: :cascade do |t|
    t.integer "product_id"
    t.integer "tag_id"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["product_id"], name: "index_taggings_on_product_id"
    t.index ["tag_id"], name: "index_taggings_on_tag_id"
  end

  create_table "tags", force: :cascade do |t|
    t.string "name"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.integer "taggings_count", default: 0
  end

  create_table "users", force: :cascade do |t|
    t.string "email"
    t.string "username"
    t.string "password_digest"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.datetime "deleted_at"
    t.index ["email"], name: "index_users_on_email"
    t.index ["username"], name: "index_users_on_username"
  end

  create_table "variants", force: :cascade do |t|
    t.decimal "price", precision: 8, scale: 2, default: "0.0"
    t.string "sku"
    t.integer "stock"
    t.integer "position"
    t.integer "product_id"
    t.string "option1"
    t.string "option2"
    t.string "option3"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.datetime "deleted_at"
  end

end
