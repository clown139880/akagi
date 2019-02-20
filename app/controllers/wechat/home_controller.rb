class Wechat::HomeController < Wechat::BaseController

  #before_action :log_in

  def index
    @events = Event.limit(5)
  end

  private
  def set_event
    @event = Event.find_by(id: params[:id])
  end

  def reply_param
    params.require('post').permit(:nickname,:content)
  end

end
