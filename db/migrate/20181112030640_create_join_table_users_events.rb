class CreateJoinTableUsersEvents < ActiveRecord::Migration[5.2]
  def change
    create_join_table :users, :events do |t|
      t.integer  :type
      t.integer  :status
      t.index [:user_id, :event_id, :type]

      t.timestamps
    end
  end
end
