module Blog::PostsHelper

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
    ol_flag = false;
    markdown(post.content).to_s.lines.each_with_index do |line, index|
      if index < 20 && ol_flag == false
        if line.include?('<ol>')
          concat(sanitize "<br>")
          ol_flag = true;
          next
        end
        concat(sanitize line)
      elsif image_count < 2 and line.include?('img')
        concat(sanitize line)
        image_count = image_count +1
      elsif line.include?('</ol>')
        ol_flag = false;
      end
    end
    if markdown(post.content).to_s.lines.count > 20
      content_tag :p, class:'tiy' do
        link_to '查看全文', [:blog, post], class:"more"
      end
    else
    end
  end

  def get_upload_token()
    return QiniuUploader.get_token_without_key();
  end

end
