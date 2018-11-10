class Event < ApplicationRecord

  has_many :photos, as: :photoable

  accepts_nested_attributes_for :photos, reject_if: proc{ |photo|
              photo['url'].blank?  }

  has_many :posts, dependent: :destroy
  belongs_to :user, optional: true

  default_scope -> { order(updated_at: :desc) }

  acts_as_taggable # Alias for acts_as_taggable_on :tags

  def get_started
    self.started_at ? self.started_at.strftime('%Y-%m-%d') : ''
  end

  def get_ended
    self.ended_at ? self.ended_at.strftime('%Y-%m-%d') : ''
  end

end