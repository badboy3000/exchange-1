class CreateDeposits < ActiveRecord::Migration[6.0]
  def change
    create_table :deposits, comment: '存款记录' do |t|
      t.integer :account_id, null: false, comment: '账户'
      t.integer :currency_id, null: false, comment: '币种'
      t.decimal :amount, default: 0, precision: 32, scale: 16, comment: '金额'
      t.decimal :fee, default: 0, precision: 32, scale: 16, comment: '手续费'

      t.timestamps
    end

    add_index :deposits, :account_id
    add_index :deposits, :currency_id
  end
end
