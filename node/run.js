var p8 = require('./p8.js').p8;

function main() {
    var p = new p8.Page(); // create a page
    var w = new p8.PageWriter(p); // TODO: a better assembler
    var asm = new p8.Asm(); // fetch the assembler
    w.u32(asm.printi(1987));
    w.u32(asm.printi(1));
    w.u32(asm.printi(21));

    var vm = new p8.VM();
    vm.mapPage(p8.pageHead(1), p);
    vm.pc = p8.pageHead(1);
    vm.ttl = 5;

    var e = vm.resume();

    console.log("(vm stops, e=" + e + ")");
    console.log("tsc="+vm.tsc, "ttl="+vm.ttl);
}

main();
