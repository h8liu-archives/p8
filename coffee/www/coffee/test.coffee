if false
    px = (n) -> '' + n + 'px'

    main = $("#main")[0]
    ref = $("#ref")
    dpr = window.devicePixelRatio
    ctx = main.getContext "2d"
    fontsize = 13

    # ctx.scale first and then zoom with width, height
    # the order here is critical

    drawText = (message) ->
        height = fontsize * dpr
        ctx.font = '' + height + 'px Consolas'
        width = ctx.measureText('M').width
        chars = message.split('')
        x = 0
        y = height
        ctx.fillStyle = "#c00" # red
        for c in chars
            ctx.clearRect x, y - height, width, height
            ctx.fillText c, x, y
            x += width

    drawText('hello')

    savedHeight = 0
    savedWidth = 0
    timer = ->
        winHeight = $(window).height()
        winWidth = $(window).width()
        if winHeight != savedHeight || winWidth != savedWidth
            console.log winHeight, winWidth
            main.style.width = px(winWidth)
            main.style.height = px(winHeight)
            ctx.scale dpr, dpr
            main.width = winWidth * dpr
            main.height = winHeight * dpr
            m = 'some ' + winWidth + 'x' + winHeight
            drawText(m)
            drawText("res: ")
            savedHeight = winHeight
            savedWidth = winWidth
        return true

    window.requestAnimationFrame(timer)

timer = (f) ->
    repeat = (t) ->
        ret = f(t)
        if ret
            window.requestAnimationFrame repeat
    window.requestAnimationFrame repeat
    return

Terminal = (canvas, _dpr) ->
    thiz = this
    px = (n) -> '' + n + 'px'
    ctx = canvas.getContext '2d'
    fontSize = 13
    charWidth = 0
    charHeight = 0
    curWidth = 0
    curHeight = 0
    dpr = _dpr

    this.resize = (w, h) ->
        if w == curWidth && h == curHeight
            return false
        canvas.style.width = px(w)
        canvas.style.height = px(h)
        ctx.scale dpr, dpr
        canvas.width = w * dpr
        canvas.height = h * dpr
        charHeight = fontSize * dpr
        ctx.font = '' + charHeight + 'px Consolas'
        charWidth = ctx.measureText('M').width
        curWidth = w
        curHeight = h
        return true

    this.sizeStr = ->
        return '' + curWidth + 'x' + curHeight
    
    this.fillWindow = (window) ->
        width = window.innerWidth
        height = window.innerHeight
        return thiz.resize(width, height)

    this.putChar = (x, y, c) ->
        _x = x * charWidth
        _y = y * charHeight
        if _x < 0 || _x + charWidth > curWidth
            return
        if _y < 0 || _y + charHeight > curHeight
            return
        ctx.clearRect _x, _y, charWidth, charHeight
        ctx.fillText c, _x, _y + charHeight

    this.print = (msg) ->
        chars = msg.split('')
        x = 0
        y = 0
        ctx.font = '' + charHeight + 'px Consolas'
        ctx.fillStyle = "#c00" # red
        for c in chars
            thiz.putChar x, y, c
            x += 1
        return
    
    return

term = new Terminal($("#main")[0], window.devicePixelRatio)
timer ->
    if term.fillWindow(window)
        term.print term.sizeStr()
    return true

