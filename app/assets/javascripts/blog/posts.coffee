# Place all the behaviors and hooks related to the matching controller here.
# All this logic will automatically be available in application.js.
# You can use CoffeeScript in this file: http://coffeescript.org/
$ ->
  $('.scroll-down').on 'click', (e) ->
    $cover = $(this).closest('.cover')
    $('html, body').animate { scrollTop: $cover.position().top + $cover.height() }, 800
    e.preventDefault()
    return

  $('.top-link').on 'click', (e) ->
    $('html, body').animate 'scrollTop': 0
    e.preventDefault()
    return