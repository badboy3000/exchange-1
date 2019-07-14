class CreateWithdraws < ActiveRecord::Migration[6.0]
  def change
    create_table :withdraws, comment: '提现记录' do |t|
      t.bigint :account_id, null: false, comment: '账户'
      t.bigint :currency_id, null: false, comment: '币种'
      t.decimal :amount, default: 0, precision: 32, scale: 16, comment: '金额'
      t.decimal :fee, default: 0, precision: 32, scale: 16, comment: '手续费'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :withdraws, :account_id
    add_index :withdraws, :currency_id
  end
end
