require 'aliyun/oss'

# ACCESS_KEY/SECRET_KEY is your access credentials
# host: your bucket's data center host, eg: oss-cn-hangzhou.aliyuncs.com
# Details: https://docs.aliyun.com/#/pub/oss/product-documentation/domain-region#menu2
# bucket: your bucket name

class OssUploader
  attr_accessor :access_key, :key, :policy, :signature, :bucket_endpoint

  def initialize()
    @access_key = Settings.oss_access_key
    secret_key = Settings.oss_secret_key
    bucket = Settings.oss_bucket
    host = Settings.oss_host
    @key = 'blog/' + DateTime.current.to_s(:number)[0..5] +'/'

    policy_hash = {
      expiration: 15.minutes.since.strftime('%Y-%m-%dT%H:%M:%S.000Z'),
      conditions: [
        { bucket: bucket }
      ]
    }

    @policy = Aliyun::Oss::Authorization.get_base64_policy(policy_hash)
    @signature = Aliyun::Oss::Authorization.get_policy_signature(secret_key, policy_hash)
    @bucket_endpoint = Aliyun::Oss::Utils.get_endpoint(bucket, host)
  end

  def upload()
    client.bucket_objects.create('image.png', File.new('path/to/image.png'), { 'Content-Type' => 'image/png' })
  end

  def list()
    client.bucket_objects.list()
  end

  def get_policy()
    # policy your policy in hash
    # Return base64 string which can used to fill your form field: policy
    client.get_base64_policy(policy)
    
    # Get Signature for policy
    # Return Signature with policy, can used to fill your form field: Signature
    client.get_policy_signature(SECRET_KEY, policy)
  end
end
