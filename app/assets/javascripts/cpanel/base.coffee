# Place all the behaviors and hooks related to the matching controller here.
# All this logic will automatically be available in application.js.
# You can use CoffeeScript in this file: http://coffeescript.org/
$ ->
    if $('#editormd').length > 0
        if window.screen.width > 767 
            window.Editor = editormd('editormd',
                width   : "80%",
                height  : "800px",
                syncScrolling : true,
                path: "    /editormdlib/")
        else 
                window.Editor = editormd('editormd',
                width   : "100%",
                height  : "800px",
                watch   : false,
                syncScrolling : true,
                toolbar : false,
                path: "    /editormdlib/")