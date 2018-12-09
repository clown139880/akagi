module Cpanel::BaseHelper

  def top_c_navbar(str)
    top_c_home(str) +  top_c_event(str) + top_c_post(str)
  end

  def top_c_home(str)
    cls = "home" == str ? "active" : ""
    content_tag :li, class: cls do
      link_to "Home", "/cpanel/"
    end
  end


  def top_c_event(str)
    cls = "events" == str ? "active" : ""
    content_tag :li, class: cls do
      link_to "地球图书馆", [:cpanel, :events]
    end
  end

  def top_c_post(str)
    cls = "posts" == str ? "active" : ""
    content_tag :li, class: cls do
      link_to "文库", [:cpanel, :posts]
    end
  end


  def active_sidebar_class(path, alias_path = "")
    current_page?(path) ? 'active' : ''
  end

end
