class Cpanel::PostsController < Cpanel::BaseController

	before_action :set_post, only: [:edit, :update, :destroy]

	def index
    @posts = Post.paginate(page: params[:page])
  end

	def new
		@events = Event.all
		@post = Post.new
    @oss_uploader = OssUploader.new
	end

  def create
    @post = User.first.posts.build post_params
    if @post.save
      @post.event.update_attribute(:updated_at, Time.zone.now)
      if params[:is_share] == "1"
        weibo_share @post
      end
			flash[:success] = "发表成功!"
    	redirect_to [:cpanel, :posts]
    else
			flash[:notice] = "失败!"
			@events = Event.all
			render new
    end
  end

	def edit
    @oss_uploader = OssUploader.new
	end

	def update
		@post.update_attributes post_params
		redirect_to [:cpanel, :posts], notice:"更新成功"
	end

	def destroy
		@post.delete
		redirect_to [:cpanel, :posts], notice:"删除成功"
	end

	def get_options
    thepost = Post.find_by(id: params[:id])
		if thepost
		  options = [[thepost.id, thepost.title]]
		  thepost.subposts.each do |s|
		    options << [s.id, s.title]
		  end
			options << [0, '重置']
		  render :json => options
		else
			options = []
			post.level(1).each do |s|
		    options << [s.id, s.title]
		  end
			render :json => options
		end
  end


  private
  def post_params
    params.require(:post).permit(:title, :content, :picture, :event_id, :tag_list, photos_attributes:[:url])
  end

	def set_post
		@post = Post.find_by(id: params[:id])
	end

end
