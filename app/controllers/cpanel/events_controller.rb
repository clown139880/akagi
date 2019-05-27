class Cpanel::EventsController < Cpanel::BaseController

	before_action :set_event, only: [:edit, :update, :destroy]

	def index
    @events = Event.paginate(page: params[:page])
  end

	def new
		@event = Event.new
	end

	def create
		@event = Event.new event_params
		if @event.save
			if params[:event][:logo]
				@logo = Photo.create ({
					photoable: @event,
					url: params[:event][:logo],
					is_logo: true
				})
			end
			redirect_to [:cpanel, :events], notice:"创建成功"
		else
			render :new
		end
	end

	def edit
	end

	def update
		@event.update_attributes event_params
		if params[:event][:logo] != @event.get_logo
			if @event.photos.logo
				@event.photos.logo.update_all :is_logo => false
			end
			Photo.create ({
				photoable: @event,
				url: params[:event][:logo],
				is_logo: true
			})
		end
		redirect_to [:cpanel, :events], notice:"更新成功"
	end

	def destroy
		@event.delete
		redirect_to [:cpanel, :events], notice:"删除成功"
	end

	def get_options
    theevent = Event.find_by(id: params[:id])
		if theevent
		  options = [[theevent.id, theevent.title]]
		  theevent.subevents.each do |s|
		    options << [s.id, s.title]
		  end
			options << [0, '重置']
		  render :json => options
		else
			options = []
			Event.level(1).each do |s|
		    options << [s.id, s.title]
		  end
			render :json => options
		end
  end


	private
	def event_params
		if params[:post]
			params[:event] = params[:event].merge params[:post]
		end
		params.require(:event).permit(:title, :content, :place, :started_at, :ended_at, :tag_list)
	end

	def set_event
		@event = Event.find_by(id: params[:id])
	end

end
