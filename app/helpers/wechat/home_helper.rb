module Wechat::HomeHelper

  def footer_nav(str = "")
    home(str) + cases(str)
  end

  def home(str = "")
    cls = "weui_tabbar_item"
    cls = "weui_tabbar_item weui-bar__item_on" if str == "home"
    icon = "gficon-home weui_tabbar_icon"
    icon = "gficon-home-fill weui_tabbar_icon" if str == "home"
    content_tag :a, href: '/wechat', class: cls do
      concat(content_tag 'i', '', class: icon)
      concat(content_tag :p, '首页', class: 'weui_tabbar_label')
    end
  end

  def cases(str = "")
    cls = "weui_tabbar_item"
    cls = "weui_tabbar_item weui-bar__item_on" if str == "cases"
    icon = "gficon-reserve weui_tabbar_icon"
    icon = "gficon-reserve-fill weui_tabbar_icon" if str == "cases"
    content_tag :a, href: '/wechat/cases', class: cls do
      concat(content_tag 'i', '', class: icon)
      concat(content_tag :p, '帖子', class: 'weui_tabbar_label')
    end
  end


  def users(str = "")
    cls = "weui_tabbar_item"
    cls = "weui_tabbar_item weui-bar__item_on " if str == "users"
    icon = "gficon-my weui_tabbar_icon"
    icon = "gficon-my-fill weui_tabbar_icon" if str == "users"
    content_tag :a, href: wechat_user_path(current_user), class: cls do
      concat(content_tag 'i', '', class: icon)
      concat(content_tag :p, '我', class: 'weui_tabbar_label')
    end
  end

end
