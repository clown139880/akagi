$ ->
  $photos = 0
  $('#case-file').change ->

    data = new FormData()
    data.append 'photo', $('#case-file')[0].files[0]

    $.ajax
      url: '/wechat/photos'
      type: 'POST',
      data: data,
      processData: false,
      contentType: false,
      success: (data) ->
        if data.success == true
          img = data.file_path

          html = '<li class="weui-uploader__file" style="background-image:url(' + img + ')"></li>'
          $('.img-preview').append html

          input = document.createElement('input')
          input.name = 'case[photos_attributes][' + $photos + '][image]'
          input.value = img
          input.type = 'hidden'
          $('.img-preview').append input
          $photos = $photos + 1

    return

  $('#post-file').change ->

    data = new FormData()
    data.append 'photo', $('#post-file')[0].files[0]

    $.ajax
      url: '/wechat/photos'
      type: 'POST',
      data: data,
      processData: false,
      contentType: false,
      success: (data) ->
        if data.success == true
          img = data.file_path

          html = '<li class="weui-uploader__file" style="background-image:url(' + img + ')"></li>'
          $('.img-preview').append html

          input = document.createElement('input')
          input.name = 'post[photos_attributes][' + $photos + '][image]'
          input.value = img
          input.type = 'hidden'
          $('.img-preview').append input
          $photos = $photos + 1

    return