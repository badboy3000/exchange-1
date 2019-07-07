class CreateUsers < ActiveRecord::Migration[6.0]
  def change
    create_table :users, comment: '用户' do |t|
      t.string :name, null: false, comment: '用户名'
      t.string :password_digest, null: false, comment: '密码'
      t.string :email, null: false, comment: '有限'
      t.string :role, comment: '角色'
      t.string :address, comment: '地址'
      t.datetime :deleted_at, comment: '删除时间'

      t.timestamps
    end

    add_index :users, :email, unique: true
  end
end
