#!/usr/bin/env ruby

require 'qiniu'

class QiniuUploader

	class << self

		def customize_put(url, tmp_file)

    end

    def default_put(tmp_file)
      # 构建鉴权对象
      Qiniu.establish_connection! access_key: '0FJbnxEroUXGGEE621NaDZzlXngpRvWeAjV_dFEM',
                                  secret_key: 'SqX854J-4RFDrA0EY3f8UQYdeV-ujEuh3HjO-AIJ'

      # 要上传的空间
      bucket = 'akagi'
			url = 'http://oes1t3t81.bkt.clouddn.com/'
			file_ext = File.extname tmp_file
      # 上传到七牛后保存的文件名
      key = "#{Time.now.strftime("%Y%m%d")}#{SecureRandom.uuid}#{file_ext}"

      # 构建上传策略
      put_policy = Qiniu::Auth::PutPolicy.new(
          bucket, # 存储空间
          key,    # 最终资源名，可省略，即缺省为“创建”语义，设置为 nil 为普通上传
          3600    # token 过期时间，默认为 3600s
      )

      # 生成上传 Token
      uptoken = Qiniu::Auth.generate_uptoken(put_policy)

      # 要上传文件的本地路径
      filePath = tmp_file.path

      # 调用 upload_file 方法上传
      info = Qiniu.upload_file uptoken: uptoken, file: filePath, bucket: bucket, key: key

			if info.nil?
				return nil
			else
				img_url = url + key #+ '?imageMogr2/format/webp/blur/1x0/quality/75|imageslim'
			end
      # 打印文件信息

    end

    def get_token(filename)
			# 构建鉴权对象
			Qiniu.establish_connection! access_key: '0FJbnxEroUXGGEE621NaDZzlXngpRvWeAjV_dFEM',
																	secret_key: 'SqX854J-4RFDrA0EY3f8UQYdeV-ujEuh3HjO-AIJ'

			# 要上传的空间
			bucket = 'akagi'
			url = 'http://oes1t3t81.bkt.clouddn.com/'
			# 上传到七牛后保存的文件名
			if Rails.env.development?
				key = "test/wechat/#{Time.now.strftime("%Y%m%d")}#{filename}"
			else
				key = "wechat/#{Time.now.strftime("%Y%m%d")}#{filename}"
			end
			# 构建上传策略
			put_policy = Qiniu::Auth::PutPolicy.new(
					bucket, # 存储空间
					key,    # 最终资源名，可省略，即缺省为“创建”语义，设置为 nil 为普通上传
					3600    # token 过期时间，默认为 3600s
			)
			img_url = url + key
			# 生成上传 Token
			uptoken = Qiniu::Auth.generate_uptoken(put_policy)
			return key, uptoken, img_url
    end

		def get_token_without_key()
			# 构建鉴权对象
			Qiniu.establish_connection! access_key: '0FJbnxEroUXGGEE621NaDZzlXngpRvWeAjV_dFEM',
																	secret_key: 'SqX854J-4RFDrA0EY3f8UQYdeV-ujEuh3HjO-AIJ'
			# 要上传的空间
			bucket = 'akagi'
			url = 'http://oes1t3t81.bkt.clouddn.com/'
			# 上传到七牛后保存的文件名
			# 构建上传策略
			put_policy = Qiniu::Auth::PutPolicy.new(
					bucket, # 存储空间
					nil,    # 最终资源名，可省略，即缺省为“创建”语义，设置为 nil 为普通上传
					3600    # token 过期时间，默认为 3600s
			)
			# 生成上传 Token
			uptoken = Qiniu::Auth.generate_uptoken(put_policy)
			return uptoken
		end

	end

end
