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

height = 13 * dpr
ctx.font = '' + height + 'px Consolas'
width = ctx.measureText('M').width

message = 'Hello, world!'
ctx.fillText message, 0, height
chars = message.split('')
x = 0
y = height * 2
for c in chars
    ctx.fillText c, x, y
    x += width
