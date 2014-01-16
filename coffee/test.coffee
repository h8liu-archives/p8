p8 = require('./p8').p8

main = ->
    p = new p8.Page()
    w = new p8.PageWriter(p)
    asm = new p8.Asm()
    w.u32 asm.printi 1987
    w.u32 asm.printi 1
    w.u32 asm.printi 21

    vm = new p8.Vm()
    vm.mapPage p8.pageHead(1), p
    vm.pc = p8.pageHead 1
    vm.ttl = 5

    e = vm.resume()

    console.log "(vm stops, e=" + e + ")"
    console.log "tsc="+vm.tsc, "ttl="+vm.ttl

main()
