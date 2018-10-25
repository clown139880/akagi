class Case < ApplicationRecord

  #self.per_page = 20

  TYPES_OF_SELF   = 1    #别人来看自己的
  TYPES_OF_SHARE  = 2   #自己给别人看的
  TYPES_OF_ANOY   = 3  #匿名的
  TYPES_OF_NONANOY = 4  #非匿名的

  STATUS_OF_NOT_START = 0

  has_many :subcases, class_name: 'Case', foreign_key: 'parent_id'
  belongs_to :parent, class_name: 'Case', optional: true

  has_many :photos, as: :photoable

  accepts_nested_attributes_for :photos, reject_if: proc{ |photo|
              photo['image'].blank?  }

  has_many :posts, dependent: :destroy
  belongs_to :user, optional: true

  default_scope -> { order(updated_at: :desc) }

  scope :level, -> (level){ where(level: level) }

  scope :bbs, -> { where(types: [2,3]) }

  scope :by_types, -> (types){ where(types: types) }

  acts_as_taggable # Alias for acts_as_taggable_on :tags

  def get_family
    family = [self]
    child = self
    (self.level-1).times do
      child = child.parent
      family.unshift child
    end
    family
  end

  def get_username
    self.types == Case::TYPES_OF_ANOY ? self.nickname : User.first.name
  end

  def get_avatar
    if self.types == Case::TYPES_OF_NONANOY and self.user.avatar.present?
      self.user.avatar
    elsif Settings.nicknames.include? self.nickname
      self.nickname + '.jpg'
    else
      Settings.nicknames.sample + '.jpg'
    end
  end

  def logo
    if self.photos.empty?
      return 'http://oes1t3t81.bkt.clouddn.com//blog/1511962137717.JPG'
    else
      self.photos.first.image
    end
  end

  #重构
  def status_str
    case self.status
    when 0
      '未开始'
    else
      '未知'
    end
  end

  def types_str
    case self.types
    when 0
      0
    else
      self.types
    end
  end

  def get_started
    self.started_at ? self.started_at.strftime('%Y-%m-%d') : '无'
  end

  def get_ended
    self.ended_at ? self.ended_at.strftime('%Y-%m-%d') : '无'
  end

end
