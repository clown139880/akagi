class AddWeixinOpenidToUsers < ActiveRecord::Migration[5.0]
  def change
    add_column :users, :weixin_openid, :string
    add_column :users, :avatar, :string
  end
end
