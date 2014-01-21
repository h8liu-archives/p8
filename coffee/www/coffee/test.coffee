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

startTimer = (f) ->
    timer = (t) ->
        ret = f(t)
        if ret
            window.requestAnimationFrame timer
    window.requestAnimationFrame timer
    return

Terminal = (canvas, _dpr) ->
    thiz = this
    c = canvas
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
            return
        c.style.width = px(w)
        c.style.height = px(h)
        ctx.scale dpr, dpr
        canvas.width = w * dpr
        canvas.height = h * dpr
        charHeight = fontSize * dpr
        charWidth = ctx.measureText('M').width
        ctx.font = 'Consolas ' + charHeight + 'px'
        return
    
    this.fillWindow = (window) ->
        width = window.innerWidth
        height = window.innerHeight
        thiz.resize(width, height)
        return

    this.putChar = (x, y, c) ->
        _x = x * charWidth
        _y = y * charHeight
        ctx.clearRect _x, _y, charWidth, charHeight
        ctx.fillText c, _x, _y + charHeight

    this.print = (msg) ->
        chars = msg.split('')
        x = 0
        y = 0
        console.log charHeight
        ctx.font = '' + charHeight + 'px Consolas'
        ctx.fillStyle = "#c00" # red
        for c in chars
            thiz.putChar x, y, c
            x += 1
        return
    
    return

term = new Terminal $("#main")[0], window.devicePixelRatio
startTimer ( -> 
    term.fillWindow(window) 
    term.print('hello')
    return true
)

