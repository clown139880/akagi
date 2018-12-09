Rails.application.routes.draw do

  root 'blog/home#index'
  get '/login/authorized', to:'blog/sessions#authorize'
  get 'sitemap', to:'blog/sitemaps#index'

  namespace :api do
    namespace :v1 do
      resources :posts do
        collection do
          post :add_note
          get :get_note
          get :check_note
        end
      end
      post '/dorothy/receive', to:'dorothy#receive'
    end
  end

  namespace :blog do

    root 'home#index'
    get '/bar', to:'home#bar'
    get '/login', to:'sessions#new',     as: :login
    post '/login', to:'sessions#create'

    resources :posts
    resources :users
    resources :events
  end

  namespace :cpanel do

    root 'events#index'

    resources :posts
    resources :users
    resources :events do
      collection do
        post :get_options
      end
    end

  end

  namespace :wechat do

    root 'home#index'

    resources :posts do
      collection do
        get :more
      end
    end
    resources :users
    resources :cases
    resources :photos do
      collection do
        post :gettoken
      end
    end


  end

end
