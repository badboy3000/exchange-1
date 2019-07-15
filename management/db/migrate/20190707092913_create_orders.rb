class CreateOrders < ActiveRecord::Migration[6.0]
  def change
    create_table :orders, comment: '订单' do |t|
      t.bigint :bid_user_id, null: false, comment: '买方'
      t.bigint :ask_user_id, null: false, comment: '卖方'
      t.string :symbol, null: false, comment: '简称 eg BTC_USD'
      t.bigint :fund_id, null: false, comment: '商品'
      t.bigint :bid_order_book_id, null: false, comment: '买方委托单'
      t.bigint :ask_order_book_id, null: false, comment: '卖方委托单'
      t.decimal :volume, default: 0, precision: 32, scale: 16, comment: '量'
      t.decimal :price, default: 0, precision: 32, scale: 16, comment: '价格'
      t.decimal :ask_fee, default: 0, precision: 32, scale: 16, comment: '卖单手续费'
      t.decimal :bid_fee, default: 0, precision: 32, scale: 16, comment: '买单手续费'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :orders, :bid_user_id
    add_index :orders, :ask_user_id
    add_index :orders, :bid_order_book_id
    add_index :orders, :ask_order_book_id
    add_index :orders, :fund_id
  end
end
