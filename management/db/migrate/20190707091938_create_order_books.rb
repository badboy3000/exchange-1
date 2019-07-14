class CreateOrderBooks < ActiveRecord::Migration[6.0]
  def change
    create_table :order_books, comment: '下单记录 包括ask and bid' do |t|
      t.bigint :user_id, null: false, comment: '用户'
      t.string :symbol, null: false, comment: '简称 eg BTC_USD'
      t.bigint :fund_id, null: false, comment: '商品'
      t.integer :status, null: false, default: 0, comment: '状态'
      t.string :order_type, null: false, comment: '订单类型 市价单market 限价单limit'
      t.string :side, null: false, comment: 'sell or buy'
      t.decimal :volume, default: 0, precision: 32, scale: 16, comment: '量'
      t.decimal :price, default: 0, precision: 32, scale: 16, comment: '价格'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :order_books, :user_id
    add_index :order_books, :fund_id
  end
end
