class Post < ApplicationRecord

  TYPES_OF_SELF    = 1
  TYPES_OF_ANOY    = 2
  TYPES_OF_NONANOY   = 3

  belongs_to :user, optional: true
  belongs_to :case

  has_many :values, class_name: 'Post', foreign_key: 'key_id'
  belongs_to :key, class_name: 'Post', optional: true

  #validates :user_id, presence:true
  validates :content, presence:true
  #mount_uploader :picture, PictureUploader
  accepts_nested_attributes_for :values

  has_many :photos, as: :photoable

  accepts_nested_attributes_for :photos, reject_if: proc{ |photo|
              photo['image'].blank?  }

  default_scope -> { order(created_at: :desc) }
  scope :created_asc, -> { order(created_at: :asc) }


  before_create :default_case
  before_save :default_case

  acts_as_taggable

  def get_nickname
    self.types == Post::TYPES_OF_ANOY ? self.nickname : self.user ? self.user.decode_nickname : "未知"
  end

  def get_avatar
    if self.user.present?
      self.user.avatar
    elsif Settings.nicknames.include? self.nickname
      self.nickname + '.jpg'
    else
      Settings.nicknames.sample + '.jpg'
    end
  end

  def logo
    self.photos.empty? ? self.case.logo : self.photos.first.image
  end

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

  private


  def default_case
    if self.case_id.nil?
      self.case_id = 1
    end
  end

end
