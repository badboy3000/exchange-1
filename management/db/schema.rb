# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `rails
# db:schema:load`. When creating a new database, `rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2019_07_13_040614) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "accounts", comment: "币种账户", force: :cascade do |t|
    t.integer "user_id", null: false, comment: "用户"
    t.integer "currency_id", null: false, comment: "币种"
    t.decimal "balance", precision: 32, scale: 16, default: "0.0", comment: "余额"
    t.decimal "locked", precision: 32, scale: 16, default: "0.0", comment: "锁定金额"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["currency_id"], name: "index_accounts_on_currency_id"
    t.index ["user_id"], name: "index_accounts_on_user_id"
  end

  create_table "currencies", comment: "币种", force: :cascade do |t|
    t.string "symbol", null: false, comment: "简称"
    t.decimal "deposit_fee", precision: 32, scale: 16, default: "0.0", comment: "存款手续费"
    t.decimal "withdraw_fee", precision: 32, scale: 16, default: "0.0", comment: "提现手续费"
    t.datetime "deleted_at", comment: "删除时间"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["symbol"], name: "index_currencies_on_symbol", unique: true
  end

  create_table "deposits", comment: "存款记录", force: :cascade do |t|
    t.integer "account_id", null: false, comment: "账户"
    t.integer "currency_id", null: false, comment: "币种"
    t.decimal "amount", precision: 32, scale: 16, default: "0.0", comment: "金额"
    t.decimal "fee", precision: 32, scale: 16, default: "0.0", comment: "手续费"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["account_id"], name: "index_deposits_on_account_id"
    t.index ["currency_id"], name: "index_deposits_on_currency_id"
  end

  create_table "funds", comment: "商品", force: :cascade do |t|
    t.string "name", null: false, comment: "名称"
    t.string "symbol", null: false, comment: "简称 eg BTC_USD"
    t.integer "right_currency_id", null: false, comment: "币种 eg BTC"
    t.integer "left_currency_id", null: false, comment: "币种 eg USD"
    t.decimal "limit_rate", precision: 32, scale: 16, default: "0.0", comment: "限价单利率"
    t.decimal "market_rate", precision: 32, scale: 16, default: "0.0", comment: "市价单利率"
    t.datetime "deleted_at", comment: "删除时间"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["left_currency_id"], name: "index_funds_on_left_currency_id"
    t.index ["right_currency_id"], name: "index_funds_on_right_currency_id"
    t.index ["symbol"], name: "index_funds_on_symbol", unique: true
  end

  create_table "order_books", comment: "下单记录 包括ask and bid", force: :cascade do |t|
    t.integer "user_id", null: false, comment: "用户"
    t.string "symbol", null: false, comment: "简称 eg BTC_USD"
    t.integer "fund_id", null: false, comment: "商品"
    t.integer "status", default: 0, null: false, comment: "状态"
    t.string "order_type", null: false, comment: "订单类型 市价单market 限价单limit"
    t.string "side", null: false, comment: "sell or buy"
    t.decimal "volume", precision: 32, scale: 16, default: "0.0", comment: "量"
    t.decimal "price", precision: 32, scale: 16, default: "0.0", comment: "价格"
    t.datetime "deleted_at", comment: "删除时间"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["fund_id"], name: "index_order_books_on_fund_id"
    t.index ["user_id"], name: "index_order_books_on_user_id"
  end

  create_table "orders", comment: "订单", force: :cascade do |t|
    t.integer "user_id", null: false, comment: "用户"
    t.string "symbol", null: false, comment: "简称 eg BTC_USD"
    t.integer "fund_id", null: false, comment: "商品"
    t.string "order_type", null: false, comment: "订单类型 市价单market 限价单limit"
    t.string "side", null: false, comment: "sell or buy"
    t.decimal "volume", precision: 32, scale: 16, default: "0.0", comment: "量"
    t.decimal "price", precision: 32, scale: 16, default: "0.0", comment: "价格"
    t.decimal "average_price", precision: 32, scale: 16, default: "0.0", comment: "开单均价"
    t.decimal "ask_fee", precision: 32, scale: 16, default: "0.0", comment: "卖单手续费"
    t.decimal "bid_fee", precision: 32, scale: 16, default: "0.0", comment: "买单手续费"
    t.datetime "deleted_at", comment: "删除时间"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["fund_id"], name: "index_orders_on_fund_id"
    t.index ["user_id"], name: "index_orders_on_user_id"
  end

  create_table "trading_records", comment: "成交记录", force: :cascade do |t|
    t.integer "user_id", null: false, comment: "用户"
    t.string "symbol", null: false, comment: "简称 eg BTC_USD"
    t.integer "fund_id", null: false, comment: "商品"
    t.integer "order_book_id", null: false, comment: "关联下单记录"
    t.integer "order_id", null: false, comment: "所属订单"
    t.string "order_type", null: false, comment: "订单类型 市价单market 限价单limit"
    t.string "side", null: false, comment: "sell or buy"
    t.decimal "volume", precision: 32, scale: 16, default: "0.0", comment: "量"
    t.decimal "price", precision: 32, scale: 16, default: "0.0", comment: "价格"
    t.decimal "ask_fee", precision: 32, scale: 16, default: "0.0", comment: "卖单手续费"
    t.decimal "bid_fee", precision: 32, scale: 16, default: "0.0", comment: "买单手续费"
    t.datetime "deleted_at", comment: "删除时间"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["fund_id"], name: "index_trading_records_on_fund_id"
    t.index ["order_book_id"], name: "index_trading_records_on_order_book_id"
    t.index ["order_id"], name: "index_trading_records_on_order_id"
    t.index ["user_id"], name: "index_trading_records_on_user_id"
  end

  create_table "users", comment: "用户", force: :cascade do |t|
    t.string "name", null: false, comment: "用户名"
    t.string "password_digest", null: false, comment: "密码"
    t.string "email", null: false, comment: "有限"
    t.string "role", comment: "角色"
    t.string "address", comment: "地址"
    t.datetime "deleted_at", comment: "删除时间"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["email"], name: "index_users_on_email", unique: true
  end

  create_table "withdraws", comment: "提现记录", force: :cascade do |t|
    t.integer "account_id", null: false, comment: "账户"
    t.integer "currency_id", null: false, comment: "币种"
    t.decimal "amount", precision: 32, scale: 16, default: "0.0", comment: "金额"
    t.decimal "fee", precision: 32, scale: 16, default: "0.0", comment: "手续费"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.index ["account_id"], name: "index_withdraws_on_account_id"
    t.index ["currency_id"], name: "index_withdraws_on_currency_id"
  end

end
