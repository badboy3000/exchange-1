class CreateOrders < ActiveRecord::Migration[6.0]
  def change
    create_table :orders, comment: '订单' do |t|
      t.bigint :user_id, null: false, comment: '用户'
      t.string :symbol, null: false, comment: '简称 eg BTC_USD'
      t.bigint :fund_id, null: false, comment: '商品'
      t.bigint :order_book_id, null: false, comment: '委托单'
      t.bigint :other_side_order_book_id, null: false, comment: '对方订单簿'
      t.bigint :other_side_order_id, null: false, comment: '对方成交记录'
      t.string :order_type, null: false, comment: '订单类型 市价单market 限价单limit'
      t.string :side, null: false, comment: 'sell or buy'
      t.decimal :volume, default: 0, precision: 32, scale: 16, comment: '量'
      t.decimal :price, default: 0, precision: 32, scale: 16, comment: '价格'
      t.decimal :average_price, default: 0, precision: 32, scale: 16, comment: '开单均价'
      t.decimal :ask_fee, default: 0, precision: 32, scale: 16, comment: '卖单手续费'
      t.decimal :bid_fee, default: 0, precision: 32, scale: 16, comment: '买单手续费'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :orders, :user_id
    add_index :orders, :fund_id
  end
end
