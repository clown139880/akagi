class Wechat::EventsController < Wechat::BaseController

  #before_action :log_in
  before_action :set_event, only:[:index, :show]

  def index
    @events = Event.all
  end

  def new
    @event = Event.new(nickname: Settings.nicknames.sample, types:3)
    @oss_uploader = OssUploader.new
  end

  def create
    @event = Event.new event_param
    @event.user = User.first
    @event.save
    redirect_to [:wechat,  @event]
  end

  def show
    @posts = @event.posts.paginate(page: params[:page], per_page: 5)
    @post = @event.posts.build(nickname: Settings.nicknames.sample, types:3)
  end



  private
  def set_event
    @event = Event.find_by(id: params[:id])
  end

  def event_param
    params['event']['photos_attributes'] = params['uploads']['photos_attributes'] if params['uploads']
    params.require(:event).permit(:nickname, :content, :title, :types, photos_attributes:[:url])
  end

end
