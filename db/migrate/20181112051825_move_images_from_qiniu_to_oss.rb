#require 'csv'

class MoveImagesFromQiniuToOss < ActiveRecord::Migration[5.2]
  def up
    remove_column :photos, :title
    add_index :photos, [:photoable_type, :photoable_id]
    add_column :photos, :key, :string, limit: 150, :after => 'id'
    #oes1t3t81 phv6tif06
    puts'「我要开始导出数据了」'
    posts = Post.where('content like "%http://phv6tif06.bkt.clouddn.com%"')
    urls = []
    posts.each do |post|
      urls += post.content.scan(/http:\/\/phv6tif06\.bkt\.clouddn\.com[^\)]+(?:\))/).map do |url|
        Photo.new({
          url: url[0..-2],
          created_at: post.updated_at,
          updated_at: post.updated_at,
          photoable_id: post.id,
          photoable_type: 'Post'
        })
      end
    end
    Event.where.not(:logo => nil).each do |event|
      urls += [Photo.new({
        url: event.logo,
        created_at: event.updated_at,
        updated_at: event.updated_at,
        photoable_id: event.id,
        photoable_type: 'Event',
        is_logo: true
      })]
    end
    saveKeyFromUrl(Photo.all + urls)
    Post.all.update :user_id => 1
    Photo.all.where(:photoable_type => 'Case').upate :photoable_type => 'Event'
    remove_column :events, :logo
  end

  def export(photos)
    CSV.open('data.csv', 'w', write_headers: true, headers: ['Id', 'OriginUrl', 'NewUrl', 'Id', 'Type']) do |csv|
      photos.each do |photo|
        newUrl = handle(photo.url)
        csv << [photo.id, photo.url, newUrl, photo.photoable_id,photo.photoable_type] #数据内容
      end
    end
  end

  def saveKeyFromUrl(photos)
    photos.each do |photo|
      if photo.url.scan(/clouddn/).length > 0
        photo.key = handle(photo.url)
      end
      photo.save
    end
  end

  def handle(url)
    url = url.gsub(')', '')
    url = url[33..(url.index('?') ? url.index('?') : 0)-1]
    if url.count('/') == 0
    ## 把根目录的文件移动到blog/2018中
      url = 'blog/2018/' + url      
    else
      url = url.gsub(/\/blog/, 'blog/201812')
      url = url.gsub(/201712/, 'blog/201712')
      url = url.gsub(/wechat/, 'wechat/201812')
    end
    return url
  end

  def down
    add_column :photos, :title ,:string
    add_column :events, :logo ,:string
    remove_column :photos, :key
    remove_index :photos, [:photoable_type, :photoable_id]
  end
end
