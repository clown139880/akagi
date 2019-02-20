$ ->
  $photos = 0

  showphoto =(photo)->
    $('#image' + $photos).removeClass 'weui-uploader__file_status'
    $('#image' + $photos).attr 'style', 'background-image:url(' + photo + '!200x200)'
    input = document.createElement('input')
    input.name = 'uploads[photos_attributes][' + $photos + '][url]'
    input.value = photo
    input.type = 'hidden'
    $('.img-preview').append input
    $photos = $photos + 1

  $('#qiniuimage').change ->
    f = $('#qiniuimage')[0].files[0]
    key = $('#key').val() + Date.parse(new Date()) + '.jpg'
    data = new FormData()
    data.append 'key', key
    data.append 'Signature', $('#Signature').val()
    data.append 'policy', $('#policy').val()
    data.append 'OSSAccessKeyId', $('#OSSAccessKeyId').val()
    data.append 'file', f
    data.append 'success_action_status', 200
    index = $photos
    li = $('<li class="weui-uploader__file weui-uploader__file_status" style="background-image:url(' + f + ')"><div class="weui-uploader__file-content"><i class="weui-loading"></i></div></li>')
    li.attr 'id', 'image' + index
    $('.img-preview').append li

    xhr = new XMLHttpRequest();
    xhr.open('POST', $('#bucket_endpoint').val(), true);
    xhr.send(data);
    xhr.onreadystatechange = (response)->
      if (xhr.readyState == 4 && xhr.status == 200)
        showphoto $('#bucket_endpoint').val() + key
      else if (xhr.status != 200 && xhr.responseText)
        console.log xhr.responseText
        wechat_notice('上传失败')

  wechat_notice = (notice)->
    $('p.weui-toast-content').html notice
    $('.wechat-notice').show()
    setTimeout  ->
      $('.wechat-notice').hide()
    , 3000