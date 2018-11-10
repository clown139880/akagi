class Api::V1::BaseController < ActionController::API
  # disable the CSRF token
  # disable cookies (no set-cookies header in response)
  before_action :destroy_session
  rescue_from ApiExceptions::BaseException, 
    :with => :api_error

  def destroy_session
    request.session_options[:skip] = true
  end

  private
  def api_error(status: 0, error: 10000, message: "something goes wrong")
    render json:{
      status: 0,
      error: error,
      message: message,
    }
  end

  def api_success(response: response, message: "success")
    render json:{
      status: 1,
      message: message,
      response: response,
    }
  end

end
