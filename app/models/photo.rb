class Photo < ApplicationRecord

  belongs_to :photoable, polymorphic:true, optional: true
  default_scope -> { order(created_at: :desc) }
  scope :logo, -> { where(is_logo: true) }

  def url
    self.key ? Settings.end_point + self.key : self.origin_url
  end

  def origin_url
    read_attribute(:url)
  end

  def url_small
    self.key ? self.url + '!200x200' : self.origin_url
  end
end
