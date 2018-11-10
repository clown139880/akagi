#namespace = "#{Settings.app_alise}:wechat"
#if Rails.env.development?
#  
#  redis = Redis.new(:host => Settings.redis_host, :port => Settings.redis_port, :db => Settings.redis_db)
#else
#  redis = Redis.new(url: ENV["REDIS_URL"])
#end
#
#exist_keys = redis.keys("#{namespace}:*")
#exist_keys.each{|key|redis.del(key)}
#redis = Redis::Namespace.new("#{namespace}", :redis => redis)
#
#wechat_timeout = 1000
#
#WeixinAuthorize.configure do |config|
#  config.redis = redis
#  config.rest_client_options = {timeout: wechat_timeout, open_timeout: wechat_timeout, verify_ssl: false }
#end
#
#$wechat_client ||= WeixinAuthorize::Client.new(Settings.wechat_appid, Settings.wechat_appsecret)
#$wechat_client.is_valid?
