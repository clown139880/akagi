class RefactorEverthing < ActiveRecord::Migration[5.0]
  def change
    rename_table :cases, :events
    rename_column :posts, :case_id, :event_id
    remove_column :posts, :picture
    add_column :photos, :is_logo, :boolean, :default => false, :after => :id
    add_column :posts, :is_synced, :boolean, :default => false, :before => :created_at
    add_column :posts, :visit_count, :int, :default => 0, :after => :id
    rename_column :photos, :image, :url
  end
end
