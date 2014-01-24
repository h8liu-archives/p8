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
    termWidth = 0
    termHeight = 0
    nchar = 0
    nline = 0
    
    cursorX = 0
    cursorY = 0
    dpr = _dpr

    this.resize = (w, h) ->
        if w == termWidth && h == termHeight
            return false
        canvas.style.width = px(w)
        canvas.style.height = px(h)
        ctx.scale dpr, dpr
        canvas.width = w * dpr
        canvas.height = h * dpr
        charHeight = fontSize * dpr
        ctx.font = '' + charHeight + 'px Consolas'
        charWidth = ctx.measureText('M').width
        termWidth = w
        termHeight = h
        nchar = Math.floor(w * dpr / charWidth)
        nline = Math.floor(h * dpr / charHeight)
        console.log nchar, nline
        return true
    this.resize $(canvas).width(), $(canvas).height()

    this.sizeStr = ->
        return '' + termWidth + 'x' + termHeight
    
    this.fillWindow = (window) ->
        width = window.innerWidth
        height = window.innerHeight
        return thiz.resize(width, height)

    this.putChar = (x, y, c) ->
        _x = x * charWidth
        _y = y * charHeight
        if _x < 0 || _x + charWidth > termWidth
            return
        if _y < 0 || _y + charHeight > termHeight
            return
        ctx.clearRect _x, _y, charWidth, charHeight
        ctx.fillText c, _x, _y + charHeight

    this.print = (msg) ->
        chars = msg.split('')
        ctx.font = '' + charHeight + 'px Consolas'
        ctx.fillStyle = "#000" # red
        for c in chars
            thiz.putChar cursorX, cursorY, c
            cursorX += 1
        return
    
    return

leftTerm = new Terminal($("#left")[0], window.devicePixelRatio)
rightTerm = new Terminal($("#right")[0], window.devicePixelRatio)
leftTerm.print "hello, this is the left terminal."
rightTerm.print "hello, this is the right terminal."

# timer ->
#     if term.fillWindow(window)
#         term.print term.sizeStr()
#     return true

