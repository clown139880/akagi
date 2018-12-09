class Blog::EventsController < Blog::BaseController

  before_action :logged_in_user, only: [:new, :create, :edit, :destroy]

  def new
    if logged_in?
      @event = current_user.events.build
    end
  end

  def index
    @events = Event.paginate(page: params[:page], per_page: 10)
  end

  def show
    @event = Event.find_by(id: params[:id])
    @posts = @event.posts.paginate(page: params[:page], per_page: 10)
    if logged_in?
      @post = current_user.posts.build(event_id: @event.id)
    end
  end

  def create
    @event = current_user.events.build(event_params)
    if params[:info][:bangumi].length > 0
      bangumi(@event,params[:info][:bangumi])
    end
    if @event.save
      @event.father(params[:father_id])
      flash[:success] = "创建成功！"
      redirect_to @event
    else
      render 'new'
    end
  end

  def edit
    @event = Event.find(params[:id])
  end

  def update
    @event = Event.find(params[:id])
    if @event.update_attributes(event_params)
      flash[:success] = "更新成功，请继续履行救世主的义务"
      redirect_to @event
    else
      render 'edit'
    end
  end

  def destroy
    @event = current_user.events.find_by(id: params[:id])
    redirect_to root_url if @event.nil?
    @event.destroy
    flash[:success] = "记录已销毁"
    redirect_to root_url
  end



  private
  def event_params
    params.require(:event).permit(:body,:title,:tag_list)
  end

  def bangumi(newevent,bgm_url)
    #发送get请求到bgmurl
    url = URI(bgm_url)
    body = Net::HTTP.get(url)
    body = Nokogiri::HTML(body)
    body = body.css('div#bangumiInfo')
    newevent.body = body
    #从request中提取信息
    #将提取到的信息保存到event.body中
  end


end
