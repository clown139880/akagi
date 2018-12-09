$(function() {
  $('#event_started_at').datetimepicker({
    format: 'YYYY-MM-DD'
  });
  $('#event_ended_at').datetimepicker({
    format: 'YYYY-MM-DD'
  });
  $("#event_started_at").on("dp.change", function (e) {
      $('#event_ended_at').data("DateTimePicker").minDate(e.date);
  });
  $("#event_ended_at").on("dp.change", function (e) {
      $('#event_started_at').data("DateTimePicker").maxDate(e.date);
  });
});
