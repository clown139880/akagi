module Wechat::EventsHelper

  def event_banner_for event
    if event.photos.empty?
      image_tag 'table.jpg'
    else
      image_tag event.photos.first.url
    end
  end

  def wechat_desc_for_event event
    if event.posts.count == 0
      desc = event.get_owner_name + '发布了'+ event.title
    else
      desc = event.posts.first.get_username + '发表了新的回复'
    end
  end

end
