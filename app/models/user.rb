class User < ApplicationRecord

  has_many :posts, dependent: :destroy
  has_many :cases
  has_many :active_relationships, class_name:"Relationship",
                              foreign_key:"follower_id",
                              dependent: :destroy
  has_many :following, through: :active_relationships, source: :followed
  has_many :passive_relationships, class_name:"Relationship",
                          foreign_key:"followed_id",
                          dependent: :destroy
  has_many :followers, through: :passive_relationships, source: :follower
  attr_accessor :remember_token,:activation_token,:reset_token

  before_save :downcase_email
  #before_create :create_activation_digest

  validates :name,presence:true
  VALID_EMAIL_REGEX = /\A[\w+\-.]+@[a-z\d\-.]+[a-z\d]\.[a-z]+\z/i

  #validates :email,presence:true,length:{ maximum:255 },
                  #format: { with: VALID_EMAIL_REGEX },
                  #uniqueness:true
  #validates :password, presence: true, length: { minimum: 6 }, allow_nil:true
  #has_secure_password



  def decode_nickname
    self.name
  end


  private
  def downcase_email
      self.email.downcase! unless self.email.nil?
  end

  def create_activation_digest
      self.activation_token = User.new_token
      self.activation_digest = User.digest(activation_token)
  end



  class << self
    def create_user(obj)
      nickname = obj[:nickname].present? ? encode_nickname(obj[:nickname]) : ''
      avatar = obj[:headimgurl].present? ? obj[:headimgurl] : ''
      user = User.create!(name: nickname,
                          avatar: avatar,
                          weixin_openid: obj[:openid])

    end

    def encode_nickname(nickname)
      Rumoji.encode(URI.encode(nickname))
    end

  end


end
