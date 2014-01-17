{ p8 } = require('../p8')

main = ->
    p = new p8.Page()
    asm = new p8.PageAsm(p)
    asm.printi 1987
    asm.printi 1
    asm.printi 21
    asm.printi 32
    asm.label "so"

    vm = new p8.Vm()
    vm.mapPage p8.pageHead(1), p
    vm.pc = p8.pageHead 1
    vm.ttl = 5

    e = vm.resume()

    console.log "(vm stops, e=" + e + ")"
    console.log "tsc="+vm.tsc, "ttl="+vm.ttl

main()
