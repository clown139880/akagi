$("#events-body").html("<%= j (render partial: 'event', collection: @events, as: :event )%>");
