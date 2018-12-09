class Blog::HomeController < Blog::BaseController

  def index
    @events = Event.paginate(page: params[:page], per_page: 10)
  end

  def about
    @content = fileread('md/loveletter.md')
  end

  def contact
  end

end
