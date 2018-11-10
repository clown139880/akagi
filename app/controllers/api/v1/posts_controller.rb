class Api::V1::PostsController < Api::V1::BaseController

  def index
    $posts = Post.paginate(page: params[:page], per_page: 10)
    api_success(response:$posts)
  end

  def create
    @post = User.first.posts.build post_params
    if @post.save
      @post.case.update_attribute(:updated_at,Time.zone.now)
      if params[:sync_to_weibo] == "1"
        weibo_share @post
      end
      api_success(response:"success")
    else
      api_error(error:40001)
    end
  end

  def add_note
    @key = @case.posts.build(content:params[:key], user_id:1)
    @value = @key.values.build(content:params[:value], case_id:@case.id, user_id:1)
    debugger
    unless @key.save and @value.save
      api_error(status:'fail',message:"保存失败")
    end
  end

  def get_note
    @key = @case.posts.find_by(content: params[:key])
    if @key.nil? or @key.values.empty?
      api_error(status:404,message:"关键字无效或没有结果")
    else
      @value = @key.values.first
    end
  end

  def check_note
    @key = @case.posts.find_by(content: params[:key])
    if @key.nil? or @key.values.empty?
      api_error(status:404,message:"关键字无效或没有结果")
    else
      @value = @key.values.first
      if @value.content == params[:value]
        api_error(status:"success",message:"")
      else
        api_error(status:"fail",message:@value.content)
      end
    end
  end

  private
  def post_params
    params.require(:post).permit(:content, :picture, :case_id, :tag_list, photos_attributes:[:image])
  end

  def set_case
    @case = Case.find_by(id: params[:case_id])
    if @case.nil?
      api_error(status:404,message:"没有case")
    end
  end
end
