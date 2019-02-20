class Event < ApplicationRecord

  has_many :photos, as: :photoable
  has_and_belongs_to_many :users

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

  def status_str
    case self.status
    when 0
      'unknow'
    else 
      'unknow'
    end
  end

  def types_str
    case self.types
    when 0
      'unknow'
    else 
      'unknow'
    end
  end

  def logo
    self.photos.logo.first ? self.photos.logo.first.url : ''
  end

  def get_owner_name
    self.user ? self.user.name : 'admin'
  end

  def get_avatar
    self.user ? self.user.avatar : 'http://tvax4.sinaimg.cn/crop.136.1.388.388.1024/6490ca36ly8flf7tqek7fj20go0gk0w7.jpg'
  end

  def get_logo
    self.photos.empty? ? Settings.default_event_logo : self.photos.logo.first ? self.photos.logo.first.url : self.photos.last.url
  end

end