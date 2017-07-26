class Blog::UsersController < Blog::BaseController
    before_action :logged_in_user, only:[:destroy, :index,:edit, :update, :following, :followers]
    before_action :correct_user, only:[:edit, :update]
    before_action :admin_user, only:[:destroy]
  def new
      @user = User.new
  end

  def index
      @users = User.where(activated:true).paginate(page:params[:page])
  end

  def show
      @user = User.find(params[:id])
      redirect_to root_url and return unless @user.activated
      @microposts = @user.microposts.paginate(page: params[:page])
  end

  def create
      @user = User.new(user_params)
      if @user.save
          @user.send_activation_email
          flash[:info] = "确认邮箱中的召唤链接"
          redirect_to root_url
      else
          render 'new'
      end
  end





  def edit
      @user = User.find(params[:id])
  end

  def update
      @user = User.find(params[:id])
      if @user.update_attributes(user_params)
          flash[:success] = "重铸成功，#{@user.name}+1"
          redirect_to @user
      else
          render 'edit'
      end
  end




  def destroy
      User.find(params[:id]).destroy
      flash[:success] = "User deleted"
      redirect_to users_url
  end


  def following
      @title = "Following"
      @user = User.find(params[:id])
      @users = @user.following.paginate(page:params[:page])
      render 'show_follow'
  end

  def followers
      @title = "Following"
      @user = User.find(params[:id])
      @users = @user.followers.paginate(page:params[:page])
      render 'show_follow'
  end

private
    def user_params
        params.require(:user).permit(:name, :email, :password,
                        :password_confirmation)
    end



    def admin_user
        redirect_to(root_url) unless current_user.admin?
    end

    def correct_user
        @user = User.find(params[:id])
        unless current_user?(@user)
            flash[:danger] = "You are not the one"
            redirect_to @user
        end
    end
end