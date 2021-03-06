class Photo < ApplicationRecord

  belongs_to :photoable, polymorphic:true, optional: true
  default_scope -> { order(created_at: :desc) }
  scope :logo, -> { where(is_logo: true) }

  before_save :set_key

  def url(style = 'small')
    self.key ? Settings.end_point + self.key + '!' + style : self.origin_url
  end

  def origin_url
    read_attribute(:url)
  end

  def set_key
    keys = self.url.split 'aliyuncs.com/'
    self.key = keys[1]
  end

  def url_small
    self.url('200x200')
  end
end
