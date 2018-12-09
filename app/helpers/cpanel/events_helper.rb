module Cpanel::EventsHelper

  def event_status_for(event)
    case event.status
    when 1
      "启用"
    when 2
      "关闭"
    else
    end
  end

  def event_status_change_for(event)
    if event.status == 1
      link_to '', [:close, :cpanel, event], class: 'cm-cancle', title: '点击关闭', remote: true
    elsif event.status == 2
      link_to '', [:open, :cpanel, event],  class: 'fa fa-check', title: '点击开启', remote: true
    else

    end
  end

  def select_for_events
    return Event.all.collect{|p| [p.title, p.id]}
  end

  def event_display_for(parent_id)
    unless parent_id.present?
      "display:none;"
    end
  end

  def event_parent_for(parent)
    unless parent.nil?
      parent.title
    end
  end

  def event_is_hot_for(is_hot)
    case is_hot
    when 1
      "热门"
    end
  end

end
