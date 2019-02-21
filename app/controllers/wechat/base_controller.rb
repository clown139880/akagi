class Wechat::BaseController < ApplicationController

  layout 'wechat'

  #before_action :wechat_oauth

  protected

  def wechat_oauth

    # session[:weixin_openid] = "oO0Ntt2Y7IKOPqOPkZgwh7ep5IUI" if Rails.env.development?
    session[:weixin_openid] = Settings.test_openid if Rails.env.development?

    if session[:weixin_openid].blank? && !request.url.include?('notify')
      if params[:code].nil? || params[:code].empty?
        redirect_to $wechat_client.authorize_url(request.url, 'snsapi_userinfo')
      end

      begin
        oauth = $wechat_client.get_oauth_access_token(params[:code])
        session[:weixin_openid] = oauth.result[:openid]
        session[:access_token] = oauth.result[:access_token]
      rescue Exception => e
        Rails.logger.info("wechat oauth exception. #{e}")
      end

    end
  end

end
