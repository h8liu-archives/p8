main = $("#main")[0]
logger = $("logger")[0]
ctx = main.getContext "2d"
ctx.font = '13px Consolas'
ctx.fillText "Hello, world!", 0, 13
ctx.fillText "And hello again!", 0, 26