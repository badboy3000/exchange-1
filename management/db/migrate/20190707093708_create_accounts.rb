class CreateAccounts < ActiveRecord::Migration[6.0]
  def change
    create_table :accounts, comment: '币种账户' do |t|
      t.bigint :user_id, null: false, comment: '用户'
      t.bigint :currency_id, null: false, comment: '币种'
      t.decimal :balance, default: 0, precision: 32, scale: 16, comment: '余额'
      t.decimal :locked, default: 0, precision: 32, scale: 16, comment: '锁定金额'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :accounts, :user_id
    add_index :accounts, :currency_id
  end
end
