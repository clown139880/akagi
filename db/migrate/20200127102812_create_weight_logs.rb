class CreateWeightLogs < ActiveRecord::Migration[5.2]
  def change
    create_table :weight_logs, id: :integer do |t|
      t.decimal :weight, precision: 10, scale: 2
      t.integer :user_id, foreign_key: true

      t.timestamps
    end

    add_index :weight_logs, [:user_id, :created_at]
  end
end