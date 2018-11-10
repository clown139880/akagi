class RefactorEverthing < ActiveRecord::Migration[5.0]
  def change
    rename_table :cases, :events
    rename_column :posts, :case_id, :event_id
    remove_column :posts, :picture
    add_column :posts, :is_synced, :smallint
    add_column :posts, :visit_count, :int
    rename_column :photos, :image, :url
  end
end
