class Post < ApplicationRecord

  belongs_to :user
  belongs_to :event

  @loaded = false 

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

  def content
    raw = read_attribute(:content)
    if @loaded 
      return raw
    else 
      self.photos.map do |photo| 
        raw.gsub! photo.origin_url, photo.url
      end
      @loaded = true
      return raw
    end
  end

  def get_avatar
    if self.user
      self.user.avatar
    else
      Settings.nicknames.sample + '.jpg'
    end
  end

  def logo
    self.photos.empty? ? self.event.get_logo : self.photos.last.url
  end

end