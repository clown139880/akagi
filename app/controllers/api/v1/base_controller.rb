class Api::V1::BaseController < ActionController::API
  # disable the CSRF token
  # disable cookies (no set-cookies header in response)
  before_action :destroy_session
  rescue_from ApiExceptions::BaseException, 
    :with => :api_error

  def destroy_session
    request.session_options[:skip] = true
  end

end
