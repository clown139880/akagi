class Wechat::PostsController < Wechat::BaseController

  #before_action :log_in
  before_action :set_event


  def new
    @oss_uploader = OssUploader.new
    @post = @event.posts.build(nickname: Settings.nicknames.sample, types:3)
  end

  def edit
    @oss_uploader = OssUploader.new
    @post = Post.find_by(id: params[:id])
    @event = @post.event
  end

  def update
    @post = Post.find_by(id: params[:id])
    @post.update_attributes post_param
    redirect_to [:wechat,  @post.event]
  end

  def create
    @post = @event.posts.build post_param
    #TODO 当前用户
    @post.user = User.first
    if @post.save!
      redirect_to [:wechat,  @event]
    else
      #获取错误信息
      @post.errors.full_messages
      render :new
    end
  end

  def more
    @posts = @event.posts.paginate(page: params[:page], per_page: 5)
    render partial: '/wechat/events/post', collection: @posts, as: :post
  end


  private
  def set_event
    @event = Event.find_by(id: params[:id])
  end

  def post_param
    params['post']['photos_attributes'] = params['uploads']['photos_attributes'] if params['uploads']
    params.require('post').permit(:nickname, :content, :types, photos_attributes:[:url])
  end

end
