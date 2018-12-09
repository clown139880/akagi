$("#events-body").html("<%= j (render partial: 'event', collection: @events, as: :event )%>");
$("#flash-notice").html("<%= j (render 'cpanel/public/flash_notice') %>");
