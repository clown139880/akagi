class SessionsController < ApplicationController

  def new
  end

  def create
      @user = User.find_by(email: session_params[:email].downcase)
      if @user && @user.authenticate(session_params[:password])
          if @user.activated?
              log_in @user
              session_params[:remember_me]== "1" ? remember(@user) : forget(@user)
              redirect_back_or @user
          else
              message = "账户还没有召唤成功"
              message += "请检查你邮箱中的召唤阵"
              flash[:warning] = message
              redirect_to root_url
          end
      else
          flash.now[:danger] = "错误的召唤阵"
          render 'new'
      end
  end



  def destroy
      log_out if logged_in?
      redirect_to root_url
  end

  private
    def session_params
        params.require(:session).permit(:email, :password, :remember_me)
    end
end
