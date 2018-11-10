class Post < ApplicationRecord

  belongs_to :user
  belongs_to :event

  validates :content, presence:true

  has_many :photos, as: :photoable

  accepts_nested_attributes_for :photos, reject_if: proc{ |photo|
              photo['url'].blank?  }

  default_scope -> { order(created_at: :desc) }
  scope :created_asc, -> { order(created_at: :asc) }

  acts_as_taggable

  def get_username
    self.user ? self.user.name : ""
  end

  def get_avatar
    if self.user
      self.user.avatar
    else
      Settings.nicknames.sample + '.jpg'
    end
  end

  def logo
    self.photos.empty? ? self.event.logo : self.photos.first.url
  end

end