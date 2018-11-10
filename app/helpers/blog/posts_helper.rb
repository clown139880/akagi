module Blog::PostsHelper
  include ActsAsTaggableOn::TagsHelper

  def post_options_for(post)
    if current_user?(uniclown)
      content_tag :div do
        concat(link_to("edit",  [:edit, :blog, post]))
        concat '|'
        concat(link_to("delete", [:blog, post], method: :delete,
                      data:{ confirm:"我再和你确认一遍，#{current_user.name}，你是否要删除这条微博"}))
      end
    end
  end

  def post_form(post)
    if logged_in?
      render partial: 'blog/public/post_form', locals:{ post: post }
    end
  end

  def post_content_for(post)
    image_count = 0
    ol_flag = false
    content = ""
    markdown(post.content).lines.each_with_index do |line, index|
      if index < 20 && ol_flag == false
        if index == 0 && line.include?('<h')
          next
        elsif line.include?('<ol>')
          #concat(sanitize "<br>")
          content += "<br>"
          ol_flag = true
          next
        else
          #concat(sanitize line)
          content += line
        end
      elsif image_count < 2 and line.include?('img')
        #concat(sanitize line)
        content += line
        image_count += 1
      elsif line.include?('</ol>')
        ol_flag = false
      end
    end
    concat(sanitize content)
    if markdown(post.content).to_s.lines.count > 20
      content_tag :p, class:'read-more' do
        link_to 'Read More', [:blog, post], class:"button"
      end
    else
    end
  end

  def get_upload_token()
    return QiniuUploader.get_token_without_key();
  end

  def post_title_for(post, needTitle = false)
    title = markdown(post.content).to_s.lines[0]
    if title.include? '<h'
      return title[4..-7]
    else
      return needTitle ? post.event.title : false;
    end
  end

  def post_show_logo_for(post)
    logo = post.photos.empty? ? post.event.logo : post.photos.first.image
    image_tag logo
  end

end
