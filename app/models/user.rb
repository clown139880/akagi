class User < ApplicationRecord
  has_secure_password
  has_many :posts, dependent: :destroy
  #TODO: 关注事件，事件邀请，事件追踪

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