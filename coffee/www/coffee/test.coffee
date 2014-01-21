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
    for c in chars
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
        m = '' + winWidth + 'x' + winHeight
        drawText(m)
        savedHeight = winHeight
        savedWidth = winWidth
    window.requestAnimationFrame timer

window.requestAnimationFrame(timer)
