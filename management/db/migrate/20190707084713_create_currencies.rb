class CreateCurrencies < ActiveRecord::Migration[6.0]
  def change
    create_table :currencies, comment: '币种' do |t|
      t.string :symbol, null: false, comment: '简称'
      t.decimal :deposit_fee, default: 0, precision: 32, scale: 16, comment: '存款手续费'
      t.decimal :withdraw_fee, default: 0, precision: 32, scale: 16, comment: '提现手续费'
      t.datetime :deleted_at, comment: '删除时间'
      
      t.timestamps
    end

    add_index :currencies, :symbol, unique: true
  end
end
