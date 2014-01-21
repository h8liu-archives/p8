px = (n) -> '' + n + 'px'

main = $("#main")[0]
logger = $("logger")[0]
dpr = window.devicePixelRatio
console.log dpr
ctx = main.getContext "2d"

# ctx.scale first and then zoom with width, height
# the order here is critical
ctx.scale dpr, dpr
main.width = 1200 * dpr
main.height = 600 * dpr
main.style.width = px(1200)
main.style.height = px(600)

drawText = ->
    message = 'Hello, world!'
    height = 13 * dpr
    ctx.font = '' + height + 'px Consolas'
    width = ctx.measureText('M').width
    console.log width
    ctx.fillText message, 0, height
    chars = message.split('')
    x = 0
    y = height * 2
    for c in chars
        ctx.fillText c, x, y
        x += width

drawText()

savedHeight = 0
savedWidth = 0
timer = ->
    winHeight = $(window).height()
    winWidth = $(window).width()
    if winHeight != savedHeight || winWidth != savedWidth
        console.log winHeight, winWidth
        ctx.canvas.width = winWidth
        ctx.canvas.height = winHeight
        main.width = winWidth * dpr
        main.height = winHeight * dpr
        drawText()
        savedHeight = winHeight
        savedWidth = winWidth
    window.requestAnimationFrame timer

window.requestAnimationFrame(timer)
