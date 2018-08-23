/*
 * FitVids 1.1
 * Copyright 2013, Chris Coyier - http://css-tricks.com + Dave Rupert - http://daverupert.com
 * Credit to Thierry Koblentz - http://www.alistapart.com/articles/creating-intrinsic-ratios-for-video/
 * Released under the WTFPL license - http://sam.zoy.org/wtfpl/
 */
! function(t) {
  "use strict";
  t.fn.fitVids = function(e) {
    var i = {
      customSelector: null
    };
    if (!document.getElementById("fit-vids-style")) {
      var r = document.head || document.getElementsByTagName("head")[0],
        d = ".fluid-width-video-wrapper{width:100%;position:relative;padding:0;}.fluid-width-video-wrapper iframe,.fluid-width-video-wrapper object,.fluid-width-video-wrapper embed {position:absolute;top:0;left:0;width:100%;height:100%;}",
        a = document.createElement("div");
      a.innerHTML = '<p>x</p><style id="fit-vids-style">' + d + "</style>", r.appendChild(a.childNodes[1])
    }
    return e && t.extend(i, e), this.each(function() {
      var e = ["iframe[src*='player.vimeo.com']", "iframe[src*='youtube.com']", "iframe[src*='youtube-nocookie.com']", "iframe[src*='kickstarter.com'][src*='video.html']", "object", "embed"];
      i.customSelector && e.push(i.customSelector);
      var r = t(this).find(e.join(","));
      r = r.not("object object"), r.each(function() {
        var e = t(this);
        if (!("embed" === this.tagName.toLowerCase() && e.parent("object").length || e.parent(".fluid-width-video-wrapper").length)) {
          var i = "object" === this.tagName.toLowerCase() || e.attr("height") && !isNaN(parseInt(e.attr("height"), 10)) ? parseInt(e.attr("height"), 10) : e.height(),
            r = isNaN(parseInt(e.attr("width"), 10)) ? e.width() : parseInt(e.attr("width"), 10),
            d = i / r;
          if (!e.attr("id")) {
            var a = "fitvid" + Math.floor(999999 * Math.random());
            e.attr("id", a)
          }
          e.wrap('<div class="fluid-width-video-wrapper"></div>').parent(".fluid-width-video-wrapper").css("padding-top", 100 * d + "%"), e.removeAttr("height").removeAttr("width")
        }
      })
    })
  }
}(window.jQuery || window.Zepto);

/*
 * Reading Time
 * Author: Michael Lynch
 * Author URL: http://michaelynch.com
 * Date Created: August 14, 2013
 * Date Updated: July 11, 2016
 * Licensed under the MIT license
 */
! function(e) {
  var n;
  e.fn.readingTime = function(i) {
    var t = {
        readingTimeTarget: ".eta",
        readingTimeAsNumber: !1,
        wordCountTarget: null,
        wordsPerMinute: 270,
        round: !0,
        lang: "en",
        lessThanAMinuteString: "",
        prependTimeString: "",
        prependWordString: "",
        remotePath: null,
        remoteTarget: null,
        success: function() {},
        error: function() {}
      },
      r = this,
      a = e(this);
    r.settings = e.extend({}, t, i);
    var s = r.settings;
    if (!this.length) return s.error.call(this), this;
    if ("it" == s.lang) var l = s.lessThanAMinuteString || "Meno di un minuto",
      u = "min";
    else if ("fr" == s.lang) var l = s.lessThanAMinuteString || "Moins d'une minute",
      u = "min";
    else if ("de" == s.lang) var l = s.lessThanAMinuteString || "Weniger als eine Minute",
      u = "min";
    else if ("es" == s.lang) var l = s.lessThanAMinuteString || "Menos de un minuto",
      u = "min";
    else if ("nl" == s.lang) var l = s.lessThanAMinuteString || "Minder dan een minuut",
      u = "min";
    else if ("sk" == s.lang) var l = s.lessThanAMinuteString || "Menej než minútu",
      u = "min";
    else if ("cz" == s.lang) var l = s.lessThanAMinuteString || "Méně než minutu",
      u = "min";
    else if ("hu" == s.lang) var l = s.lessThanAMinuteString || "Kevesebb mint egy perc",
      u = "perc";
    else if ("ru" == s.lang) var l = s.lessThanAMinuteString || "Меньше минуты",
      u = "мин";
    else if ("ar" == s.lang) var l = s.lessThanAMinuteString || "أقل من دقيقة",
      u = "دقيقة";
    else if ("da" == s.lang) var l = s.lessThanAMinuteString || "Mindre end et minut",
      u = "min";
    else if ("is" == s.lang) var l = s.lessThanAMinuteString || "Minna en eina mínútu",
      u = "min";
    else if ("no" == s.lang) var l = s.lessThanAMinuteString || "Mindre enn ett minutt",
      u = "min";
    else if ("pl" == s.lang) var l = s.lessThanAMinuteString || "Mniej niż minutę",
      u = "min";
    else if ("ru" == s.lang) var l = s.lessThanAMinuteString || "Меньше минуты",
      u = "мой";
    else if ("sv" == s.lang) var l = s.lessThanAMinuteString || "Mindre än en minut",
      u = "min";
    else if ("tr" == s.lang) var l = s.lessThanAMinuteString || "Bir dakikadan az",
      u = "dk";
    else var l = s.lessThanAMinuteString || "Less than a minute",
      u = "min";
    var g = function(i) {
      if ("" !== i) {
        var t = i.trim().split(/\s+/g).length,
          r = s.wordsPerMinute / 60;
        if (n = t / r, !0 === s.round) a = Math.round(n / 60);
        else var a = Math.floor(n / 60);
        var g = Math.round(n - 60 * a);
        if (!0 === s.round) a > 0 ? e(s.readingTimeTarget).text(s.prependTimeString + a + (s.readingTimeAsNumber ? "" : " " + u)) : e(s.readingTimeTarget).text(s.readingTimeAsNumber ? a : s.prependTimeString + l);
        else {
          var m = a + ":" + g;
          e(s.readingTimeTarget).text(s.prependTimeString + m)
        }
        "" !== s.wordCountTarget && void 0 !== s.wordCountTarget && e(s.wordCountTarget).text(s.prependWordString + t), s.success.call(this)
      } else s.error.call(this, "The element is empty.")
    };
    return a.each(function() {
      null != s.remotePath && null != s.remoteTarget ? e.get(s.remotePath, function(n) {
        g(e("<div>").html(n).find(s.remoteTarget).text())
      }) : g(a.text())
    }), n
  }
}(jQuery);

/**
 * Slick: 1.8.0
 * Author: Ken Wheeler
 * Website: http://kenwheeler.github.io
 * Docs: http://kenwheeler.github.io/slick
 * Repo: http://github.com/kenwheeler/slick
 * Issues: http://github.com/kenwheeler/slick/issues
 */
! function(i) {
  "use strict";
  "function" == typeof define && define.amd ? define(["jquery"], i) : "undefined" != typeof exports ? module.exports = i(require("jquery")) : i(jQuery)
}(function(i) {
  "use strict";
  var e = window.Slick || {};
  (e = function() {
    var e = 0;
    return function(t, o) {
      var s, n = this;
      n.defaults = {
        accessibility: !0,
        adaptiveHeight: !1,
        appendArrows: i(t),
        appendDots: i(t),
        arrows: !0,
        asNavFor: null,
        prevArrow: '<button class="slick-prev" aria-label="Previous" type="button">Previous</button>',
        nextArrow: '<button class="slick-next" aria-label="Next" type="button">Next</button>',
        autoplay: !1,
        autoplaySpeed: 3e3,
        centerMode: !1,
        centerPadding: "50px",
        cssEase: "ease",
        customPaging: function(e, t) {
          return i('<button type="button" />').text(t + 1)
        },
        dots: !1,
        dotsClass: "slick-dots",
        draggable: !0,
        easing: "linear",
        edgeFriction: .35,
        fade: !1,
        focusOnSelect: !1,
        focusOnChange: !1,
        infinite: !0,
        initialSlide: 0,
        lazyLoad: "ondemand",
        mobileFirst: !1,
        pauseOnHover: !0,
        pauseOnFocus: !0,
        pauseOnDotsHover: !1,
        respondTo: "window",
        responsive: null,
        rows: 1,
        rtl: !1,
        slide: "",
        slidesPerRow: 1,
        slidesToShow: 1,
        slidesToScroll: 1,
        speed: 500,
        swipe: !0,
        swipeToSlide: !1,
        touchMove: !0,
        touchThreshold: 5,
        useCSS: !0,
        useTransform: !0,
        variableWidth: !1,
        vertical: !1,
        verticalSwiping: !1,
        waitForAnimate: !0,
        zIndex: 1e3
      }, n.initials = {
        animating: !1,
        dragging: !1,
        autoPlayTimer: null,
        currentDirection: 0,
        currentLeft: null,
        currentSlide: 0,
        direction: 1,
        $dots: null,
        listWidth: null,
        listHeight: null,
        loadIndex: 0,
        $nextArrow: null,
        $prevArrow: null,
        scrolling: !1,
        slideCount: null,
        slideWidth: null,
        $slideTrack: null,
        $slides: null,
        sliding: !1,
        slideOffset: 0,
        swipeLeft: null,
        swiping: !1,
        $list: null,
        touchObject: {},
        transformsEnabled: !1,
        unslicked: !1
      }, i.extend(n, n.initials), n.activeBreakpoint = null, n.animType = null, n.animProp = null, n.breakpoints = [], n.breakpointSettings = [], n.cssTransitions = !1, n.focussed = !1, n.interrupted = !1, n.hidden = "hidden", n.paused = !0, n.positionProp = null, n.respondTo = null, n.rowCount = 1, n.shouldClick = !0, n.$slider = i(t), n.$slidesCache = null, n.transformType = null, n.transitionType = null, n.visibilityChange = "visibilitychange", n.windowWidth = 0, n.windowTimer = null, s = i(t).data("slick") || {}, n.options = i.extend({}, n.defaults, o, s), n.currentSlide = n.options.initialSlide, n.originalSettings = n.options, void 0 !== document.mozHidden ? (n.hidden = "mozHidden", n.visibilityChange = "mozvisibilitychange") : void 0 !== document.webkitHidden && (n.hidden = "webkitHidden", n.visibilityChange = "webkitvisibilitychange"), n.autoPlay = i.proxy(n.autoPlay, n), n.autoPlayClear = i.proxy(n.autoPlayClear, n), n.autoPlayIterator = i.proxy(n.autoPlayIterator, n), n.changeSlide = i.proxy(n.changeSlide, n), n.clickHandler = i.proxy(n.clickHandler, n), n.selectHandler = i.proxy(n.selectHandler, n), n.setPosition = i.proxy(n.setPosition, n), n.swipeHandler = i.proxy(n.swipeHandler, n), n.dragHandler = i.proxy(n.dragHandler, n), n.keyHandler = i.proxy(n.keyHandler, n), n.instanceUid = e++, n.htmlExpr = /^(?:\s*(<[\w\W]+>)[^>]*)$/, n.registerBreakpoints(), n.init(!0)
    }
  }()).prototype.activateADA = function() {
    this.$slideTrack.find(".slick-active").attr({
      "aria-hidden": "false"
    }).find("a, input, button, select").attr({
      tabindex: "0"
    })
  }, e.prototype.addSlide = e.prototype.slickAdd = function(e, t, o) {
    var s = this;
    if ("boolean" == typeof t) o = t, t = null;
    else if (t < 0 || t >= s.slideCount) return !1;
    s.unload(), "number" == typeof t ? 0 === t && 0 === s.$slides.length ? i(e).appendTo(s.$slideTrack) : o ? i(e).insertBefore(s.$slides.eq(t)) : i(e).insertAfter(s.$slides.eq(t)) : !0 === o ? i(e).prependTo(s.$slideTrack) : i(e).appendTo(s.$slideTrack), s.$slides = s.$slideTrack.children(this.options.slide), s.$slideTrack.children(this.options.slide).detach(), s.$slideTrack.append(s.$slides), s.$slides.each(function(e, t) {
      i(t).attr("data-slick-index", e)
    }), s.$slidesCache = s.$slides, s.reinit()
  }, e.prototype.animateHeight = function() {
    var i = this;
    if (1 === i.options.slidesToShow && !0 === i.options.adaptiveHeight && !1 === i.options.vertical) {
      var e = i.$slides.eq(i.currentSlide).outerHeight(!0);
      i.$list.animate({
        height: e
      }, i.options.speed)
    }
  }, e.prototype.animateSlide = function(e, t) {
    var o = {},
      s = this;
    s.animateHeight(), !0 === s.options.rtl && !1 === s.options.vertical && (e = -e), !1 === s.transformsEnabled ? !1 === s.options.vertical ? s.$slideTrack.animate({
      left: e
    }, s.options.speed, s.options.easing, t) : s.$slideTrack.animate({
      top: e
    }, s.options.speed, s.options.easing, t) : !1 === s.cssTransitions ? (!0 === s.options.rtl && (s.currentLeft = -s.currentLeft), i({
      animStart: s.currentLeft
    }).animate({
      animStart: e
    }, {
      duration: s.options.speed,
      easing: s.options.easing,
      step: function(i) {
        i = Math.ceil(i), !1 === s.options.vertical ? (o[s.animType] = "translate(" + i + "px, 0px)", s.$slideTrack.css(o)) : (o[s.animType] = "translate(0px," + i + "px)", s.$slideTrack.css(o))
      },
      complete: function() {
        t && t.call()
      }
    })) : (s.applyTransition(), e = Math.ceil(e), !1 === s.options.vertical ? o[s.animType] = "translate3d(" + e + "px, 0px, 0px)" : o[s.animType] = "translate3d(0px," + e + "px, 0px)", s.$slideTrack.css(o), t && setTimeout(function() {
      s.disableTransition(), t.call()
    }, s.options.speed))
  }, e.prototype.getNavTarget = function() {
    var e = this,
      t = e.options.asNavFor;
    return t && null !== t && (t = i(t).not(e.$slider)), t
  }, e.prototype.asNavFor = function(e) {
    var t = this.getNavTarget();
    null !== t && "object" == typeof t && t.each(function() {
      var t = i(this).slick("getSlick");
      t.unslicked || t.slideHandler(e, !0)
    })
  }, e.prototype.applyTransition = function(i) {
    var e = this,
      t = {};
    !1 === e.options.fade ? t[e.transitionType] = e.transformType + " " + e.options.speed + "ms " + e.options.cssEase : t[e.transitionType] = "opacity " + e.options.speed + "ms " + e.options.cssEase, !1 === e.options.fade ? e.$slideTrack.css(t) : e.$slides.eq(i).css(t)
  }, e.prototype.autoPlay = function() {
    var i = this;
    i.autoPlayClear(), i.slideCount > i.options.slidesToShow && (i.autoPlayTimer = setInterval(i.autoPlayIterator, i.options.autoplaySpeed))
  }, e.prototype.autoPlayClear = function() {
    var i = this;
    i.autoPlayTimer && clearInterval(i.autoPlayTimer)
  }, e.prototype.autoPlayIterator = function() {
    var i = this,
      e = i.currentSlide + i.options.slidesToScroll;
    i.paused || i.interrupted || i.focussed || (!1 === i.options.infinite && (1 === i.direction && i.currentSlide + 1 === i.slideCount - 1 ? i.direction = 0 : 0 === i.direction && (e = i.currentSlide - i.options.slidesToScroll, i.currentSlide - 1 == 0 && (i.direction = 1))), i.slideHandler(e))
  }, e.prototype.buildArrows = function() {
    var e = this;
    !0 === e.options.arrows && (e.$prevArrow = i(e.options.prevArrow).addClass("slick-arrow"), e.$nextArrow = i(e.options.nextArrow).addClass("slick-arrow"), e.slideCount > e.options.slidesToShow ? (e.$prevArrow.removeClass("slick-hidden").removeAttr("aria-hidden tabindex"), e.$nextArrow.removeClass("slick-hidden").removeAttr("aria-hidden tabindex"), e.htmlExpr.test(e.options.prevArrow) && e.$prevArrow.prependTo(e.options.appendArrows), e.htmlExpr.test(e.options.nextArrow) && e.$nextArrow.appendTo(e.options.appendArrows), !0 !== e.options.infinite && e.$prevArrow.addClass("slick-disabled").attr("aria-disabled", "true")) : e.$prevArrow.add(e.$nextArrow).addClass("slick-hidden").attr({
      "aria-disabled": "true",
      tabindex: "-1"
    }))
  }, e.prototype.buildDots = function() {
    var e, t, o = this;
    if (!0 === o.options.dots) {
      for (o.$slider.addClass("slick-dotted"), t = i("<ul />").addClass(o.options.dotsClass), e = 0; e <= o.getDotCount(); e += 1) t.append(i("<li />").append(o.options.customPaging.call(this, o, e)));
      o.$dots = t.appendTo(o.options.appendDots), o.$dots.find("li").first().addClass("slick-active")
    }
  }, e.prototype.buildOut = function() {
    var e = this;
    e.$slides = e.$slider.children(e.options.slide + ":not(.slick-cloned)").addClass("slick-slide"), e.slideCount = e.$slides.length, e.$slides.each(function(e, t) {
      i(t).attr("data-slick-index", e).data("originalStyling", i(t).attr("style") || "")
    }), e.$slider.addClass("slick-slider"), e.$slideTrack = 0 === e.slideCount ? i('<div class="slick-track"/>').appendTo(e.$slider) : e.$slides.wrapAll('<div class="slick-track"/>').parent(), e.$list = e.$slideTrack.wrap('<div class="slick-list"/>').parent(), e.$slideTrack.css("opacity", 0), !0 !== e.options.centerMode && !0 !== e.options.swipeToSlide || (e.options.slidesToScroll = 1), i("img[data-lazy]", e.$slider).not("[src]").addClass("slick-loading"), e.setupInfinite(), e.buildArrows(), e.buildDots(), e.updateDots(), e.setSlideClasses("number" == typeof e.currentSlide ? e.currentSlide : 0), !0 === e.options.draggable && e.$list.addClass("draggable")
  }, e.prototype.buildRows = function() {
    var i, e, t, o, s, n, r, l = this;
    if (o = document.createDocumentFragment(), n = l.$slider.children(), l.options.rows > 1) {
      for (r = l.options.slidesPerRow * l.options.rows, s = Math.ceil(n.length / r), i = 0; i < s; i++) {
        var d = document.createElement("div");
        for (e = 0; e < l.options.rows; e++) {
          var a = document.createElement("div");
          for (t = 0; t < l.options.slidesPerRow; t++) {
            var c = i * r + (e * l.options.slidesPerRow + t);
            n.get(c) && a.appendChild(n.get(c))
          }
          d.appendChild(a)
        }
        o.appendChild(d)
      }
      l.$slider.empty().append(o), l.$slider.children().children().children().css({
        width: 100 / l.options.slidesPerRow + "%",
        display: "inline-block"
      })
    }
  }, e.prototype.checkResponsive = function(e, t) {
    var o, s, n, r = this,
      l = !1,
      d = r.$slider.width(),
      a = window.innerWidth || i(window).width();
    if ("window" === r.respondTo ? n = a : "slider" === r.respondTo ? n = d : "min" === r.respondTo && (n = Math.min(a, d)), r.options.responsive && r.options.responsive.length && null !== r.options.responsive) {
      s = null;
      for (o in r.breakpoints) r.breakpoints.hasOwnProperty(o) && (!1 === r.originalSettings.mobileFirst ? n < r.breakpoints[o] && (s = r.breakpoints[o]) : n > r.breakpoints[o] && (s = r.breakpoints[o]));
      null !== s ? null !== r.activeBreakpoint ? (s !== r.activeBreakpoint || t) && (r.activeBreakpoint = s, "unslick" === r.breakpointSettings[s] ? r.unslick(s) : (r.options = i.extend({}, r.originalSettings, r.breakpointSettings[s]), !0 === e && (r.currentSlide = r.options.initialSlide), r.refresh(e)), l = s) : (r.activeBreakpoint = s, "unslick" === r.breakpointSettings[s] ? r.unslick(s) : (r.options = i.extend({}, r.originalSettings, r.breakpointSettings[s]), !0 === e && (r.currentSlide = r.options.initialSlide), r.refresh(e)), l = s) : null !== r.activeBreakpoint && (r.activeBreakpoint = null, r.options = r.originalSettings, !0 === e && (r.currentSlide = r.options.initialSlide), r.refresh(e), l = s), e || !1 === l || r.$slider.trigger("breakpoint", [r, l])
    }
  }, e.prototype.changeSlide = function(e, t) {
    var o, s, n, r = this,
      l = i(e.currentTarget);
    switch (l.is("a") && e.preventDefault(), l.is("li") || (l = l.closest("li")), n = r.slideCount % r.options.slidesToScroll != 0, o = n ? 0 : (r.slideCount - r.currentSlide) % r.options.slidesToScroll, e.data.message) {
      case "previous":
        s = 0 === o ? r.options.slidesToScroll : r.options.slidesToShow - o, r.slideCount > r.options.slidesToShow && r.slideHandler(r.currentSlide - s, !1, t);
        break;
      case "next":
        s = 0 === o ? r.options.slidesToScroll : o, r.slideCount > r.options.slidesToShow && r.slideHandler(r.currentSlide + s, !1, t);
        break;
      case "index":
        var d = 0 === e.data.index ? 0 : e.data.index || l.index() * r.options.slidesToScroll;
        r.slideHandler(r.checkNavigable(d), !1, t), l.children().trigger("focus");
        break;
      default:
        return
    }
  }, e.prototype.checkNavigable = function(i) {
    var e, t;
    if (e = this.getNavigableIndexes(), t = 0, i > e[e.length - 1]) i = e[e.length - 1];
    else
      for (var o in e) {
        if (i < e[o]) {
          i = t;
          break
        }
        t = e[o]
      }
    return i
  }, e.prototype.cleanUpEvents = function() {
    var e = this;
    e.options.dots && null !== e.$dots && (i("li", e.$dots).off("click.slick", e.changeSlide).off("mouseenter.slick", i.proxy(e.interrupt, e, !0)).off("mouseleave.slick", i.proxy(e.interrupt, e, !1)), !0 === e.options.accessibility && e.$dots.off("keydown.slick", e.keyHandler)), e.$slider.off("focus.slick blur.slick"), !0 === e.options.arrows && e.slideCount > e.options.slidesToShow && (e.$prevArrow && e.$prevArrow.off("click.slick", e.changeSlide), e.$nextArrow && e.$nextArrow.off("click.slick", e.changeSlide), !0 === e.options.accessibility && (e.$prevArrow && e.$prevArrow.off("keydown.slick", e.keyHandler), e.$nextArrow && e.$nextArrow.off("keydown.slick", e.keyHandler))), e.$list.off("touchstart.slick mousedown.slick", e.swipeHandler), e.$list.off("touchmove.slick mousemove.slick", e.swipeHandler), e.$list.off("touchend.slick mouseup.slick", e.swipeHandler), e.$list.off("touchcancel.slick mouseleave.slick", e.swipeHandler), e.$list.off("click.slick", e.clickHandler), i(document).off(e.visibilityChange, e.visibility), e.cleanUpSlideEvents(), !0 === e.options.accessibility && e.$list.off("keydown.slick", e.keyHandler), !0 === e.options.focusOnSelect && i(e.$slideTrack).children().off("click.slick", e.selectHandler), i(window).off("orientationchange.slick.slick-" + e.instanceUid, e.orientationChange), i(window).off("resize.slick.slick-" + e.instanceUid, e.resize), i("[draggable!=true]", e.$slideTrack).off("dragstart", e.preventDefault), i(window).off("load.slick.slick-" + e.instanceUid, e.setPosition)
  }, e.prototype.cleanUpSlideEvents = function() {
    var e = this;
    e.$list.off("mouseenter.slick", i.proxy(e.interrupt, e, !0)), e.$list.off("mouseleave.slick", i.proxy(e.interrupt, e, !1))
  }, e.prototype.cleanUpRows = function() {
    var i, e = this;
    e.options.rows > 1 && ((i = e.$slides.children().children()).removeAttr("style"), e.$slider.empty().append(i))
  }, e.prototype.clickHandler = function(i) {
    !1 === this.shouldClick && (i.stopImmediatePropagation(), i.stopPropagation(), i.preventDefault())
  }, e.prototype.destroy = function(e) {
    var t = this;
    t.autoPlayClear(), t.touchObject = {}, t.cleanUpEvents(), i(".slick-cloned", t.$slider).detach(), t.$dots && t.$dots.remove(), t.$prevArrow && t.$prevArrow.length && (t.$prevArrow.removeClass("slick-disabled slick-arrow slick-hidden").removeAttr("aria-hidden aria-disabled tabindex").css("display", ""), t.htmlExpr.test(t.options.prevArrow) && t.$prevArrow.remove()), t.$nextArrow && t.$nextArrow.length && (t.$nextArrow.removeClass("slick-disabled slick-arrow slick-hidden").removeAttr("aria-hidden aria-disabled tabindex").css("display", ""), t.htmlExpr.test(t.options.nextArrow) && t.$nextArrow.remove()), t.$slides && (t.$slides.removeClass("slick-slide slick-active slick-center slick-visible slick-current").removeAttr("aria-hidden").removeAttr("data-slick-index").each(function() {
      i(this).attr("style", i(this).data("originalStyling"))
    }), t.$slideTrack.children(this.options.slide).detach(), t.$slideTrack.detach(), t.$list.detach(), t.$slider.append(t.$slides)), t.cleanUpRows(), t.$slider.removeClass("slick-slider"), t.$slider.removeClass("slick-initialized"), t.$slider.removeClass("slick-dotted"), t.unslicked = !0, e || t.$slider.trigger("destroy", [t])
  }, e.prototype.disableTransition = function(i) {
    var e = this,
      t = {};
    t[e.transitionType] = "", !1 === e.options.fade ? e.$slideTrack.css(t) : e.$slides.eq(i).css(t)
  }, e.prototype.fadeSlide = function(i, e) {
    var t = this;
    !1 === t.cssTransitions ? (t.$slides.eq(i).css({
      zIndex: t.options.zIndex
    }), t.$slides.eq(i).animate({
      opacity: 1
    }, t.options.speed, t.options.easing, e)) : (t.applyTransition(i), t.$slides.eq(i).css({
      opacity: 1,
      zIndex: t.options.zIndex
    }), e && setTimeout(function() {
      t.disableTransition(i), e.call()
    }, t.options.speed))
  }, e.prototype.fadeSlideOut = function(i) {
    var e = this;
    !1 === e.cssTransitions ? e.$slides.eq(i).animate({
      opacity: 0,
      zIndex: e.options.zIndex - 2
    }, e.options.speed, e.options.easing) : (e.applyTransition(i), e.$slides.eq(i).css({
      opacity: 0,
      zIndex: e.options.zIndex - 2
    }))
  }, e.prototype.filterSlides = e.prototype.slickFilter = function(i) {
    var e = this;
    null !== i && (e.$slidesCache = e.$slides, e.unload(), e.$slideTrack.children(this.options.slide).detach(), e.$slidesCache.filter(i).appendTo(e.$slideTrack), e.reinit())
  }, e.prototype.focusHandler = function() {
    var e = this;
    e.$slider.off("focus.slick blur.slick").on("focus.slick blur.slick", "*", function(t) {
      t.stopImmediatePropagation();
      var o = i(this);
      setTimeout(function() {
        e.options.pauseOnFocus && (e.focussed = o.is(":focus"), e.autoPlay())
      }, 0)
    })
  }, e.prototype.getCurrent = e.prototype.slickCurrentSlide = function() {
    return this.currentSlide
  }, e.prototype.getDotCount = function() {
    var i = this,
      e = 0,
      t = 0,
      o = 0;
    if (!0 === i.options.infinite)
      if (i.slideCount <= i.options.slidesToShow) ++o;
      else
        for (; e < i.slideCount;) ++o, e = t + i.options.slidesToScroll, t += i.options.slidesToScroll <= i.options.slidesToShow ? i.options.slidesToScroll : i.options.slidesToShow;
    else if (!0 === i.options.centerMode) o = i.slideCount;
    else if (i.options.asNavFor)
      for (; e < i.slideCount;) ++o, e = t + i.options.slidesToScroll, t += i.options.slidesToScroll <= i.options.slidesToShow ? i.options.slidesToScroll : i.options.slidesToShow;
    else o = 1 + Math.ceil((i.slideCount - i.options.slidesToShow) / i.options.slidesToScroll);
    return o - 1
  }, e.prototype.getLeft = function(i) {
    var e, t, o, s, n = this,
      r = 0;
    return n.slideOffset = 0, t = n.$slides.first().outerHeight(!0), !0 === n.options.infinite ? (n.slideCount > n.options.slidesToShow && (n.slideOffset = n.slideWidth * n.options.slidesToShow * -1, s = -1, !0 === n.options.vertical && !0 === n.options.centerMode && (2 === n.options.slidesToShow ? s = -1.5 : 1 === n.options.slidesToShow && (s = -2)), r = t * n.options.slidesToShow * s), n.slideCount % n.options.slidesToScroll != 0 && i + n.options.slidesToScroll > n.slideCount && n.slideCount > n.options.slidesToShow && (i > n.slideCount ? (n.slideOffset = (n.options.slidesToShow - (i - n.slideCount)) * n.slideWidth * -1, r = (n.options.slidesToShow - (i - n.slideCount)) * t * -1) : (n.slideOffset = n.slideCount % n.options.slidesToScroll * n.slideWidth * -1, r = n.slideCount % n.options.slidesToScroll * t * -1))) : i + n.options.slidesToShow > n.slideCount && (n.slideOffset = (i + n.options.slidesToShow - n.slideCount) * n.slideWidth, r = (i + n.options.slidesToShow - n.slideCount) * t), n.slideCount <= n.options.slidesToShow && (n.slideOffset = 0, r = 0), !0 === n.options.centerMode && n.slideCount <= n.options.slidesToShow ? n.slideOffset = n.slideWidth * Math.floor(n.options.slidesToShow) / 2 - n.slideWidth * n.slideCount / 2 : !0 === n.options.centerMode && !0 === n.options.infinite ? n.slideOffset += n.slideWidth * Math.floor(n.options.slidesToShow / 2) - n.slideWidth : !0 === n.options.centerMode && (n.slideOffset = 0, n.slideOffset += n.slideWidth * Math.floor(n.options.slidesToShow / 2)), e = !1 === n.options.vertical ? i * n.slideWidth * -1 + n.slideOffset : i * t * -1 + r, !0 === n.options.variableWidth && (o = n.slideCount <= n.options.slidesToShow || !1 === n.options.infinite ? n.$slideTrack.children(".slick-slide").eq(i) : n.$slideTrack.children(".slick-slide").eq(i + n.options.slidesToShow), e = !0 === n.options.rtl ? o[0] ? -1 * (n.$slideTrack.width() - o[0].offsetLeft - o.width()) : 0 : o[0] ? -1 * o[0].offsetLeft : 0, !0 === n.options.centerMode && (o = n.slideCount <= n.options.slidesToShow || !1 === n.options.infinite ? n.$slideTrack.children(".slick-slide").eq(i) : n.$slideTrack.children(".slick-slide").eq(i + n.options.slidesToShow + 1), e = !0 === n.options.rtl ? o[0] ? -1 * (n.$slideTrack.width() - o[0].offsetLeft - o.width()) : 0 : o[0] ? -1 * o[0].offsetLeft : 0, e += (n.$list.width() - o.outerWidth()) / 2)), e
  }, e.prototype.getOption = e.prototype.slickGetOption = function(i) {
    return this.options[i]
  }, e.prototype.getNavigableIndexes = function() {
    var i, e = this,
      t = 0,
      o = 0,
      s = [];
    for (!1 === e.options.infinite ? i = e.slideCount : (t = -1 * e.options.slidesToScroll, o = -1 * e.options.slidesToScroll, i = 2 * e.slideCount); t < i;) s.push(t), t = o + e.options.slidesToScroll, o += e.options.slidesToScroll <= e.options.slidesToShow ? e.options.slidesToScroll : e.options.slidesToShow;
    return s
  }, e.prototype.getSlick = function() {
    return this
  }, e.prototype.getSlideCount = function() {
    var e, t, o = this;
    return t = !0 === o.options.centerMode ? o.slideWidth * Math.floor(o.options.slidesToShow / 2) : 0, !0 === o.options.swipeToSlide ? (o.$slideTrack.find(".slick-slide").each(function(s, n) {
      if (n.offsetLeft - t + i(n).outerWidth() / 2 > -1 * o.swipeLeft) return e = n, !1
    }), Math.abs(i(e).attr("data-slick-index") - o.currentSlide) || 1) : o.options.slidesToScroll
  }, e.prototype.goTo = e.prototype.slickGoTo = function(i, e) {
    this.changeSlide({
      data: {
        message: "index",
        index: parseInt(i)
      }
    }, e)
  }, e.prototype.init = function(e) {
    var t = this;
    i(t.$slider).hasClass("slick-initialized") || (i(t.$slider).addClass("slick-initialized"), t.buildRows(), t.buildOut(), t.setProps(), t.startLoad(), t.loadSlider(), t.initializeEvents(), t.updateArrows(), t.updateDots(), t.checkResponsive(!0), t.focusHandler()), e && t.$slider.trigger("init", [t]), !0 === t.options.accessibility && t.initADA(), t.options.autoplay && (t.paused = !1, t.autoPlay())
  }, e.prototype.initADA = function() {
    var e = this,
      t = Math.ceil(e.slideCount / e.options.slidesToShow),
      o = e.getNavigableIndexes().filter(function(i) {
        return i >= 0 && i < e.slideCount
      });
    e.$slides.add(e.$slideTrack.find(".slick-cloned")).attr({
      "aria-hidden": "true",
      tabindex: "-1"
    }).find("a, input, button, select").attr({
      tabindex: "-1"
    }), null !== e.$dots && (e.$slides.not(e.$slideTrack.find(".slick-cloned")).each(function(t) {
      var s = o.indexOf(t);
      i(this).attr({
        role: "tabpanel",
        id: "slick-slide" + e.instanceUid + t,
        tabindex: -1
      }), -1 !== s && i(this).attr({
        "aria-describedby": "slick-slide-control" + e.instanceUid + s
      })
    }), e.$dots.attr("role", "tablist").find("li").each(function(s) {
      var n = o[s];
      i(this).attr({
        role: "presentation"
      }), i(this).find("button").first().attr({
        role: "tab",
        id: "slick-slide-control" + e.instanceUid + s,
        "aria-controls": "slick-slide" + e.instanceUid + n,
        "aria-label": s + 1 + " of " + t,
        "aria-selected": null,
        tabindex: "-1"
      })
    }).eq(e.currentSlide).find("button").attr({
      "aria-selected": "true",
      tabindex: "0"
    }).end());
    for (var s = e.currentSlide, n = s + e.options.slidesToShow; s < n; s++) e.$slides.eq(s).attr("tabindex", 0);
    e.activateADA()
  }, e.prototype.initArrowEvents = function() {
    var i = this;
    !0 === i.options.arrows && i.slideCount > i.options.slidesToShow && (i.$prevArrow.off("click.slick").on("click.slick", {
      message: "previous"
    }, i.changeSlide), i.$nextArrow.off("click.slick").on("click.slick", {
      message: "next"
    }, i.changeSlide), !0 === i.options.accessibility && (i.$prevArrow.on("keydown.slick", i.keyHandler), i.$nextArrow.on("keydown.slick", i.keyHandler)))
  }, e.prototype.initDotEvents = function() {
    var e = this;
    !0 === e.options.dots && (i("li", e.$dots).on("click.slick", {
      message: "index"
    }, e.changeSlide), !0 === e.options.accessibility && e.$dots.on("keydown.slick", e.keyHandler)), !0 === e.options.dots && !0 === e.options.pauseOnDotsHover && i("li", e.$dots).on("mouseenter.slick", i.proxy(e.interrupt, e, !0)).on("mouseleave.slick", i.proxy(e.interrupt, e, !1))
  }, e.prototype.initSlideEvents = function() {
    var e = this;
    e.options.pauseOnHover && (e.$list.on("mouseenter.slick", i.proxy(e.interrupt, e, !0)), e.$list.on("mouseleave.slick", i.proxy(e.interrupt, e, !1)))
  }, e.prototype.initializeEvents = function() {
    var e = this;
    e.initArrowEvents(), e.initDotEvents(), e.initSlideEvents(), e.$list.on("touchstart.slick mousedown.slick", {
      action: "start"
    }, e.swipeHandler), e.$list.on("touchmove.slick mousemove.slick", {
      action: "move"
    }, e.swipeHandler), e.$list.on("touchend.slick mouseup.slick", {
      action: "end"
    }, e.swipeHandler), e.$list.on("touchcancel.slick mouseleave.slick", {
      action: "end"
    }, e.swipeHandler), e.$list.on("click.slick", e.clickHandler), i(document).on(e.visibilityChange, i.proxy(e.visibility, e)), !0 === e.options.accessibility && e.$list.on("keydown.slick", e.keyHandler), !0 === e.options.focusOnSelect && i(e.$slideTrack).children().on("click.slick", e.selectHandler), i(window).on("orientationchange.slick.slick-" + e.instanceUid, i.proxy(e.orientationChange, e)), i(window).on("resize.slick.slick-" + e.instanceUid, i.proxy(e.resize, e)), i("[draggable!=true]", e.$slideTrack).on("dragstart", e.preventDefault), i(window).on("load.slick.slick-" + e.instanceUid, e.setPosition), i(e.setPosition)
  }, e.prototype.initUI = function() {
    var i = this;
    !0 === i.options.arrows && i.slideCount > i.options.slidesToShow && (i.$prevArrow.show(), i.$nextArrow.show()), !0 === i.options.dots && i.slideCount > i.options.slidesToShow && i.$dots.show()
  }, e.prototype.keyHandler = function(i) {
    var e = this;
    i.target.tagName.match("TEXTAREA|INPUT|SELECT") || (37 === i.keyCode && !0 === e.options.accessibility ? e.changeSlide({
      data: {
        message: !0 === e.options.rtl ? "next" : "previous"
      }
    }) : 39 === i.keyCode && !0 === e.options.accessibility && e.changeSlide({
      data: {
        message: !0 === e.options.rtl ? "previous" : "next"
      }
    }))
  }, e.prototype.lazyLoad = function() {
    function e(e) {
      i("img[data-lazy]", e).each(function() {
        var e = i(this),
          t = i(this).attr("data-lazy"),
          o = i(this).attr("data-srcset"),
          s = i(this).attr("data-sizes") || n.$slider.attr("data-sizes"),
          r = document.createElement("img");
        r.onload = function() {
          e.animate({
            opacity: 0
          }, 100, function() {
            o && (e.attr("srcset", o), s && e.attr("sizes", s)), e.attr("src", t).animate({
              opacity: 1
            }, 200, function() {
              e.removeAttr("data-lazy data-srcset data-sizes").removeClass("slick-loading")
            }), n.$slider.trigger("lazyLoaded", [n, e, t])
          })
        }, r.onerror = function() {
          e.removeAttr("data-lazy").removeClass("slick-loading").addClass("slick-lazyload-error"), n.$slider.trigger("lazyLoadError", [n, e, t])
        }, r.src = t
      })
    }
    var t, o, s, n = this;
    if (!0 === n.options.centerMode ? !0 === n.options.infinite ? s = (o = n.currentSlide + (n.options.slidesToShow / 2 + 1)) + n.options.slidesToShow + 2 : (o = Math.max(0, n.currentSlide - (n.options.slidesToShow / 2 + 1)), s = n.options.slidesToShow / 2 + 1 + 2 + n.currentSlide) : (o = n.options.infinite ? n.options.slidesToShow + n.currentSlide : n.currentSlide, s = Math.ceil(o + n.options.slidesToShow), !0 === n.options.fade && (o > 0 && o--, s <= n.slideCount && s++)), t = n.$slider.find(".slick-slide").slice(o, s), "anticipated" === n.options.lazyLoad)
      for (var r = o - 1, l = s, d = n.$slider.find(".slick-slide"), a = 0; a < n.options.slidesToScroll; a++) r < 0 && (r = n.slideCount - 1), t = (t = t.add(d.eq(r))).add(d.eq(l)), r--, l++;
    e(t), n.slideCount <= n.options.slidesToShow ? e(n.$slider.find(".slick-slide")) : n.currentSlide >= n.slideCount - n.options.slidesToShow ? e(n.$slider.find(".slick-cloned").slice(0, n.options.slidesToShow)) : 0 === n.currentSlide && e(n.$slider.find(".slick-cloned").slice(-1 * n.options.slidesToShow))
  }, e.prototype.loadSlider = function() {
    var i = this;
    i.setPosition(), i.$slideTrack.css({
      opacity: 1
    }), i.$slider.removeClass("slick-loading"), i.initUI(), "progressive" === i.options.lazyLoad && i.progressiveLazyLoad()
  }, e.prototype.next = e.prototype.slickNext = function() {
    this.changeSlide({
      data: {
        message: "next"
      }
    })
  }, e.prototype.orientationChange = function() {
    var i = this;
    i.checkResponsive(), i.setPosition()
  }, e.prototype.pause = e.prototype.slickPause = function() {
    var i = this;
    i.autoPlayClear(), i.paused = !0
  }, e.prototype.play = e.prototype.slickPlay = function() {
    var i = this;
    i.autoPlay(), i.options.autoplay = !0, i.paused = !1, i.focussed = !1, i.interrupted = !1
  }, e.prototype.postSlide = function(e) {
    var t = this;
    t.unslicked || (t.$slider.trigger("afterChange", [t, e]), t.animating = !1, t.slideCount > t.options.slidesToShow && t.setPosition(), t.swipeLeft = null, t.options.autoplay && t.autoPlay(), !0 === t.options.accessibility && (t.initADA(), t.options.focusOnChange && i(t.$slides.get(t.currentSlide)).attr("tabindex", 0).focus()))
  }, e.prototype.prev = e.prototype.slickPrev = function() {
    this.changeSlide({
      data: {
        message: "previous"
      }
    })
  }, e.prototype.preventDefault = function(i) {
    i.preventDefault()
  }, e.prototype.progressiveLazyLoad = function(e) {
    e = e || 1;
    var t, o, s, n, r, l = this,
      d = i("img[data-lazy]", l.$slider);
    d.length ? (t = d.first(), o = t.attr("data-lazy"), s = t.attr("data-srcset"), n = t.attr("data-sizes") || l.$slider.attr("data-sizes"), (r = document.createElement("img")).onload = function() {
      s && (t.attr("srcset", s), n && t.attr("sizes", n)), t.attr("src", o).removeAttr("data-lazy data-srcset data-sizes").removeClass("slick-loading"), !0 === l.options.adaptiveHeight && l.setPosition(), l.$slider.trigger("lazyLoaded", [l, t, o]), l.progressiveLazyLoad()
    }, r.onerror = function() {
      e < 3 ? setTimeout(function() {
        l.progressiveLazyLoad(e + 1)
      }, 500) : (t.removeAttr("data-lazy").removeClass("slick-loading").addClass("slick-lazyload-error"), l.$slider.trigger("lazyLoadError", [l, t, o]), l.progressiveLazyLoad())
    }, r.src = o) : l.$slider.trigger("allImagesLoaded", [l])
  }, e.prototype.refresh = function(e) {
    var t, o, s = this;
    o = s.slideCount - s.options.slidesToShow, !s.options.infinite && s.currentSlide > o && (s.currentSlide = o), s.slideCount <= s.options.slidesToShow && (s.currentSlide = 0), t = s.currentSlide, s.destroy(!0), i.extend(s, s.initials, {
      currentSlide: t
    }), s.init(), e || s.changeSlide({
      data: {
        message: "index",
        index: t
      }
    }, !1)
  }, e.prototype.registerBreakpoints = function() {
    var e, t, o, s = this,
      n = s.options.responsive || null;
    if ("array" === i.type(n) && n.length) {
      s.respondTo = s.options.respondTo || "window";
      for (e in n)
        if (o = s.breakpoints.length - 1, n.hasOwnProperty(e)) {
          for (t = n[e].breakpoint; o >= 0;) s.breakpoints[o] && s.breakpoints[o] === t && s.breakpoints.splice(o, 1), o--;
          s.breakpoints.push(t), s.breakpointSettings[t] = n[e].settings
        }
      s.breakpoints.sort(function(i, e) {
        return s.options.mobileFirst ? i - e : e - i
      })
    }
  }, e.prototype.reinit = function() {
    var e = this;
    e.$slides = e.$slideTrack.children(e.options.slide).addClass("slick-slide"), e.slideCount = e.$slides.length, e.currentSlide >= e.slideCount && 0 !== e.currentSlide && (e.currentSlide = e.currentSlide - e.options.slidesToScroll), e.slideCount <= e.options.slidesToShow && (e.currentSlide = 0), e.registerBreakpoints(), e.setProps(), e.setupInfinite(), e.buildArrows(), e.updateArrows(), e.initArrowEvents(), e.buildDots(), e.updateDots(), e.initDotEvents(), e.cleanUpSlideEvents(), e.initSlideEvents(), e.checkResponsive(!1, !0), !0 === e.options.focusOnSelect && i(e.$slideTrack).children().on("click.slick", e.selectHandler), e.setSlideClasses("number" == typeof e.currentSlide ? e.currentSlide : 0), e.setPosition(), e.focusHandler(), e.paused = !e.options.autoplay, e.autoPlay(), e.$slider.trigger("reInit", [e])
  }, e.prototype.resize = function() {
    var e = this;
    i(window).width() !== e.windowWidth && (clearTimeout(e.windowDelay), e.windowDelay = window.setTimeout(function() {
      e.windowWidth = i(window).width(), e.checkResponsive(), e.unslicked || e.setPosition()
    }, 50))
  }, e.prototype.removeSlide = e.prototype.slickRemove = function(i, e, t) {
    var o = this;
    if (i = "boolean" == typeof i ? !0 === (e = i) ? 0 : o.slideCount - 1 : !0 === e ? --i : i, o.slideCount < 1 || i < 0 || i > o.slideCount - 1) return !1;
    o.unload(), !0 === t ? o.$slideTrack.children().remove() : o.$slideTrack.children(this.options.slide).eq(i).remove(), o.$slides = o.$slideTrack.children(this.options.slide), o.$slideTrack.children(this.options.slide).detach(), o.$slideTrack.append(o.$slides), o.$slidesCache = o.$slides, o.reinit()
  }, e.prototype.setCSS = function(i) {
    var e, t, o = this,
      s = {};
    !0 === o.options.rtl && (i = -i), e = "left" == o.positionProp ? Math.ceil(i) + "px" : "0px", t = "top" == o.positionProp ? Math.ceil(i) + "px" : "0px", s[o.positionProp] = i, !1 === o.transformsEnabled ? o.$slideTrack.css(s) : (s = {}, !1 === o.cssTransitions ? (s[o.animType] = "translate(" + e + ", " + t + ")", o.$slideTrack.css(s)) : (s[o.animType] = "translate3d(" + e + ", " + t + ", 0px)", o.$slideTrack.css(s)))
  }, e.prototype.setDimensions = function() {
    var i = this;
    !1 === i.options.vertical ? !0 === i.options.centerMode && i.$list.css({
      padding: "0px " + i.options.centerPadding
    }) : (i.$list.height(i.$slides.first().outerHeight(!0) * i.options.slidesToShow), !0 === i.options.centerMode && i.$list.css({
      padding: i.options.centerPadding + " 0px"
    })), i.listWidth = i.$list.width(), i.listHeight = i.$list.height(), !1 === i.options.vertical && !1 === i.options.variableWidth ? (i.slideWidth = Math.ceil(i.listWidth / i.options.slidesToShow), i.$slideTrack.width(Math.ceil(i.slideWidth * i.$slideTrack.children(".slick-slide").length))) : !0 === i.options.variableWidth ? i.$slideTrack.width(5e3 * i.slideCount) : (i.slideWidth = Math.ceil(i.listWidth), i.$slideTrack.height(Math.ceil(i.$slides.first().outerHeight(!0) * i.$slideTrack.children(".slick-slide").length)));
    var e = i.$slides.first().outerWidth(!0) - i.$slides.first().width();
    !1 === i.options.variableWidth && i.$slideTrack.children(".slick-slide").width(i.slideWidth - e)
  }, e.prototype.setFade = function() {
    var e, t = this;
    t.$slides.each(function(o, s) {
      e = t.slideWidth * o * -1, !0 === t.options.rtl ? i(s).css({
        position: "relative",
        right: e,
        top: 0,
        zIndex: t.options.zIndex - 2,
        opacity: 0
      }) : i(s).css({
        position: "relative",
        left: e,
        top: 0,
        zIndex: t.options.zIndex - 2,
        opacity: 0
      })
    }), t.$slides.eq(t.currentSlide).css({
      zIndex: t.options.zIndex - 1,
      opacity: 1
    })
  }, e.prototype.setHeight = function() {
    var i = this;
    if (1 === i.options.slidesToShow && !0 === i.options.adaptiveHeight && !1 === i.options.vertical) {
      var e = i.$slides.eq(i.currentSlide).outerHeight(!0);
      i.$list.css("height", e)
    }
  }, e.prototype.setOption = e.prototype.slickSetOption = function() {
    var e, t, o, s, n, r = this,
      l = !1;
    if ("object" === i.type(arguments[0]) ? (o = arguments[0], l = arguments[1], n = "multiple") : "string" === i.type(arguments[0]) && (o = arguments[0], s = arguments[1], l = arguments[2], "responsive" === arguments[0] && "array" === i.type(arguments[1]) ? n = "responsive" : void 0 !== arguments[1] && (n = "single")), "single" === n) r.options[o] = s;
    else if ("multiple" === n) i.each(o, function(i, e) {
      r.options[i] = e
    });
    else if ("responsive" === n)
      for (t in s)
        if ("array" !== i.type(r.options.responsive)) r.options.responsive = [s[t]];
        else {
          for (e = r.options.responsive.length - 1; e >= 0;) r.options.responsive[e].breakpoint === s[t].breakpoint && r.options.responsive.splice(e, 1), e--;
          r.options.responsive.push(s[t])
        }
    l && (r.unload(), r.reinit())
  }, e.prototype.setPosition = function() {
    var i = this;
    i.setDimensions(), i.setHeight(), !1 === i.options.fade ? i.setCSS(i.getLeft(i.currentSlide)) : i.setFade(), i.$slider.trigger("setPosition", [i])
  }, e.prototype.setProps = function() {
    var i = this,
      e = document.body.style;
    i.positionProp = !0 === i.options.vertical ? "top" : "left", "top" === i.positionProp ? i.$slider.addClass("slick-vertical") : i.$slider.removeClass("slick-vertical"), void 0 === e.WebkitTransition && void 0 === e.MozTransition && void 0 === e.msTransition || !0 === i.options.useCSS && (i.cssTransitions = !0), i.options.fade && ("number" == typeof i.options.zIndex ? i.options.zIndex < 3 && (i.options.zIndex = 3) : i.options.zIndex = i.defaults.zIndex), void 0 !== e.OTransform && (i.animType = "OTransform", i.transformType = "-o-transform", i.transitionType = "OTransition", void 0 === e.perspectiveProperty && void 0 === e.webkitPerspective && (i.animType = !1)), void 0 !== e.MozTransform && (i.animType = "MozTransform", i.transformType = "-moz-transform", i.transitionType = "MozTransition", void 0 === e.perspectiveProperty && void 0 === e.MozPerspective && (i.animType = !1)), void 0 !== e.webkitTransform && (i.animType = "webkitTransform", i.transformType = "-webkit-transform", i.transitionType = "webkitTransition", void 0 === e.perspectiveProperty && void 0 === e.webkitPerspective && (i.animType = !1)), void 0 !== e.msTransform && (i.animType = "msTransform", i.transformType = "-ms-transform", i.transitionType = "msTransition", void 0 === e.msTransform && (i.animType = !1)), void 0 !== e.transform && !1 !== i.animType && (i.animType = "transform", i.transformType = "transform", i.transitionType = "transition"), i.transformsEnabled = i.options.useTransform && null !== i.animType && !1 !== i.animType
  }, e.prototype.setSlideClasses = function(i) {
    var e, t, o, s, n = this;
    if (t = n.$slider.find(".slick-slide").removeClass("slick-active slick-center slick-current").attr("aria-hidden", "true"), n.$slides.eq(i).addClass("slick-current"), !0 === n.options.centerMode) {
      var r = n.options.slidesToShow % 2 == 0 ? 1 : 0;
      e = Math.floor(n.options.slidesToShow / 2), !0 === n.options.infinite && (i >= e && i <= n.slideCount - 1 - e ? n.$slides.slice(i - e + r, i + e + 1).addClass("slick-active").attr("aria-hidden", "false") : (o = n.options.slidesToShow + i, t.slice(o - e + 1 + r, o + e + 2).addClass("slick-active").attr("aria-hidden", "false")), 0 === i ? t.eq(t.length - 1 - n.options.slidesToShow).addClass("slick-center") : i === n.slideCount - 1 && t.eq(n.options.slidesToShow).addClass("slick-center")), n.$slides.eq(i).addClass("slick-center")
    } else i >= 0 && i <= n.slideCount - n.options.slidesToShow ? n.$slides.slice(i, i + n.options.slidesToShow).addClass("slick-active").attr("aria-hidden", "false") : t.length <= n.options.slidesToShow ? t.addClass("slick-active").attr("aria-hidden", "false") : (s = n.slideCount % n.options.slidesToShow, o = !0 === n.options.infinite ? n.options.slidesToShow + i : i, n.options.slidesToShow == n.options.slidesToScroll && n.slideCount - i < n.options.slidesToShow ? t.slice(o - (n.options.slidesToShow - s), o + s).addClass("slick-active").attr("aria-hidden", "false") : t.slice(o, o + n.options.slidesToShow).addClass("slick-active").attr("aria-hidden", "false"));
    "ondemand" !== n.options.lazyLoad && "anticipated" !== n.options.lazyLoad || n.lazyLoad()
  }, e.prototype.setupInfinite = function() {
    var e, t, o, s = this;
    if (!0 === s.options.fade && (s.options.centerMode = !1), !0 === s.options.infinite && !1 === s.options.fade && (t = null, s.slideCount > s.options.slidesToShow)) {
      for (o = !0 === s.options.centerMode ? s.options.slidesToShow + 1 : s.options.slidesToShow, e = s.slideCount; e > s.slideCount - o; e -= 1) t = e - 1, i(s.$slides[t]).clone(!0).attr("id", "").attr("data-slick-index", t - s.slideCount).prependTo(s.$slideTrack).addClass("slick-cloned");
      for (e = 0; e < o + s.slideCount; e += 1) t = e, i(s.$slides[t]).clone(!0).attr("id", "").attr("data-slick-index", t + s.slideCount).appendTo(s.$slideTrack).addClass("slick-cloned");
      s.$slideTrack.find(".slick-cloned").find("[id]").each(function() {
        i(this).attr("id", "")
      })
    }
  }, e.prototype.interrupt = function(i) {
    var e = this;
    i || e.autoPlay(), e.interrupted = i
  }, e.prototype.selectHandler = function(e) {
    var t = this,
      o = i(e.target).is(".slick-slide") ? i(e.target) : i(e.target).parents(".slick-slide"),
      s = parseInt(o.attr("data-slick-index"));
    s || (s = 0), t.slideCount <= t.options.slidesToShow ? t.slideHandler(s, !1, !0) : t.slideHandler(s)
  }, e.prototype.slideHandler = function(i, e, t) {
    var o, s, n, r, l, d = null,
      a = this;
    if (e = e || !1, !(!0 === a.animating && !0 === a.options.waitForAnimate || !0 === a.options.fade && a.currentSlide === i))
      if (!1 === e && a.asNavFor(i), o = i, d = a.getLeft(o), r = a.getLeft(a.currentSlide), a.currentLeft = null === a.swipeLeft ? r : a.swipeLeft, !1 === a.options.infinite && !1 === a.options.centerMode && (i < 0 || i > a.getDotCount() * a.options.slidesToScroll)) !1 === a.options.fade && (o = a.currentSlide, !0 !== t ? a.animateSlide(r, function() {
        a.postSlide(o)
      }) : a.postSlide(o));
      else if (!1 === a.options.infinite && !0 === a.options.centerMode && (i < 0 || i > a.slideCount - a.options.slidesToScroll)) !1 === a.options.fade && (o = a.currentSlide, !0 !== t ? a.animateSlide(r, function() {
      a.postSlide(o)
    }) : a.postSlide(o));
    else {
      if (a.options.autoplay && clearInterval(a.autoPlayTimer), s = o < 0 ? a.slideCount % a.options.slidesToScroll != 0 ? a.slideCount - a.slideCount % a.options.slidesToScroll : a.slideCount + o : o >= a.slideCount ? a.slideCount % a.options.slidesToScroll != 0 ? 0 : o - a.slideCount : o, a.animating = !0, a.$slider.trigger("beforeChange", [a, a.currentSlide, s]), n = a.currentSlide, a.currentSlide = s, a.setSlideClasses(a.currentSlide), a.options.asNavFor && (l = (l = a.getNavTarget()).slick("getSlick")).slideCount <= l.options.slidesToShow && l.setSlideClasses(a.currentSlide), a.updateDots(), a.updateArrows(), !0 === a.options.fade) return !0 !== t ? (a.fadeSlideOut(n), a.fadeSlide(s, function() {
        a.postSlide(s)
      })) : a.postSlide(s), void a.animateHeight();
      !0 !== t ? a.animateSlide(d, function() {
        a.postSlide(s)
      }) : a.postSlide(s)
    }
  }, e.prototype.startLoad = function() {
    var i = this;
    !0 === i.options.arrows && i.slideCount > i.options.slidesToShow && (i.$prevArrow.hide(), i.$nextArrow.hide()), !0 === i.options.dots && i.slideCount > i.options.slidesToShow && i.$dots.hide(), i.$slider.addClass("slick-loading")
  }, e.prototype.swipeDirection = function() {
    var i, e, t, o, s = this;
    return i = s.touchObject.startX - s.touchObject.curX, e = s.touchObject.startY - s.touchObject.curY, t = Math.atan2(e, i), (o = Math.round(180 * t / Math.PI)) < 0 && (o = 360 - Math.abs(o)), o <= 45 && o >= 0 ? !1 === s.options.rtl ? "left" : "right" : o <= 360 && o >= 315 ? !1 === s.options.rtl ? "left" : "right" : o >= 135 && o <= 225 ? !1 === s.options.rtl ? "right" : "left" : !0 === s.options.verticalSwiping ? o >= 35 && o <= 135 ? "down" : "up" : "vertical"
  }, e.prototype.swipeEnd = function(i) {
    var e, t, o = this;
    if (o.dragging = !1, o.swiping = !1, o.scrolling) return o.scrolling = !1, !1;
    if (o.interrupted = !1, o.shouldClick = !(o.touchObject.swipeLength > 10), void 0 === o.touchObject.curX) return !1;
    if (!0 === o.touchObject.edgeHit && o.$slider.trigger("edge", [o, o.swipeDirection()]), o.touchObject.swipeLength >= o.touchObject.minSwipe) {
      switch (t = o.swipeDirection()) {
        case "left":
        case "down":
          e = o.options.swipeToSlide ? o.checkNavigable(o.currentSlide + o.getSlideCount()) : o.currentSlide + o.getSlideCount(), o.currentDirection = 0;
          break;
        case "right":
        case "up":
          e = o.options.swipeToSlide ? o.checkNavigable(o.currentSlide - o.getSlideCount()) : o.currentSlide - o.getSlideCount(), o.currentDirection = 1
      }
      "vertical" != t && (o.slideHandler(e), o.touchObject = {}, o.$slider.trigger("swipe", [o, t]))
    } else o.touchObject.startX !== o.touchObject.curX && (o.slideHandler(o.currentSlide), o.touchObject = {})
  }, e.prototype.swipeHandler = function(i) {
    var e = this;
    if (!(!1 === e.options.swipe || "ontouchend" in document && !1 === e.options.swipe || !1 === e.options.draggable && -1 !== i.type.indexOf("mouse"))) switch (e.touchObject.fingerCount = i.originalEvent && void 0 !== i.originalEvent.touches ? i.originalEvent.touches.length : 1, e.touchObject.minSwipe = e.listWidth / e.options.touchThreshold, !0 === e.options.verticalSwiping && (e.touchObject.minSwipe = e.listHeight / e.options.touchThreshold), i.data.action) {
      case "start":
        e.swipeStart(i);
        break;
      case "move":
        e.swipeMove(i);
        break;
      case "end":
        e.swipeEnd(i)
    }
  }, e.prototype.swipeMove = function(i) {
    var e, t, o, s, n, r, l = this;
    return n = void 0 !== i.originalEvent ? i.originalEvent.touches : null, !(!l.dragging || l.scrolling || n && 1 !== n.length) && (e = l.getLeft(l.currentSlide), l.touchObject.curX = void 0 !== n ? n[0].pageX : i.clientX, l.touchObject.curY = void 0 !== n ? n[0].pageY : i.clientY, l.touchObject.swipeLength = Math.round(Math.sqrt(Math.pow(l.touchObject.curX - l.touchObject.startX, 2))), r = Math.round(Math.sqrt(Math.pow(l.touchObject.curY - l.touchObject.startY, 2))), !l.options.verticalSwiping && !l.swiping && r > 4 ? (l.scrolling = !0, !1) : (!0 === l.options.verticalSwiping && (l.touchObject.swipeLength = r), t = l.swipeDirection(), void 0 !== i.originalEvent && l.touchObject.swipeLength > 4 && (l.swiping = !0, i.preventDefault()), s = (!1 === l.options.rtl ? 1 : -1) * (l.touchObject.curX > l.touchObject.startX ? 1 : -1), !0 === l.options.verticalSwiping && (s = l.touchObject.curY > l.touchObject.startY ? 1 : -1), o = l.touchObject.swipeLength, l.touchObject.edgeHit = !1, !1 === l.options.infinite && (0 === l.currentSlide && "right" === t || l.currentSlide >= l.getDotCount() && "left" === t) && (o = l.touchObject.swipeLength * l.options.edgeFriction, l.touchObject.edgeHit = !0), !1 === l.options.vertical ? l.swipeLeft = e + o * s : l.swipeLeft = e + o * (l.$list.height() / l.listWidth) * s, !0 === l.options.verticalSwiping && (l.swipeLeft = e + o * s), !0 !== l.options.fade && !1 !== l.options.touchMove && (!0 === l.animating ? (l.swipeLeft = null, !1) : void l.setCSS(l.swipeLeft))))
  }, e.prototype.swipeStart = function(i) {
    var e, t = this;
    if (t.interrupted = !0, 1 !== t.touchObject.fingerCount || t.slideCount <= t.options.slidesToShow) return t.touchObject = {}, !1;
    void 0 !== i.originalEvent && void 0 !== i.originalEvent.touches && (e = i.originalEvent.touches[0]), t.touchObject.startX = t.touchObject.curX = void 0 !== e ? e.pageX : i.clientX, t.touchObject.startY = t.touchObject.curY = void 0 !== e ? e.pageY : i.clientY, t.dragging = !0
  }, e.prototype.unfilterSlides = e.prototype.slickUnfilter = function() {
    var i = this;
    null !== i.$slidesCache && (i.unload(), i.$slideTrack.children(this.options.slide).detach(), i.$slidesCache.appendTo(i.$slideTrack), i.reinit())
  }, e.prototype.unload = function() {
    var e = this;
    i(".slick-cloned", e.$slider).remove(), e.$dots && e.$dots.remove(), e.$prevArrow && e.htmlExpr.test(e.options.prevArrow) && e.$prevArrow.remove(), e.$nextArrow && e.htmlExpr.test(e.options.nextArrow) && e.$nextArrow.remove(), e.$slides.removeClass("slick-slide slick-active slick-visible slick-current").attr("aria-hidden", "true").css("width", "")
  }, e.prototype.unslick = function(i) {
    var e = this;
    e.$slider.trigger("unslick", [e, i]), e.destroy()
  }, e.prototype.updateArrows = function() {
    var i = this;
    Math.floor(i.options.slidesToShow / 2), !0 === i.options.arrows && i.slideCount > i.options.slidesToShow && !i.options.infinite && (i.$prevArrow.removeClass("slick-disabled").attr("aria-disabled", "false"), i.$nextArrow.removeClass("slick-disabled").attr("aria-disabled", "false"), 0 === i.currentSlide ? (i.$prevArrow.addClass("slick-disabled").attr("aria-disabled", "true"), i.$nextArrow.removeClass("slick-disabled").attr("aria-disabled", "false")) : i.currentSlide >= i.slideCount - i.options.slidesToShow && !1 === i.options.centerMode ? (i.$nextArrow.addClass("slick-disabled").attr("aria-disabled", "true"), i.$prevArrow.removeClass("slick-disabled").attr("aria-disabled", "false")) : i.currentSlide >= i.slideCount - 1 && !0 === i.options.centerMode && (i.$nextArrow.addClass("slick-disabled").attr("aria-disabled", "true"), i.$prevArrow.removeClass("slick-disabled").attr("aria-disabled", "false")))
  }, e.prototype.updateDots = function() {
    var i = this;
    null !== i.$dots && (i.$dots.find("li").removeClass("slick-active").end(), i.$dots.find("li").eq(Math.floor(i.currentSlide / i.options.slidesToScroll)).addClass("slick-active"))
  }, e.prototype.visibility = function() {
    var i = this;
    i.options.autoplay && (document[i.hidden] ? i.interrupted = !0 : i.interrupted = !1)
  }, i.fn.slick = function() {
    var i, t, o = this,
      s = arguments[0],
      n = Array.prototype.slice.call(arguments, 1),
      r = o.length;
    for (i = 0; i < r; i++)
      if ("object" == typeof s || void 0 === s ? o[i].slick = new e(o[i], s) : t = o[i].slick[s].apply(o[i].slick, n), void 0 !== t) return t;
    return o
  }
});

/**
 * ghostHunter - 0.4.0
 * Copyright (C) 2014 Jamal Neufeld (jamal@i11u.me)
 * MIT Licensed
 * @license
 */
! function(t) {
  ! function() {
    var t = function(e) {
      var i = new t.Index;
      return i.pipeline.add(t.trimmer, t.stopWordFilter, t.stemmer), e && e.call(i, i), i
    };
    t.version = "0.7.0", t.utils = {}, t.utils.warn = function(t) {
        return function(e) {
          t.console && console.warn && console.warn(e)
        }
      }(this), t.utils.asString = function(t) {
        return void 0 === t || null === t ? "" : t.toString()
      }, t.EventEmitter = function() {
        this.events = {}
      }, t.EventEmitter.prototype.addListener = function() {
        var t = Array.prototype.slice.call(arguments),
          e = t.pop(),
          i = t;
        if ("function" != typeof e) throw new TypeError("last argument must be a function");
        i.forEach(function(t) {
          this.hasHandler(t) || (this.events[t] = []), this.events[t].push(e)
        }, this)
      }, t.EventEmitter.prototype.removeListener = function(t, e) {
        if (this.hasHandler(t)) {
          var i = this.events[t].indexOf(e);
          this.events[t].splice(i, 1), this.events[t].length || delete this.events[t]
        }
      }, t.EventEmitter.prototype.emit = function(t) {
        if (this.hasHandler(t)) {
          var e = Array.prototype.slice.call(arguments, 1);
          this.events[t].forEach(function(t) {
            t.apply(void 0, e)
          })
        }
      }, t.EventEmitter.prototype.hasHandler = function(t) {
        return t in this.events
      }, t.tokenizer = function(e) {
        return arguments.length && null != e && void 0 != e ? Array.isArray(e) ? e.map(function(e) {
          return t.utils.asString(e).toLowerCase()
        }) : e.toString().trim().toLowerCase().split(t.tokenizer.seperator) : []
      }, t.tokenizer.seperator = /[\s\-]+/, t.tokenizer.load = function(t) {
        var e = this.registeredFunctions[t];
        if (!e) throw new Error("Cannot load un-registered function: " + t);
        return e
      }, t.tokenizer.label = "default", t.tokenizer.registeredFunctions = {
        default: t.tokenizer
      }, t.tokenizer.registerFunction = function(e, i) {
        i in this.registeredFunctions && t.utils.warn("Overwriting existing tokenizer: " + i), e.label = i, this.registeredFunctions[i] = e
      }, t.Pipeline = function() {
        this._stack = []
      }, t.Pipeline.registeredFunctions = {}, t.Pipeline.registerFunction = function(e, i) {
        i in this.registeredFunctions && t.utils.warn("Overwriting existing registered function: " + i), e.label = i, t.Pipeline.registeredFunctions[e.label] = e
      }, t.Pipeline.warnIfFunctionNotRegistered = function(e) {
        e.label && e.label in this.registeredFunctions || t.utils.warn("Function is not registered with pipeline. This may cause problems when serialising the index.\n", e)
      }, t.Pipeline.load = function(e) {
        var i = new t.Pipeline;
        return e.forEach(function(e) {
          var n = t.Pipeline.registeredFunctions[e];
          if (!n) throw new Error("Cannot load un-registered function: " + e);
          i.add(n)
        }), i
      }, t.Pipeline.prototype.add = function() {
        Array.prototype.slice.call(arguments).forEach(function(e) {
          t.Pipeline.warnIfFunctionNotRegistered(e), this._stack.push(e)
        }, this)
      }, t.Pipeline.prototype.after = function(e, i) {
        t.Pipeline.warnIfFunctionNotRegistered(i);
        var n = this._stack.indexOf(e);
        if (-1 == n) throw new Error("Cannot find existingFn");
        n += 1, this._stack.splice(n, 0, i)
      }, t.Pipeline.prototype.before = function(e, i) {
        t.Pipeline.warnIfFunctionNotRegistered(i);
        var n = this._stack.indexOf(e);
        if (-1 == n) throw new Error("Cannot find existingFn");
        this._stack.splice(n, 0, i)
      }, t.Pipeline.prototype.remove = function(t) {
        var e = this._stack.indexOf(t); - 1 != e && this._stack.splice(e, 1)
      }, t.Pipeline.prototype.run = function(t) {
        for (var e = [], i = t.length, n = this._stack.length, o = 0; i > o; o++) {
          for (var r = t[o], s = 0; n > s && void 0 !== (r = this._stack[s](r, o, t)) && "" !== r; s++);
          void 0 !== r && "" !== r && e.push(r)
        }
        return e
      }, t.Pipeline.prototype.reset = function() {
        this._stack = []
      }, t.Pipeline.prototype.toJSON = function() {
        return this._stack.map(function(e) {
          return t.Pipeline.warnIfFunctionNotRegistered(e), e.label
        })
      }, t.Vector = function() {
        this._magnitude = null, this.list = void 0, this.length = 0
      }, t.Vector.Node = function(t, e, i) {
        this.idx = t, this.val = e, this.next = i
      }, t.Vector.prototype.insert = function(e, i) {
        this._magnitude = void 0;
        var n = this.list;
        if (!n) return this.list = new t.Vector.Node(e, i, n), this.length++;
        if (e < n.idx) return this.list = new t.Vector.Node(e, i, n), this.length++;
        for (var o = n, r = n.next; void 0 != r;) {
          if (e < r.idx) return o.next = new t.Vector.Node(e, i, r), this.length++;
          o = r, r = r.next
        }
        return o.next = new t.Vector.Node(e, i, r), this.length++
      }, t.Vector.prototype.magnitude = function() {
        if (this._magnitude) return this._magnitude;
        for (var t, e = this.list, i = 0; e;) t = e.val, i += t * t, e = e.next;
        return this._magnitude = Math.sqrt(i)
      }, t.Vector.prototype.dot = function(t) {
        for (var e = this.list, i = t.list, n = 0; e && i;) e.idx < i.idx ? e = e.next : e.idx > i.idx ? i = i.next : (n += e.val * i.val, e = e.next, i = i.next);
        return n
      }, t.Vector.prototype.similarity = function(t) {
        return this.dot(t) / (this.magnitude() * t.magnitude())
      }, t.SortedSet = function() {
        this.length = 0, this.elements = []
      }, t.SortedSet.load = function(t) {
        var e = new this;
        return e.elements = t, e.length = t.length, e
      }, t.SortedSet.prototype.add = function() {
        var t, e;
        for (t = 0; t < arguments.length; t++) e = arguments[t], ~this.indexOf(e) || this.elements.splice(this.locationFor(e), 0, e);
        this.length = this.elements.length
      }, t.SortedSet.prototype.toArray = function() {
        return this.elements.slice()
      }, t.SortedSet.prototype.map = function(t, e) {
        return this.elements.map(t, e)
      }, t.SortedSet.prototype.forEach = function(t, e) {
        return this.elements.forEach(t, e)
      }, t.SortedSet.prototype.indexOf = function(t) {
        for (var e = 0, i = this.elements.length, n = i - e, o = e + Math.floor(n / 2), r = this.elements[o]; n > 1;) {
          if (r === t) return o;
          t > r && (e = o), r > t && (i = o), n = i - e, o = e + Math.floor(n / 2), r = this.elements[o]
        }
        return r === t ? o : -1
      }, t.SortedSet.prototype.locationFor = function(t) {
        for (var e = 0, i = this.elements.length, n = i - e, o = e + Math.floor(n / 2), r = this.elements[o]; n > 1;) t > r && (e = o), r > t && (i = o), n = i - e, o = e + Math.floor(n / 2), r = this.elements[o];
        return r > t ? o : t > r ? o + 1 : void 0
      }, t.SortedSet.prototype.intersect = function(e) {
        for (var i = new t.SortedSet, n = 0, o = 0, r = this.length, s = e.length, a = this.elements, u = e.elements; !(n > r - 1 || o > s - 1);) a[n] !== u[o] ? a[n] < u[o] ? n++ : a[n] > u[o] && o++ : (i.add(a[n]), n++, o++);
        return i
      }, t.SortedSet.prototype.clone = function() {
        var e = new t.SortedSet;
        return e.elements = this.toArray(), e.length = e.elements.length, e
      }, t.SortedSet.prototype.union = function(t) {
        var e, i, n;
        this.length >= t.length ? (e = this, i = t) : (e = t, i = this), n = e.clone();
        for (var o = 0, r = i.toArray(); o < r.length; o++) n.add(r[o]);
        return n
      }, t.SortedSet.prototype.toJSON = function() {
        return this.toArray()
      }, t.Index = function() {
        this._fields = [], this._ref = "id", this.pipeline = new t.Pipeline, this.documentStore = new t.Store, this.tokenStore = new t.TokenStore, this.corpusTokens = new t.SortedSet, this.eventEmitter = new t.EventEmitter, this.tokenizerFn = t.tokenizer, this._idfCache = {}, this.on("add", "remove", "update", function() {
          this._idfCache = {}
        }.bind(this))
      }, t.Index.prototype.on = function() {
        var t = Array.prototype.slice.call(arguments);
        return this.eventEmitter.addListener.apply(this.eventEmitter, t)
      }, t.Index.prototype.off = function(t, e) {
        return this.eventEmitter.removeListener(t, e)
      }, t.Index.load = function(e) {
        e.version !== t.version && t.utils.warn("version mismatch: current " + t.version + " importing " + e.version);
        var i = new this;
        return i._fields = e.fields, i._ref = e.ref, i.tokenizer = t.tokenizer.load(e.tokenizer), i.documentStore = t.Store.load(e.documentStore), i.tokenStore = t.TokenStore.load(e.tokenStore), i.corpusTokens = t.SortedSet.load(e.corpusTokens), i.pipeline = t.Pipeline.load(e.pipeline), i
      }, t.Index.prototype.field = function(t, e) {
        var i = {
          name: t,
          boost: (e = e || {}).boost || 1
        };
        return this._fields.push(i), this
      }, t.Index.prototype.ref = function(t) {
        return this._ref = t, this
      }, t.Index.prototype.tokenizer = function(e) {
        return e.label && e.label in t.tokenizer.registeredFunctions || t.utils.warn("Function is not a registered tokenizer. This may cause problems when serialising the index"), this.tokenizerFn = e, this
      }, t.Index.prototype.add = function(e, i) {
        var n = {},
          o = new t.SortedSet,
          r = e[this._ref],
          i = void 0 === i || i;
        this._fields.forEach(function(t) {
          var i = this.pipeline.run(this.tokenizerFn(e[t.name]));
          n[t.name] = i;
          for (var r = 0; r < i.length; r++) {
            var s = i[r];
            o.add(s), this.corpusTokens.add(s)
          }
        }, this), this.documentStore.set(r, o);
        for (var s = 0; s < o.length; s++) {
          for (var a = o.elements[s], u = 0, l = 0; l < this._fields.length; l++) {
            var h = this._fields[l],
              c = n[h.name],
              f = c.length;
            if (f) {
              for (var d = 0, p = 0; f > p; p++) c[p] === a && d++;
              u += d / f * h.boost
            }
          }
          this.tokenStore.add(a, {
            ref: r,
            tf: u
          })
        }
        i && this.eventEmitter.emit("add", e, this)
      }, t.Index.prototype.remove = function(t, e) {
        var i = t[this._ref],
          e = void 0 === e || e;
        if (this.documentStore.has(i)) {
          var n = this.documentStore.get(i);
          this.documentStore.remove(i), n.forEach(function(t) {
            this.tokenStore.remove(t, i)
          }, this), e && this.eventEmitter.emit("remove", t, this)
        }
      }, t.Index.prototype.update = function(t, e) {
        var e = void 0 === e || e;
        this.remove(t, !1), this.add(t, !1), e && this.eventEmitter.emit("update", t, this)
      }, t.Index.prototype.idf = function(t) {
        var e = "@" + t;
        if (Object.prototype.hasOwnProperty.call(this._idfCache, e)) return this._idfCache[e];
        var i = this.tokenStore.count(t),
          n = 1;
        return i > 0 && (n = 1 + Math.log(this.documentStore.length / i)), this._idfCache[e] = n
      }, t.Index.prototype.search = function(e) {
        var i = this.pipeline.run(this.tokenizerFn(e)),
          n = new t.Vector,
          o = [],
          r = this._fields.reduce(function(t, e) {
            return t + e.boost
          }, 0);
        return i.some(function(t) {
          return this.tokenStore.has(t)
        }, this) ? (i.forEach(function(e, i, s) {
          var a = 1 / s.length * this._fields.length * r,
            u = this,
            l = this.tokenStore.expand(e).reduce(function(i, o) {
              var r = u.corpusTokens.indexOf(o),
                s = u.idf(o),
                l = 1,
                h = new t.SortedSet;
              if (o !== e) {
                var c = Math.max(3, o.length - e.length);
                l = 1 / Math.log(c)
              }
              r > -1 && n.insert(r, a * s * l);
              for (var f = u.tokenStore.get(o), d = Object.keys(f), p = d.length, g = 0; p > g; g++) h.add(f[d[g]].ref);
              return i.union(h)
            }, new t.SortedSet);
          o.push(l)
        }, this), o.reduce(function(t, e) {
          return t.intersect(e)
        }).map(function(t) {
          return {
            ref: t,
            score: n.similarity(this.documentVector(t))
          }
        }, this).sort(function(t, e) {
          return e.score - t.score
        })) : []
      }, t.Index.prototype.documentVector = function(e) {
        for (var i = this.documentStore.get(e), n = i.length, o = new t.Vector, r = 0; n > r; r++) {
          var s = i.elements[r],
            a = this.tokenStore.get(s)[e].tf,
            u = this.idf(s);
          o.insert(this.corpusTokens.indexOf(s), a * u)
        }
        return o
      }, t.Index.prototype.toJSON = function() {
        return {
          version: t.version,
          fields: this._fields,
          ref: this._ref,
          tokenizer: this.tokenizerFn.label,
          documentStore: this.documentStore.toJSON(),
          tokenStore: this.tokenStore.toJSON(),
          corpusTokens: this.corpusTokens.toJSON(),
          pipeline: this.pipeline.toJSON()
        }
      }, t.Index.prototype.use = function(t) {
        var e = Array.prototype.slice.call(arguments, 1);
        e.unshift(this), t.apply(this, e)
      }, t.Store = function() {
        this.store = {}, this.length = 0
      }, t.Store.load = function(e) {
        var i = new this;
        return i.length = e.length, i.store = Object.keys(e.store).reduce(function(i, n) {
          return i[n] = t.SortedSet.load(e.store[n]), i
        }, {}), i
      }, t.Store.prototype.set = function(t, e) {
        this.has(t) || this.length++, this.store[t] = e
      }, t.Store.prototype.get = function(t) {
        return this.store[t]
      }, t.Store.prototype.has = function(t) {
        return t in this.store
      }, t.Store.prototype.remove = function(t) {
        this.has(t) && (delete this.store[t], this.length--)
      }, t.Store.prototype.toJSON = function() {
        return {
          store: this.store,
          length: this.length
        }
      }, t.stemmer = function() {
        var t = {
            ational: "ate",
            tional: "tion",
            enci: "ence",
            anci: "ance",
            izer: "ize",
            bli: "ble",
            alli: "al",
            entli: "ent",
            eli: "e",
            ousli: "ous",
            ization: "ize",
            ation: "ate",
            ator: "ate",
            alism: "al",
            iveness: "ive",
            fulness: "ful",
            ousness: "ous",
            aliti: "al",
            iviti: "ive",
            biliti: "ble",
            logi: "log"
          },
          e = {
            icate: "ic",
            ative: "",
            alize: "al",
            iciti: "ic",
            ical: "ic",
            ful: "",
            ness: ""
          },
          i = "[aeiouy]",
          n = "[^aeiou][^aeiouy]*",
          o = i + "[aeiou]*",
          r = new RegExp("^([^aeiou][^aeiouy]*)?[aeiouy][aeiou]*[^aeiou][^aeiouy]*"),
          s = new RegExp("^([^aeiou][^aeiouy]*)?[aeiouy][aeiou]*[^aeiou][^aeiouy]*[aeiouy][aeiou]*[^aeiou][^aeiouy]*"),
          a = new RegExp("^([^aeiou][^aeiouy]*)?[aeiouy][aeiou]*[^aeiou][^aeiouy]*([aeiouy][aeiou]*)?$"),
          u = new RegExp("^([^aeiou][^aeiouy]*)?[aeiouy]"),
          l = /^(.+?)(ss|i)es$/,
          h = /^(.+?)([^s])s$/,
          c = /^(.+?)eed$/,
          f = /^(.+?)(ed|ing)$/,
          d = /.$/,
          p = /(at|bl|iz)$/,
          g = new RegExp("([^aeiouylsz])\\1$"),
          m = new RegExp("^" + n + i + "[^aeiouwxy]$"),
          v = /^(.+?[^aeiou])y$/,
          y = /^(.+?)(ational|tional|enci|anci|izer|bli|alli|entli|eli|ousli|ization|ation|ator|alism|iveness|fulness|ousness|aliti|iviti|biliti|logi)$/,
          S = /^(.+?)(icate|ative|alize|iciti|ical|ful|ness)$/,
          w = /^(.+?)(al|ance|ence|er|ic|able|ible|ant|ement|ment|ent|ou|ism|ate|iti|ous|ive|ize)$/,
          k = /^(.+?)(s|t)(ion)$/,
          x = /^(.+?)e$/,
          b = /ll$/,
          _ = new RegExp("^" + n + i + "[^aeiouwxy]$");
        return function(i) {
          var n, o, E, F, z, I, P;
          if (i.length < 3) return i;
          if ("y" == (E = i.substr(0, 1)) && (i = E.toUpperCase() + i.substr(1)), F = l, z = h, F.test(i) ? i = i.replace(F, "$1$2") : z.test(i) && (i = i.replace(z, "$1$2")), F = c, z = f, F.test(i)) {
            var O = F.exec(i);
            (F = r).test(O[1]) && (F = d, i = i.replace(F, ""))
          } else z.test(i) && (n = (O = z.exec(i))[1], (z = u).test(n) && (i = n, z = p, I = g, P = m, z.test(i) ? i += "e" : I.test(i) ? (F = d, i = i.replace(F, "")) : P.test(i) && (i += "e")));
          return (F = v).test(i) && (i = (n = (O = F.exec(i))[1]) + "i"), (F = y).test(i) && (n = (O = F.exec(i))[1], o = O[2], (F = r).test(n) && (i = n + t[o])), (F = S).test(i) && (n = (O = F.exec(i))[1], o = O[2], (F = r).test(n) && (i = n + e[o])), F = w, z = k, F.test(i) ? (n = (O = F.exec(i))[1], (F = s).test(n) && (i = n)) : z.test(i) && (n = (O = z.exec(i))[1] + O[2], (z = s).test(n) && (i = n)), (F = x).test(i) && (n = (O = F.exec(i))[1], z = a, I = _, ((F = s).test(n) || z.test(n) && !I.test(n)) && (i = n)), F = b, z = s, F.test(i) && z.test(i) && (F = d, i = i.replace(F, "")), "y" == E && (i = E.toLowerCase() + i.substr(1)), i
        }
      }(), t.Pipeline.registerFunction(t.stemmer, "stemmer"), t.generateStopWordFilter = function(t) {
        var e = t.reduce(function(t, e) {
          return t[e] = e, t
        }, {});
        return function(t) {
          return t && e[t] !== t ? t : void 0
        }
      }, t.stopWordFilter = t.generateStopWordFilter(["a", "able", "about", "across", "after", "all", "almost", "also", "am", "among", "an", "and", "any", "are", "as", "at", "be", "because", "been", "but", "by", "can", "cannot", "could", "dear", "did", "do", "does", "either", "else", "ever", "every", "for", "from", "get", "got", "had", "has", "have", "he", "her", "hers", "him", "his", "how", "however", "i", "if", "in", "into", "is", "it", "its", "just", "least", "let", "like", "likely", "may", "me", "might", "most", "must", "my", "neither", "no", "nor", "not", "of", "off", "often", "on", "only", "or", "other", "our", "own", "rather", "said", "say", "says", "she", "should", "since", "so", "some", "than", "that", "the", "their", "them", "then", "there", "these", "they", "this", "tis", "to", "too", "twas", "us", "wants", "was", "we", "were", "what", "when", "where", "which", "while", "who", "whom", "why", "will", "with", "would", "yet", "you", "your"]), t.Pipeline.registerFunction(t.stopWordFilter, "stopWordFilter"), t.trimmer = function(t) {
        return t.replace(/^\W+/, "").replace(/\W+$/, "")
      }, t.Pipeline.registerFunction(t.trimmer, "trimmer"), t.TokenStore = function() {
        this.root = {
          docs: {}
        }, this.length = 0
      }, t.TokenStore.load = function(t) {
        var e = new this;
        return e.root = t.root, e.length = t.length, e
      }, t.TokenStore.prototype.add = function(t, e, i) {
        var i = i || this.root,
          n = t.charAt(0),
          o = t.slice(1);
        return n in i || (i[n] = {
          docs: {}
        }), 0 === o.length ? (i[n].docs[e.ref] = e, void(this.length += 1)) : this.add(o, e, i[n])
      }, t.TokenStore.prototype.has = function(t) {
        if (!t) return !1;
        for (var e = this.root, i = 0; i < t.length; i++) {
          if (!e[t.charAt(i)]) return !1;
          e = e[t.charAt(i)]
        }
        return !0
      }, t.TokenStore.prototype.getNode = function(t) {
        if (!t) return {};
        for (var e = this.root, i = 0; i < t.length; i++) {
          if (!e[t.charAt(i)]) return {};
          e = e[t.charAt(i)]
        }
        return e
      }, t.TokenStore.prototype.get = function(t, e) {
        return this.getNode(t, e).docs || {}
      }, t.TokenStore.prototype.count = function(t, e) {
        return Object.keys(this.get(t, e)).length
      }, t.TokenStore.prototype.remove = function(t, e) {
        if (t) {
          for (var i = this.root, n = 0; n < t.length; n++) {
            if (!(t.charAt(n) in i)) return;
            i = i[t.charAt(n)]
          }
          delete i.docs[e]
        }
      }, t.TokenStore.prototype.expand = function(t, e) {
        var i = this.getNode(t),
          n = i.docs || {},
          e = e || [];
        return Object.keys(n).length && e.push(t), Object.keys(i).forEach(function(i) {
          "docs" !== i && e.concat(this.expand(t + i, e))
        }, this), e
      }, t.TokenStore.prototype.toJSON = function() {
        return {
          root: this.root,
          length: this.length
        }
      },
      function(t, e) {
        "function" == typeof define && define.amd ? define(e) : "object" == typeof exports ? module.exports = e() : t.lunr = e()
      }(this, function() {
        return t
      })
  }(), t.fn.ghostHunter = function(e) {
    var n = t.extend({}, t.fn.ghostHunter.defaults, e);
    if (n.results) return i.init(this, n), i
  }, t.fn.ghostHunter.defaults = {
    resultsData: !1,
    onPageLoad: !0,
    onKeyUp: !1,
    result_template: "<a href='{{link}}'><p><h2>{{title}}</h2><h4>{{prettyPubDate}}</h4></p></a>",
    info_template: "<p>Number of posts found: {{amount}}</p>",
    displaySearchInfo: !0,
    zeroResultsInfo: !0,
    before: !1,
    onComplete: !1,
    includepages: !1,
    filterfields: !1
  };
  var e = function(t) {
      var e = new Date(t),
        i = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
      return e.getDate() + " " + i[e.getMonth()] + " " + e.getFullYear()
    },
    i = {
      isInit: !1,
      init: function(t, e) {
        var i = this;
        this.target = t, this.results = e.results, this.blogData = {}, this.result_template = e.result_template, this.info_template = e.info_template, this.zeroResultsInfo = e.zeroResultsInfo, this.displaySearchInfo = e.displaySearchInfo, this.before = e.before, this.onComplete = e.onComplete, this.includepages = e.includepages, this.filterfields = e.filterfields, this.index = lunr(function() {
          this.field("title", {
            boost: 10
          }), this.field("description"), this.field("link"), this.field("plaintext", {
            boost: 5
          }), this.field("pubDate"), this.field("tag"), this.ref("id")
        }), e.onPageLoad ? window.setTimeout(function() {
          i.loadAPI()
        }, 1) : t.focus(function() {
          i.loadAPI()
        }), t.closest("form").submit(function(e) {
          e.preventDefault(), i.find(t.val())
        }), e.onKeyUp && t.keyup(function() {
          i.find(t.val())
        })
      },
      loadAPI: function() {
        if (this.isInit) return !1;
        var i = this.index,
          n = this.blogData;
        obj = {
          limit: "all",
          include: "tags",
          formats: ["plaintext"]
        }, this.includepages && (obj.filter = "(page:true,page:false)"), t.get(ghost.url.api("posts", obj)).done(function(t) {
          searchData = t.posts, searchData.forEach(function(t) {
            var o = t.tags.map(function(t) {
              return t.name
            });
            null == t.meta_description && (t.meta_description = "");
            var r = o.join(", ");
            r.length < 1 && (r = "undefined");
            var s = {
              id: String(t.id),
              title: String(t.title),
              description: String(t.meta_description),
              plaintext: String(t.plaintext),
              pubDate: String(t.created_at),
              tag: r,
              link: String(t.url)
            };
            s.prettyPubDate = e(s.pubDate);
            var a = e(s.pubDate);
            i.add(s), n[t.id] = {
              title: t.title,
              description: t.meta_description,
              pubDate: a,
              link: t.url
            }
          })
        }), this.isInit = !0
      },
      find: function(e) {
        var i = this.index.search(e),
          n = t(this.results),
          o = [];
        n.empty(), this.before && this.before(), (this.zeroResultsInfo || i.length > 0) && this.displaySearchInfo && n.append(this.format(this.info_template, {
          amount: i.length
        }));
        for (var r = 0; r < i.length; r++) {
          var s = i[r].ref,
            a = this.blogData[s];
          n.append(this.format(this.result_template, a)), o.push(a)
        }
        this.onComplete && this.onComplete(o)
      },
      clear: function() {
        t(this.results).empty(), this.target.val("")
      },
      format: function(t, e) {
        return t.replace(/{{([^{}]*)}}/g, function(t, i) {
          var n = e[i];
          return "string" == typeof n || "number" == typeof n ? n : t
        })
      }
    }
}(jQuery);

/*
 * Instafeed.js by Steven Schobert
 * Version: 1.4.1
 * Generated by CoffeeScript 1.9.3
 */
(function() {
  var e;
  e = function() {
      function e(e, t) {
        var n, r;
        this.options = {
          target: "instafeed",
          get: "popular",
          resolution: "thumbnail",
          sortBy: "none",
          links: !0,
          mock: !1,
          useHttp: !1
        };
        if (typeof e == "object")
          for (n in e) r = e[n], this.options[n] = r;
        this.context = t != null ? t : this, this.unique = this._genKey()
      }
      return e.prototype.hasNext = function() {
        return typeof this.context.nextUrl == "string" && this.context.nextUrl.length > 0
      }, e.prototype.next = function() {
        return this.hasNext() ? this.run(this.context.nextUrl) : !1
      }, e.prototype.run = function(t) {
        var n, r, i;
        if (typeof this.options.clientId != "string" && typeof this.options.accessToken != "string") throw new Error("Missing clientId or accessToken.");
        if (typeof this.options.accessToken != "string" && typeof this.options.clientId != "string") throw new Error("Missing clientId or accessToken.");
        return this.options.before != null && typeof this.options.before == "function" && this.options.before.call(this), typeof document != "undefined" && document !== null && (i = document.createElement("script"), i.id = "instafeed-fetcher", i.src = t || this._buildUrl(), n = document.getElementsByTagName("head"), n[0].appendChild(i), r = "instafeedCache" + this.unique, window[r] = new e(this.options, this), window[r].unique = this.unique), !0
      }, e.prototype.parse = function(e) {
        var t, n, r, i, s, o, u, a, f, l, c, h, p, d, v, m, g, y, b, w, E, S, x, T, N, C, k, L, A, O, M, _, D;
        if (typeof e != "object") {
          if (this.options.error != null && typeof this.options.error == "function") return this.options.error.call(this, "Invalid JSON data"), !1;
          throw new Error("Invalid JSON response")
        }
        if (e.meta.code !== 200) {
          if (this.options.error != null && typeof this.options.error == "function") return this.options.error.call(this, e.meta.error_message), !1;
          throw new Error("Error from Instagram: " + e.meta.error_message)
        }
        if (e.data.length === 0) {
          if (this.options.error != null && typeof this.options.error == "function") return this.options.error.call(this, "No images were returned from Instagram"), !1;
          throw new Error("No images were returned from Instagram")
        }
        this.options.success != null && typeof this.options.success == "function" && this.options.success.call(this, e), this.context.nextUrl = "", e.pagination != null && (this.context.nextUrl = e.pagination.next_url);
        if (this.options.sortBy !== "none") {
          this.options.sortBy === "random" ? M = ["", "random"] : M = this.options.sortBy.split("-"), O = M[0] === "least" ? !0 : !1;
          switch (M[1]) {
            case "random":
              e.data.sort(function() {
                return .5 - Math.random()
              });
              break;
            case "recent":
              e.data = this._sortBy(e.data, "created_time", O);
              break;
            case "liked":
              e.data = this._sortBy(e.data, "likes.count", O);
              break;
            case "commented":
              e.data = this._sortBy(e.data, "comments.count", O);
              break;
            default:
              throw new Error("Invalid option for sortBy: '" + this.options.sortBy + "'.")
          }
        }
        if (typeof document != "undefined" && document !== null && this.options.mock === !1) {
          m = e.data, A = parseInt(this.options.limit, 10), this.options.limit != null && m.length > A && (m = m.slice(0, A)), u = document.createDocumentFragment(), this.options.filter != null && typeof this.options.filter == "function" && (m = this._filter(m, this.options.filter));
          if (this.options.template != null && typeof this.options.template == "string") {
            f = "", d = "", w = "", D = document.createElement("div");
            for (c = 0, N = m.length; c < N; c++) {
              h = m[c], p = h.images[this.options.resolution];
              if (typeof p != "object") throw o = "No image found for resolution: " + this.options.resolution + ".", new Error(o);
              E = p.width, y = p.height, b = "square", E > y && (b = "landscape"), E < y && (b = "portrait"), v = p.url, l = window.location.protocol.indexOf("http") >= 0, l && !this.options.useHttp && (v = v.replace(/https?:\/\//, "//")), d = this._makeTemplate(this.options.template, {
                model: h,
                id: h.id,
                link: h.link,
                type: h.type,
                image: v,
                width: E,
                height: y,
                orientation: b,
                caption: this._getObjectProperty(h, "caption.text"),
                likes: h.likes.count,
                comments: h.comments.count,
                location: this._getObjectProperty(h, "location.name")
              }), f += d
            }
            D.innerHTML = f, i = [], r = 0, n = D.childNodes.length;
            while (r < n) i.push(D.childNodes[r]), r += 1;
            for (x = 0, C = i.length; x < C; x++) L = i[x], u.appendChild(L)
          } else
            for (T = 0, k = m.length; T < k; T++) {
              h = m[T], g = document.createElement("img"), p = h.images[this.options.resolution];
              if (typeof p != "object") throw o = "No image found for resolution: " + this.options.resolution + ".", new Error(o);
              v = p.url, l = window.location.protocol.indexOf("http") >= 0, l && !this.options.useHttp && (v = v.replace(/https?:\/\//, "//")), g.src = v, this.options.links === !0 ? (t = document.createElement("a"), t.href = h.link, t.appendChild(g), u.appendChild(t)) : u.appendChild(g)
            }
          _ = this.options.target, typeof _ == "string" && (_ = document.getElementById(_));
          if (_ == null) throw o = 'No element with id="' + this.options.target + '" on page.', new Error(o);
          _.appendChild(u), a = document.getElementsByTagName("head")[0], a.removeChild(document.getElementById("instafeed-fetcher")), S = "instafeedCache" + this.unique, window[S] = void 0;
          try {
            delete window[S]
          } catch (P) {
            s = P
          }
        }
        return this.options.after != null && typeof this.options.after == "function" && this.options.after.call(this), !0
      }, e.prototype._buildUrl = function() {
        var e, t, n;
        e = "https://api.instagram.com/v1";
        switch (this.options.get) {
          case "popular":
            t = "media/popular";
            break;
          case "tagged":
            if (!this.options.tagName) throw new Error("No tag name specified. Use the 'tagName' option.");
            t = "tags/" + this.options.tagName + "/media/recent";
            break;
          case "location":
            if (!this.options.locationId) throw new Error("No location specified. Use the 'locationId' option.");
            t = "locations/" + this.options.locationId + "/media/recent";
            break;
          case "user":
            if (!this.options.userId) throw new Error("No user specified. Use the 'userId' option.");
            t = "users/" + this.options.userId + "/media/recent";
            break;
          default:
            throw new Error("Invalid option for get: '" + this.options.get + "'.")
        }
        return n = e + "/" + t, this.options.accessToken != null ? n += "?access_token=" + this.options.accessToken : n += "?client_id=" + this.options.clientId, this.options.limit != null && (n += "&count=" + this.options.limit), n += "&callback=instafeedCache" + this.unique + ".parse", n
      }, e.prototype._genKey = function() {
        var e;
        return e = function() {
          return ((1 + Math.random()) * 65536 | 0).toString(16).substring(1)
        }, "" + e() + e() + e() + e()
      }, e.prototype._makeTemplate = function(e, t) {
        var n, r, i, s, o;
        r = /(?:\{{2})([\w\[\]\.]+)(?:\}{2})/, n = e;
        while (r.test(n)) s = n.match(r)[1], o = (i = this._getObjectProperty(t, s)) != null ? i : "", n = n.replace(r, function() {
          return "" + o
        });
        return n
      }, e.prototype._getObjectProperty = function(e, t) {
        var n, r;
        t = t.replace(/\[(\w+)\]/g, ".$1"), r = t.split(".");
        while (r.length) {
          n = r.shift();
          if (!(e != null && n in e)) return null;
          e = e[n]
        }
        return e
      }, e.prototype._sortBy = function(e, t, n) {
        var r;
        return r = function(e, r) {
          var i, s;
          return i = this._getObjectProperty(e, t), s = this._getObjectProperty(r, t), n ? i > s ? 1 : -1 : i < s ? 1 : -1
        }, e.sort(r.bind(this)), e
      }, e.prototype._filter = function(e, t) {
        var n, r, i, s, o;
        n = [], r = function(e) {
          if (t(e)) return n.push(e)
        };
        for (i = 0, o = e.length; i < o; i++) s = e[i], r(s);
        return n
      }, e
    }(),
    function(e, t) {
      return typeof define == "function" && define.amd ? define([], t) : typeof module == "object" && module.exports ? module.exports = t() : e.Instafeed = t()
    }(this, function() {
      return e
    })
}).call(this);

/* Rainbow v2.1.3
 * rainbowco.de
 * included languages: c, csharp, css, generic, html, java, javascript, json, php, python, ruby
 */
! function(e, t) {
  "object" == typeof exports && "undefined" != typeof module ? module.exports = t() : "function" == typeof define && define.amd ? define(t) : e.Rainbow = t()
}(this, function() {
  "use strict";

  function e() {
    return "undefined" != typeof module && "object" == typeof module.exports
  }

  function t(e) {
    var t = e.getAttribute("data-language") || e.parentNode.getAttribute("data-language");
    if (!t) {
      var n = e.className.match(/\blang(?:uage)?-(\w+)/) || e.parentNode.className.match(/\blang(?:uage)?-(\w+)/);
      n && (t = n[1])
    }
    return t ? t.toLowerCase() : null
  }

  function n(e, t, n, r) {
    return (n !== e || r !== t) && (n <= e && r >= t)
  }

  function r(e) {
    return e.replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/&(?![\w\#]+;)/g, "&amp;")
  }

  function a(e, t) {
    for (var n = 0, r = 1; r < t; ++r) e[r] && (n += e[r].length);
    return n
  }

  function o(e, t, n, r) {
    return n >= e && n < t || r > e && r < t
  }

  function i(e) {
    var t = [];
    for (var n in e) e.hasOwnProperty(n) && t.push(n);
    return t.sort(function(e, t) {
      return t - e
    })
  }

  function u(e, t, n, r) {
    var a = r.substr(e);
    return r.substr(0, e) + a.replace(t, n)
  }

  function s(t, Prism) {
    if (e()) return global.Worker = require("webworker-threads").Worker, new Worker(__filename);
    var s = Prism.toString(),
      c = i.toString();
    c += r.toString(), c += n.toString(), c += o.toString(), c += u.toString(), c += a.toString(), c += s;
    var f = c + "\tthis.onmessage=" + t.toString(),
      l = new Blob([f], {
        type: "text/javascript"
      });
    return new Worker((window.URL || window.webkitURL).createObjectURL(l))
  }

  function c(e) {
    function t() {
      self.postMessage({
        id: n.id,
        lang: n.lang,
        result: a
      })
    }
    var n = e.data,
      r = new Prism(n.options),
      a = r.refract(n.code, n.lang);
    if (n.isNode) return t(), void self.close();
    setTimeout(function() {
      t()
    }, 1e3 * n.options.delay)
  }

  function f() {
    return (A || null === k) && (k = s(c, Prism)), k
  }

  function l(e, t) {
    function n(a) {
      a.data.id === e.id && (t(a.data), r.removeEventListener("message", n))
    }
    var r = f();
    r.addEventListener("message", n), r.postMessage(e)
  }

  function d(e, t, n) {
    return function(r) {
      e.innerHTML = r.result, e.classList.remove("loading"), "PRE" === e.parentNode.tagName && e.parentNode.classList.remove("loading"), E && E(e, r.lang), 0 == --t.c && n()
    }
  }

  function g(e) {
    return {
      patterns: M,
      inheritenceMap: C,
      aliases: S,
      globalClass: e.globalClass,
      delay: isNaN(e.delay) ? 0 : e.delay
    }
  }

  function m(e, t) {
    var n = {};
    return "object" == typeof t && (n = t, t = n.language), t = S[t] || t, {
      id: x++,
      code: e,
      lang: t,
      options: g(n),
      isNode: A
    }
  }

  function p(e, n) {
    for (var r = {
        c: 0
      }, a = 0, o = e; a < o.length; a += 1) {
      var i = o[a],
        u = t(i);
      if (!i.classList.contains("rainbow") && u) {
        i.classList.add("loading"), i.classList.add("rainbow"), "PRE" === i.parentNode.tagName && i.parentNode.classList.add("loading");
        var s = i.getAttribute("data-global-class"),
          c = parseInt(i.getAttribute("data-delay"), 10);
        ++r.c, l(m(i.innerHTML, {
          language: u,
          globalClass: s,
          delay: c
        }), d(i, r, n))
      }
    }
    0 === r.c && n()
  }

  function v(e) {
    var t = document.createElement("div");
    t.className = "preloader";
    for (var n = 0; n < 7; n++) t.appendChild(document.createElement("div"));
    e.appendChild(t)
  }

  function h(e, t) {
    t = t || function() {}, e = e && "function" == typeof e.getElementsByTagName ? e : document;
    for (var n = e.getElementsByTagName("pre"), r = e.getElementsByTagName("code"), a = [], o = [], i = 0, u = n; i < u.length; i += 1) {
      var s = u[i];
      v(s), s.getElementsByTagName("code").length ? s.getAttribute("data-trimmed") || (s.setAttribute("data-trimmed", !0), s.innerHTML = s.innerHTML.trim()) : a.push(s)
    }
    for (var c = 0, f = r; c < f.length; c += 1) {
      var l = f[c];
      o.push(l)
    }
    p(o.concat(a), t)
  }

  function b(e) {
    E = e
  }

  function y(e, t, n) {
    C[e] || (C[e] = n), M[e] = t.concat(M[e] || [])
  }

  function w(e) {
    delete C[e], delete M[e]
  }

  function N() {
    for (var e = [], t = arguments.length; t--;) e[t] = arguments[t];
    if ("string" == typeof e[0]) {
      return void l(m(e[0], e[1]), function(e) {
        return function(t) {
          e && e(t.result, t.lang)
        }
      }(e[2]))
    }
    if ("function" == typeof e[0]) return void h(0, e[0]);
    h(e[0], e[1])
  }

  function L(e, t) {
    S[e] = t
  }
  var E, Prism = function Prism(e) {
      function t(e, t) {
        for (var r in h)
          if (r = parseInt(r, 10), n(r, h[r], e, t) && (delete h[r], delete v[r]), o(r, h[r], e, t)) return !0;
        return !1
      }

      function s(t, n) {
        var r = t.replace(/\./g, " "),
          a = e.globalClass;
        return a && (r += " " + a), '<span class="' + r + '">' + n + "</span>"
      }

      function c(e) {
        for (var t = i(v), n = 0, r = t; n < r.length; n += 1) {
          var a = r[n],
            o = v[a];
          e = u(a, o.replace, o.with, e)
        }
        return e
      }

      function f(e) {
        var t = "";
        return e.ignoreCase && (t += "i"), e.multiline && (t += "m"), new RegExp(e.source, t)
      }

      function l(n, r, o) {
        void 0 === o && (o = 0);
        var c = n.pattern;
        if (!c) return !1;
        var l = !c.global;
        c = f(c);
        var d = c.exec(r);
        if (!d) return !1;
        !n.name && n.matches && "string" == typeof n.matches[0] && (n.name = n.matches[0], delete n.matches[0]);
        var g = d[0],
          m = d.index + o,
          b = d[0].length + m;
        if (m === b) return !1;
        if (t(m, b)) return {
          remaining: r.substr(b - o),
          offset: b
        };
        for (var y = i(n.matches), w = 0, N = y; w < N.length; w += 1) {
          var L = N[w];
          ! function(t) {
            var r = d[t];
            if (r) {
              var o = n.matches[t],
                i = o.language,
                c = o.name && o.matches ? o.matches : o,
                f = function(e, n, r) {
                  g = u(a(d, t), e, r ? s(r, n) : n, g)
                };
              if ("string" == typeof o) return void f(r, r, o);
              var l, m = new Prism(e);
              if (i) return l = m.refract(r, i), void f(r, l);
              l = m.refract(r, p, c.length ? c : [c]), f(r, l, o.matches ? o.name : 0)
            }
          }(L)
        }
        return function(e) {
          return n.name && (e = s(n.name, e)), v[m] = {
            replace: d[0],
            with: e
          }, h[m] = b, !l && {
            remaining: r.substr(b - o),
            offset: b
          }
        }(g)
      }

      function d(e, t) {
        for (var n = 0, r = t; n < r.length; n += 1)
          for (var a = r[n], o = l(a, e); o;) o = l(a, o.remaining, o.offset);
        return c(e)
      }

      function g(t) {
        for (var n = e.patterns[t] || []; e.inheritenceMap[t];) t = e.inheritenceMap[t], n = n.concat(e.patterns[t] || []);
        return n
      }

      function m(e, t, n) {
        return p = t, n = n || g(t), d(r(e), n)
      }
      var p, v = {},
        h = {};
      this.refract = m
    },
    M = {},
    C = {},
    S = {},
    T = {},
    x = 0,
    A = e(),
    R = function() {
      return "undefined" == typeof document && "undefined" != typeof self
    }(),
    k = null;
  return T = {
    extend: y,
    remove: w,
    onHighlight: b,
    addAlias: L,
    color: N
  }, A && (T.colorSync = function(e, t) {
    var n = m(e, t);
    return new Prism(n.options).refract(n.code, n.lang)
  }), A || R || document.addEventListener("DOMContentLoaded", function(e) {
    T.defer || T.color(e)
  }, !1), R && (self.onmessage = c), T
});
Rainbow.extend("generic", [{
  matches: {
    1: [{
      name: "keyword.operator",
      pattern: /\=|\+/g
    }, {
      name: "keyword.dot",
      pattern: /\./g
    }],
    2: {
      name: "string",
      matches: {
        name: "constant.character.escape",
        pattern: /\\('|"){1}/g
      }
    }
  },
  pattern: /(\(|\s|\[|\=|:|\+|\.|\{|,)(('|")([^\\\1]|\\.)*?(\3))/gm
}, {
  name: "comment",
  pattern: /\/\*[\s\S]*?\*\/|(\/\/|\#)(?!.*('|").*?[^:](\/\/|\#)).*?$/gm
}, {
  name: "constant.numeric",
  pattern: /\b(\d+(\.\d+)?(e(\+|\-)?\d+)?(f|d)?|0x[\da-f]+)\b/gi
}, {
  matches: {
    1: "keyword"
  },
  pattern: /\b(and|array|as|b(ool(ean)?|reak)|c(ase|atch|har|lass|on(st|tinue))|d(ef|elete|o(uble)?)|e(cho|lse(if)?|xit|xtends|xcept)|f(inally|loat|or(each)?|unction)|global|if|import|int(eger)?|long|new|object|or|pr(int|ivate|otected)|public|return|self|st(ring|ruct|atic)|switch|th(en|is|row)|try|(un)?signed|var|void|while)(?=\b)/gi
}, {
  name: "constant.language",
  pattern: /true|false|null/g
}, {
  name: "keyword.operator",
  pattern: /\+|\!|\-|&(gt|lt|amp);|\||\*|\=/g
}, {
  matches: {
    1: "function.call"
  },
  pattern: /(\w+?)(?=\()/g
}, {
  matches: {
    1: "storage.function",
    2: "entity.name.function"
  },
  pattern: /(function)\s(.*?)(?=\()/g
}]), Rainbow.extend("java", [{
  name: "constant",
  pattern: /\b(false|null|true|[A-Z_]+)\b/g
}, {
  matches: {
    1: "keyword",
    2: "support.namespace"
  },
  pattern: /(import|package)\s(.+)/g
}, {
  name: "keyword",
  pattern: /\b(abstract|assert|boolean|break|byte|case|catch|char|class|const|continue|default|do|double|else|enum|extends|final|finally|float|for|goto|if|implements|import|instanceof|int|interface|long|native|new|package|private|protected|public|return|short|static|strictfp|super|switch|synchronized|this|throw|throws|transient|try|void|volatile|while)\b/g
}, {
  name: "string",
  pattern: /(".*?")/g
}, {
  name: "char",
  pattern: /(')(.|\\.|\\u[\dA-Fa-f]{4})\1/g
}, {
  name: "integer",
  pattern: /\b(0x[\da-f]+|\d+)L?\b/g
}, {
  name: "comment",
  pattern: /\/\*[\s\S]*?\*\/|(\/\/).*?$/gm
}, {
  name: "support.annotation",
  pattern: /@\w+/g
}, {
  matches: {
    1: "entity.function"
  },
  pattern: /([^@\.\s]+)\(/g
}, {
  name: "entity.class",
  pattern: /\b([A-Z]\w*)\b/g
}, {
  name: "operator",
  pattern: /(\+{1,2}|-{1,2}|~|!|\*|\/|%|(?:&lt;){1,2}|(?:&gt;){1,3}|instanceof|(?:&amp;){1,2}|\^|\|{1,2}|\?|:|(?:=|!|\+|-|\*|\/|%|\^|\||(?:&lt;){1,2}|(?:&gt;){1,3})?=)/g
}]), Rainbow.extend("javascript", [{
  name: "selector",
  pattern: /\$(?=\.|\()/g
}, {
  name: "support",
  pattern: /\b(window|document)\b/g
}, {
  name: "keyword",
  pattern: /\b(export|default|from)\b/g
}, {
  name: "function.call",
  pattern: /\b(then)(?=\()/g
}, {
  name: "variable.language.this",
  pattern: /\bthis\b/g
}, {
  name: "variable.language.super",
  pattern: /super(?=\.|\()/g
}, {
  name: "storage.type",
  pattern: /\b(const|let|var)(?=\s)/g
}, {
  matches: {
    1: "support.property"
  },
  pattern: /\.(length|node(Name|Value))\b/g
}, {
  matches: {
    1: "support.function"
  },
  pattern: /(setTimeout|setInterval)(?=\()/g
}, {
  matches: {
    1: "support.method"
  },
  pattern: /\.(getAttribute|replace|push|getElementById|getElementsByClassName|setTimeout|setInterval)(?=\()/g
}, {
  name: "string.regexp",
  matches: {
    1: "string.regexp.open",
    2: {
      name: "constant.regexp.escape",
      pattern: /\\(.){1}/g
    },
    3: "string.regexp.close",
    4: "string.regexp.modifier"
  },
  pattern: /(\/)((?![*+?])(?:[^\r\n\[\/\\]|\\.|\[(?:[^\r\n\]\\]|\\.)*\])+)(\/)(?!\/)([igm]{0,3})/g
}, {
  matches: {
    1: "storage.type",
    3: "entity.function"
  },
  pattern: /(var)?(\s|^)(\S+)(?=\s?=\s?function\()/g
}, {
  matches: {
    1: "keyword",
    2: "variable.type"
  },
  pattern: /(new)\s+(?!Promise)([^\(]*)(?=\()/g
}, {
  name: "entity.function",
  pattern: /(\w+)(?=:\s{0,}function)/g
}, {
  name: "constant.other",
  pattern: /\*(?= as)/g
}, {
  matches: {
    1: "keyword",
    2: "constant.other"
  },
  pattern: /(export)\s+(\*)/g
}, {
  matches: {
    1: "storage.type.accessor",
    2: "entity.name.function"
  },
  pattern: /(get|set)\s+(\w+)(?=\()/g
}, {
  matches: {
    2: "entity.name.function"
  },
  pattern: /(^\s*)(\w+)(?=\([^\)]*?\)\s*\{)/gm
}, {
  matches: {
    1: "storage.type.class",
    2: "entity.name.class",
    3: "storage.modifier.extends",
    4: "entity.other.inherited-class"
  },
  pattern: /(class)\s+(\w+)(?:\s+(extends)\s+(\w+))?(?=\s*\{)/g
}, {
  name: "storage.type.function.arrow",
  pattern: /=&gt;/g
}, {
  name: "support.class.promise",
  pattern: /\bPromise(?=(\(|\.))/g
}], "generic"), Rainbow.addAlias("js", "javascript"), Rainbow.extend("csharp", [{
  name: "constant",
  pattern: /\b(false|null|true)\b/g
}, {
  name: "keyword",
  pattern: /\b(abstract|add|alias|ascending|as|async|await|base|bool|break|byte|case|catch|char|checked|class|const|continue|decimal|default|delegate|descending|double|do|dynamic|else|enum|event|explicit|extern|false|finally|fixed|float|foreach|for|from|get|global|goto|group|if|implicit|int|interface|internal|into|in|is|join|let|lock|long|namespace|new|object|operator|orderby|out|override|params|partial|private|protected|public|readonly|ref|remove|return|sbyte|sealed|select|set|short|sizeof|stackalloc|static|string|struct|switch|this|throw|try|typeof|uint|unchecked|ulong|unsafe|ushort|using|value|var|virtual|void|volatile|where|while|yield)\b/g
}, {
  matches: {
    1: "keyword",
    2: {
      name: "support.class",
      pattern: /\w+/g
    }
  },
  pattern: /(typeof)\s([^\$].*?)(\)|;)/g
}, {
  matches: {
    1: "keyword.namespace",
    2: {
      name: "support.namespace",
      pattern: /\w+/g
    }
  },
  pattern: /\b(namespace)\s(.*?);/g
}, {
  matches: {
    1: "storage.modifier",
    2: "storage.class",
    3: "entity.name.class",
    4: "storage.modifier.extends",
    5: "entity.other.inherited-class"
  },
  pattern: /\b(abstract|sealed)?\s?(class)\s(\w+)(\sextends\s)?([\w\\]*)?\s?\{?(\n|\})/g
}, {
  name: "keyword.static",
  pattern: /\b(static)\b/g
}, {
  matches: {
    1: "keyword.new",
    2: {
      name: "support.class",
      pattern: /\w+/g
    }
  },
  pattern: /\b(new)\s([^\$].*?)(?=\)|\(|;|&)/g
}, {
  name: "string",
  pattern: /(")(.*?)\1/g
}, {
  name: "integer",
  pattern: /\b(0x[\da-f]+|\d+)\b/g
}, {
  name: "comment",
  pattern: /\/\*[\s\S]*?\*\/|(\/\/)[\s\S]*?$/gm
}, {
  name: "operator",
  pattern: /(\+\+|\+=|\+|--|-=|-|&lt;&lt;=|&lt;&lt;|&lt;=|=&gt;|&gt;&gt;=|&gt;&gt;|&gt;=|!=|!|~|\^|\|\||&amp;&amp;|&amp;=|&amp;|\?\?|::|:|\*=|\*|\/=|%=|\|=|==|=)/g
}, {
  name: "preprocessor",
  pattern: /(\#if|\#else|\#elif|\#endif|\#define|\#undef|\#warning|\#error|\#line|\#region|\#endregion|\#pragma)[\s\S]*?$/gm
}]), Rainbow.extend("python", [{
  name: "variable.self",
  pattern: /self/g
}, {
  name: "constant.language",
  pattern: /None|True|False|NotImplemented|\.\.\./g
}, {
  name: "support.object",
  pattern: /object/g
}, {
  name: "support.function.python",
  pattern: /\b(bs|divmod|input|open|staticmethod|all|enumerate|int|ord|str|any|eval|isinstance|pow|sum|basestring|execfile|issubclass|print|super|bin|file|iter|property|tuple|bool|filter|len|range|type|bytearray|float|list|raw_input|unichr|callable|format|locals|reduce|unicode|chr|frozenset|long|reload|vars|classmethod|getattr|map|repr|xrange|cmp|globals|max|reversed|zip|compile|hasattr|memoryview|round|__import__|complex|hash|min|set|apply|delattr|help|next|setattr|buffer|dict|hex|object|slice|coerce|dir|id|oct|sorted|intern)(?=\()/g
}, {
  matches: {
    1: "keyword"
  },
  pattern: /\b(pass|lambda|with|is|not|in|from|elif|raise|del)(?=\b)/g
}, {
  matches: {
    1: "storage.class",
    2: "entity.name.class",
    3: "entity.other.inherited-class"
  },
  pattern: /(class)\s+(\w+)\((\w+?)\)/g
}, {
  matches: {
    1: "storage.function",
    2: "support.magic"
  },
  pattern: /(def)\s+(__\w+)(?=\()/g
}, {
  name: "support.magic",
  pattern: /__(name)__/g
}, {
  matches: {
    1: "keyword.control",
    2: "support.exception.type"
  },
  pattern: /(except) (\w+):/g
}, {
  matches: {
    1: "storage.function",
    2: "entity.name.function"
  },
  pattern: /(def)\s+(\w+)(?=\()/g
}, {
  name: "entity.name.function.decorator",
  pattern: /@([\w\.]+)/g
}, {
  name: "comment.docstring",
  pattern: /('{3}|"{3})[\s\S]*?\1/gm
}], "generic"), Rainbow.extend("c", [{
  name: "meta.preprocessor",
  matches: {
    1: [{
      matches: {
        1: "keyword.define",
        2: "entity.name"
      },
      pattern: /(\w+)\s(\w+)\b/g
    }, {
      name: "keyword.define",
      pattern: /endif/g
    }, {
      name: "constant.numeric",
      pattern: /\d+/g
    }, {
      matches: {
        1: "keyword.include",
        2: "string"
      },
      pattern: /(include)\s(.*?)$/g
    }]
  },
  pattern: /\#([\S\s]*?)$/gm
}, {
  name: "keyword",
  pattern: /\b(do|goto|typedef)\b/g
}, {
  name: "entity.label",
  pattern: /\w+:/g
}, {
  matches: {
    1: "storage.type",
    3: "storage.type",
    4: "entity.name.function"
  },
  pattern: /\b((un)?signed|const)? ?(void|char|short|int|long|float|double)\*? +((\w+)(?= ?\())?/g
}, {
  matches: {
    2: "entity.name.function"
  },
  pattern: /(\w|\*) +((\w+)(?= ?\())/g
}, {
  name: "storage.modifier",
  pattern: /\b(static|extern|auto|register|volatile|inline)\b/g
}, {
  name: "support.type",
  pattern: /\b(struct|union|enum)\b/g
}], "generic"), Rainbow.extend("php", [{
  name: "support",
  pattern: /\becho\b/gi
}, {
  matches: {
    1: "variable.dollar-sign",
    2: "variable"
  },
  pattern: /(\$)(\w+)\b/g
}, {
  name: "constant.language",
  pattern: /true|false|null/gi
}, {
  name: "constant",
  pattern: /\b[A-Z0-9_]{2,}\b/g
}, {
  name: "keyword.dot",
  pattern: /\./g
}, {
  name: "keyword",
  pattern: /\b(die|end(for(each)?|switch|if)|case|require(_once)?|include(_once)?)(?=\b)/gi
}, {
  matches: {
    1: "keyword",
    2: {
      name: "support.class",
      pattern: /\w+/g
    }
  },
  pattern: /(instanceof)\s([^\$].*?)(\)|;)/gi
}, {
  matches: {
    1: "support.function"
  },
  pattern: /\b(array(_key_exists|_merge|_keys|_shift)?|isset|count|empty|unset|printf|is_(array|string|numeric|object)|sprintf|each|date|time|substr|pos|str(len|pos|tolower|_replace|totime)?|ord|trim|in_array|implode|end|preg_match|explode|fmod|define|link|list|get_class|serialize|file|sort|mail|dir|idate|log|intval|header|chr|function_exists|dirname|preg_replace|file_exists)(?=\()/gi
}, {
  name: "variable.language.php-tag",
  pattern: /(&lt;\?(php)?|\?&gt;)/gi
}, {
  matches: {
    1: "keyword.namespace",
    2: {
      name: "support.namespace",
      pattern: /\w+/g
    }
  },
  pattern: /\b(namespace|use)\s(.*?);/gi
}, {
  matches: {
    1: "storage.modifier",
    2: "storage.class",
    3: "entity.name.class",
    4: "storage.modifier.extends",
    5: "entity.other.inherited-class",
    6: "storage.modifier.extends",
    7: "entity.other.inherited-class"
  },
  pattern: /\b(abstract|final)?\s?(class|interface|trait)\s(\w+)(\sextends\s)?([\w\\]*)?(\simplements\s)?([\w\\]*)?\s?\{?(\n|\})/gi
}, {
  name: "keyword.static",
  pattern: /self::|static::/gi
}, {
  matches: {
    1: "storage.function",
    2: "entity.name.function.magic"
  },
  pattern: /(function)\s(__.*?)(?=\()/gi
}, {
  matches: {
    1: "storage.function",
    2: "entity.name.function"
  },
  pattern: /(function)\s(.*?)(?=\()/gi
}, {
  matches: {
    1: "keyword.new",
    2: {
      name: "support.class",
      pattern: /\w+/g
    }
  },
  pattern: /\b(new)\s([^\$][a-z0-9_\\]*?)(?=\)|\(|;)/gi
}, {
  matches: {
    1: {
      name: "support.class",
      pattern: /\w+/g
    },
    2: "keyword.static"
  },
  pattern: /([\w\\]*?)(::)(?=\b|\$)/g
}, {
  matches: {
    2: {
      name: "support.class",
      pattern: /\w+/g
    }
  },
  pattern: /(\(|,\s?)([\w\\]*?)(?=\s\$)/g
}], "generic"), Rainbow.extend("ruby", [{
  matches: {
    1: "variable.language",
    2: {
      language: null
    }
  },
  pattern: /^(__END__)\n((?:.*\n)*)/gm
}, {
  name: "string",
  matches: {
    1: "string.open",
    2: [{
      name: "string.interpolation",
      matches: {
        1: "string.open",
        2: {
          language: "ruby"
        },
        3: "string.close"
      },
      pattern: /(\#\{)(.*?)(\})/g
    }],
    3: "string.close"
  },
  pattern: /("|`)(.*?[^\\\1])?(\1)/g
}, {
  name: "string",
  pattern: /('|"|`)([^\\\1\n]|\\.)*?\1/g
}, {
  name: "string",
  pattern: /%[qQ](?=(\(|\[|\{|&lt;|.)(.*?)(?:'|\)|\]|\}|&gt;|\1))(?:\(\2\)|\[\2\]|\{\2\}|\&lt;\2&gt;|\1\2\1)/g
}, {
  matches: {
    1: "string",
    2: "string",
    3: "string"
  },
  pattern: /(&lt;&lt;)(\w+).*?$([\s\S]*?^\2)/gm
}, {
  matches: {
    1: "string",
    2: "string",
    3: "string"
  },
  pattern: /(&lt;&lt;\-)(\w+).*?$([\s\S]*?\2)/gm
}, {
  name: "string.regexp",
  matches: {
    1: "string.regexp",
    2: {
      name: "string.regexp",
      pattern: /\\(.){1}/g
    },
    3: "string.regexp",
    4: "string.regexp"
  },
  pattern: /(\/)(.*?)(\/)([a-z]*)/g
}, {
  name: "string.regexp",
  matches: {
    1: "string.regexp",
    2: {
      name: "string.regexp",
      pattern: /\\(.){1}/g
    },
    3: "string.regexp",
    4: "string.regexp"
  },
  pattern: /%r(?=(\(|\[|\{|&lt;|.)(.*?)('|\)|\]|\}|&gt;|\1))(?:\(\2\)|\[\2\]|\{\2\}|\&lt;\2&gt;|\1\2\1)([a-z]*)/g
}, {
  name: "comment",
  pattern: /#.*$/gm
}, {
  name: "comment",
  pattern: /^\=begin[\s\S]*?\=end$/gm
}, {
  matches: {
    1: "constant"
  },
  pattern: /(\w+:)[^:]/g
}, {
  matches: {
    1: "constant.symbol"
  },
  pattern: /[^:](:(?:\w+|(?=['"](.*?)['"])(?:"\2"|'\2')))/g
}, {
  name: "constant.numeric",
  pattern: /\b(0x[\da-f]+|[\d_]+)\b/g
}, {
  name: "support.class",
  pattern: /\b[A-Z]\w*(?=((\.|::)[A-Za-z]|\[))/g
}, {
  name: "constant",
  pattern: /\b[A-Z]\w*\b/g
}, {
  matches: {
    1: "storage.class",
    2: "entity.name.class",
    3: "entity.other.inherited-class"
  },
  pattern: /\s*(class)\s+((?:(?:::)?[A-Z]\w*)+)(?:\s+&lt;\s+((?:(?:::)?[A-Z]\w*)+))?/g
}, {
  matches: {
    1: "storage.module",
    2: "entity.name.class"
  },
  pattern: /\s*(module)\s+((?:(?:::)?[A-Z]\w*)+)/g
}, {
  name: "variable.global",
  pattern: /\$([a-zA-Z_]\w*)\b/g
}, {
  name: "variable.class",
  pattern: /@@([a-zA-Z_]\w*)\b/g
}, {
  name: "variable.instance",
  pattern: /@([a-zA-Z_]\w*)\b/g
}, {
  matches: {
    1: "keyword.control"
  },
  pattern: /[^\.]\b(BEGIN|begin|case|class|do|else|elsif|END|end|ensure|for|if|in|module|rescue|then|unless|until|when|while)\b(?![?!])/g
}, {
  matches: {
    1: "keyword.control.pseudo-method"
  },
  pattern: /[^\.]\b(alias|alias_method|break|next|redo|retry|return|super|undef|yield)\b(?![?!])|\bdefined\?|\bblock_given\?/g
}, {
  matches: {
    1: "constant.language"
  },
  pattern: /\b(nil|true|false)\b(?![?!])/g
}, {
  matches: {
    1: "variable.language"
  },
  pattern: /\b(__(FILE|LINE)__|self)\b(?![?!])/g
}, {
  matches: {
    1: "keyword.special-method"
  },
  pattern: /\b(require|gem|initialize|new|loop|include|extend|raise|attr_reader|attr_writer|attr_accessor|attr|catch|throw|private|module_function|public|protected)\b(?![?!])/g
}, {
  name: "keyword.operator",
  pattern: /\s\?\s|=|&lt;&lt;|&lt;&lt;=|%=|&=|\*=|\*\*=|\+=|\-=|\^=|\|{1,2}=|&lt;&lt;|&lt;=&gt;|&lt;(?!&lt;|=)|&gt;(?!&lt;|=|&gt;)|&lt;=|&gt;=|===|==|=~|!=|!~|%|&amp;|\*\*|\*|\+|\-|\/|\||~|&gt;&gt;/g
}, {
  matches: {
    1: "keyword.operator.logical"
  },
  pattern: /[^\.]\b(and|not|or)\b/g
}, {
  matches: {
    1: "storage.function",
    2: "entity.name.function"
  },
  pattern: /(def)\s(.*?)(?=(\s|\())/g
}]), Rainbow.extend("html", [{
  name: "source.php.embedded",
  matches: {
    1: "variable.language.php-tag",
    2: {
      language: "php"
    },
    3: "variable.language.php-tag"
  },
  pattern: /(&lt;\?php|&lt;\?=?(?!xml))([\s\S]*?)(\?&gt;)/gm
}, {
  name: "source.css.embedded",
  matches: {
    1: {
      matches: {
        1: "support.tag.style",
        2: [{
          name: "entity.tag.style",
          pattern: /^style/g
        }, {
          name: "string",
          pattern: /('|")(.*?)(\1)/g
        }, {
          name: "entity.tag.style.attribute",
          pattern: /(\w+)/g
        }],
        3: "support.tag.style"
      },
      pattern: /(&lt;\/?)(style.*?)(&gt;)/g
    },
    2: {
      language: "css"
    },
    3: "support.tag.style",
    4: "entity.tag.style",
    5: "support.tag.style"
  },
  pattern: /(&lt;style.*?&gt;)([\s\S]*?)(&lt;\/)(style)(&gt;)/gm
}, {
  name: "source.js.embedded",
  matches: {
    1: {
      matches: {
        1: "support.tag.script",
        2: [{
          name: "entity.tag.script",
          pattern: /^script/g
        }, {
          name: "string",
          pattern: /('|")(.*?)(\1)/g
        }, {
          name: "entity.tag.script.attribute",
          pattern: /(\w+)/g
        }],
        3: "support.tag.script"
      },
      pattern: /(&lt;\/?)(script.*?)(&gt;)/g
    },
    2: {
      language: "javascript"
    },
    3: "support.tag.script",
    4: "entity.tag.script",
    5: "support.tag.script"
  },
  pattern: /(&lt;script(?! src).*?&gt;)([\s\S]*?)(&lt;\/)(script)(&gt;)/gm
}, {
  name: "comment.html",
  pattern: /&lt;\!--[\S\s]*?--&gt;/g
}, {
  matches: {
    1: "support.tag.open",
    2: "support.tag.close"
  },
  pattern: /(&lt;)|(\/?\??&gt;)/g
}, {
  name: "support.tag",
  matches: {
    1: "support.tag",
    2: "support.tag.special",
    3: "support.tag-name"
  },
  pattern: /(&lt;\??)(\/|\!?)(\w+)/g
}, {
  matches: {
    1: "support.attribute"
  },
  pattern: /([a-z-]+)(?=\=)/gi
}, {
  matches: {
    1: "support.operator",
    2: "string.quote",
    3: "string.value",
    4: "string.quote"
  },
  pattern: /(=)('|")(.*?)(\2)/g
}, {
  matches: {
    1: "support.operator",
    2: "support.value"
  },
  pattern: /(=)([a-zA-Z\-0-9]*)\b/g
}, {
  matches: {
    1: "support.attribute"
  },
  pattern: /\s([\w-]+)(?=\s|&gt;)(?![\s\S]*&lt;)/g
}]), Rainbow.addAlias("xml", "html"), Rainbow.extend("css", [{
  name: "comment",
  pattern: /\/\*[\s\S]*?\*\//gm
}, {
  name: "constant.hex-color",
  pattern: /#([a-f0-9]{3}|[a-f0-9]{6})(?=;|\s|,|\))/gi
}, {
  matches: {
    1: "constant.numeric",
    2: "keyword.unit"
  },
  pattern: /(\d+)(px|em|cm|s|%)?/g
}, {
  name: "string",
  pattern: /('|")(.*?)\1/g
}, {
  name: "support.css-property",
  matches: {
    1: "support.vendor-prefix"
  },
  pattern: /(-o-|-moz-|-webkit-|-ms-)?[\w-]+(?=\s?:)(?!.*\{)/g
}, {
  matches: {
    1: [{
      name: "entity.name.sass",
      pattern: /&amp;/g
    }, {
      name: "direct-descendant",
      pattern: /&gt;/g
    }, {
      name: "entity.name.class",
      pattern: /\.[\w\-_]+/g
    }, {
      name: "entity.name.id",
      pattern: /\#[\w\-_]+/g
    }, {
      name: "entity.name.pseudo",
      pattern: /:[\w\-_]+/g
    }, {
      name: "entity.name.tag",
      pattern: /\w+/g
    }]
  },
  pattern: /([\w\ ,\n:\.\#\&\;\-_]+)(?=.*\{)/g
}, {
  matches: {
    2: "support.vendor-prefix",
    3: "support.css-value"
  },
  pattern: /(:|,)\s*(-o-|-moz-|-webkit-|-ms-)?([a-zA-Z-]*)(?=\b)(?!.*\{)/g
}]), Rainbow.addAlias("scss", "css"), Rainbow.extend("json", [{
  matches: {
    0: {
      name: "string",
      matches: {
        name: "constant.character.escape",
        pattern: /\\('|"){1}/g
      }
    }
  },
  pattern: /(\"|\')(\\?.)*?\1/g
}, {
  name: "constant.numeric",
  pattern: /\b(-?(0x)?\d*\.?[\da-f]+|NaN|-?Infinity)\b/gi
}, {
  name: "constant.language",
  pattern: /\b(true|false|null)\b/g
}]);
