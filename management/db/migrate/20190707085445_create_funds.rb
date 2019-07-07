class CreateFunds < ActiveRecord::Migration[6.0]
  def change
    create_table :funds, comment: '商品' do |t|
      t.string :name, null: false, comment: '名称'
      t.string :symbol, null: false, comment: '简称 eg BTC_USD'
      t.integer :right_currency_id, null: false, comment: '币种 eg BTC'
      t.integer :left_currency_id, null: false, comment: '币种 eg USD'
      t.decimal :limit_rate, default: 0, precision: 32, scale: 16, comment: '限价单利率'
      t.decimal :market_rate, default: 0, precision: 32, scale: 16, comment: '市价单利率'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :funds, :symbol, unique: true
    add_index :funds, :right_currency_id
    add_index :funds, :left_currency_id
  end
end
