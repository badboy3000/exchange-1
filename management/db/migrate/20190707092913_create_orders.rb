class CreateOrders < ActiveRecord::Migration[6.0]
  def change
    create_table :orders, comment: '订单' do |t|
      t.integer :user_id, null: false, comment: '用户'
      t.string :symbol, null: false, comment: '简称 eg BTC_USD'
      t.integer :fund_id, null: false, comment: '商品'
      t.integer :order_book_id, null: false, comment: '关联下单记录'
      t.string :order_type, null: false, comment: '订单类型 市价单market 限价单limit'
      t.string :side, null: false, comment: 'sell or buy'
      t.decimal :volume, default: 0, precision: 32, scale: 16, comment: '量'
      t.decimal :price, default: 0, precision: 32, scale: 16, comment: '价格'
      t.datetime :deleted_at, comment: '删除时间'
      t.decimal :ask_fee, default: 0, precision: 32, scale: 16, comment: '卖单手续费'
      t.decimal :bid_fee, default: 0, precision: 32, scale: 16, comment: '买单手续费'

      t.timestamps
    end

    add_index :orders, :user_id
    add_index :orders, :fund_id
    add_index :orders, :order_book_id
  end
end
