"use strict";

(function() {
    var errHalt = 1;
    var errAddr = 2;
    var errDeath = 3;

    function todo() { console.error("todo"); }

    var pageSize = 4096;
    var pageShift = 12;

    function pageId(a) { return a >> pageShift; }
    function pageHead(i) { return i << pageShift; }
    function pageAlign(a) { return a >> pageShift << pageShift; }
    function pageOff(a) { return a & ((1 << pageShift) - 1); }

    function a16(p) { return p >> 1 << 1; }
    function a32(p) { return p >> 2 << 2; }
    function a64(p) { return p >> 3 << 3; }

    function Page() {
        var buffer = new ArrayBuffer(pageSize);
        var view = new DataView(buffer);

        this.i8 = function(p) { return view.getInt8(p); }
        this.u8 = function(p) { return view.getUint8(p); }
        this.i16 = function(p) { return view.getInt16(a16(p)); }
        this.u16 = function(p) { return view.getUint16(a16(p)); }
        this.i32 = function(p) { return view.getInt32(a32(p));  }
        this.u32 = function(p) { return view.getUint32(a32(p));  }
        this.f64 = function(p) { return view.getFloat64(a64(p));  }
        this.pi8 = function(p, v) { view.setInt8(p, v); }
        this.pu8 = function(p, v) { view.setUint8(p, v); }
        this.pi16 = function(p, v) { view.setInt16(a16(p), v); }
        this.pu16 = function(p, v) { view.setUint16(a16(p), v); }
        this.pi32 = function(p, v) { view.setInt32(a32(p), v);  }
        this.pu32 = function(p, v) { view.setUint32(a32(p), v); }
        this.pf64 = function(p, v) { view.setFloat64(a64(p), v); }
    }

    function NoopPage() {
        function r0(p) { return 0; }
        function noop(p, v) {}

        this.i8 = r0;
        this.u8 = r0;
        this.i16 = r0;
        this.u16 = r0;
        this.i32 = r0;
        this.u32 = r0;
        this.f64 = r0;

        this.pi8 = noop;
        this.pu8 = noop;
        this.pi16 = noop;
        this.pu16 = noop;
        this.pi32 = noop;
        this.pu32 = noop;
        this.pf64 = noop;
    }

    function Asm() {
        var thiz = this;
        var byname = {}; // from name to Op
        var bycode = {};
        var nop = 0;

        function makeInst(x, p, q, i) {
            var ret = 0;
            ret = ret | (i & 0xffff);
            ret = ret | ((q & 0x7) << 16);
            ret = ret | ((p & 0x7) << 19);
            ret = ret | ((x & 0x7) << 22);
            return ret;
        }

        function execnoop() { }

        function add(name, str, make, parse) {
            var code = nop;
            var op = {
                code: code,
                name: name,
                exec: execnoop,
                str: str,
                make: make,
                parse: parse,
            };
            byname[name] = op;
            bycode[nop] = op;
            thiz[name] = function(a, b, c, d, e) {
                var ret = make(a, b, c, d, e);
                ret |= code << 25;
                return ret;
            }

            nop += 1;
        }

        function add0(name) {
            add(name, function(p) { return name; },
                    function() { return makeInst(0, 0, 0, 0); },
                    function(line) { todo(); });
        }

        function add0Ims(name) {
            add(name, function(p) { return name; },
                    function() { return name + " " + p.ims; },
                    function() { return makeInst(0, 0, 0, ims); },
                    function(line) { todo(); });
        }

        add0("halt");
        add0Ims("printi");

        this.code = function(name) { 
            if (name in byname) {
                return byname[name]
            }
            return false;
        }
    }

    function VM() {
        var thiz = this;
        var pages = {};
        var nreg = 8;
        var iregs = new Int32Array(nreg);
        var uregs = new Uint32Array(iregs);
        var e = 0;
        var fakePage = new Page();

        // fetch the page for address
        function page(a) {
            var id = pageId(a);
            if (id in pages) {
                return pages[id];
            }
            e = errAddr;
            return fakePage;
        }

        // check alignments
        function c16(a) { if (a16(a) != a) e = errAddr; }
        function c32(a) { if (a32(a) != a) e = errAddr; }
        function c64(a) { if (a64(a) != a) e = errAddr; }

        // memory operations
        function i8(a) { return page(a).i8(pageOff(a)); }
        function u8(a) { return page(a).u8(pageOff(a)); }
        function i16(a) { c16(a); return page(a).i16(pageOff(a)); }
        function u16(a) { c16(a); return page(a).u16(pageOff(a)); }
        function i32(a) { c32(a); return page(a).i32(pageOff(a)); }
        function u32(a) { c32(a); return page(a).u32(pageOff(a)); }
        function f64(a) { c64(a); return page(a).u64(pageOff(a)); }

        function pi8(a, v) { return page(a).pi8(pageOff(a)); }
        function pu8(a, v) { return page(a).pu8(pageOff(a)); }
        function pi16(a, v) { c16(a); return page(a).pi16(pageOff(a)); }
        function pu16(a, v) { c16(a); return page(a).pu16(pageOff(a)); }
        function pi32(a, v) { c32(a); return page(a).pi32(pageOff(a)); }
        function pu32(a, v) { c32(a); return page(a).pu32(pageOff(a)); }
        function pf64(a, v) { c64(a); return page(a).pf64(pageOff(a)); }

        var ops = [];
        var asm = {};

        function o(name, exec, str, make) {
            var op = {
                code: ops.length,
                exec: exec,
                str: str,
                make: make,
            };

            ops.push(op);
            asm[name] = function(a, b, c, d, e) {
                var ret = op.make(a, b, c, d, e);
                ret |= op.code << 25;
                return ret
            };
        }

        function makeArgs(x, p, q, i) {
            var ret = 0;
            ret = ret | (i & 0xffff);
            ret = ret | ((q & 0x7) << 16);
            ret = ret | ((p & 0x7) << 19);
            ret = ret | ((x & 0x7) << 22);
            return ret;
        }
        function o0(name, f) { o(name, f, 
                function(p) { return name; }, 
                function() { return makeArgs(0, 0, 0, 0) }); }
        function oims(name, f) { o(name, f, 
                function(p) { return name + " " + p.ims; }, 
                function(ims) { return makeArgs(0, 0, 0, ims) }); }

        o0("halt", function(p) { e = errHalt });
        oims("printi", function(p) { thiz.printi(p.ims) });

        function exec(inst, oldPc) {
            var opcode = (inst >> 25) & 0x7f;
            var parts = {
                code: opcode,
                x: (inst >> 22) & 0x7,
                p: (inst >> 19) & 0x7,
                q: (inst >> 16) & 0x7,
                imu: (inst & 0xffff),
                ims: (inst << 16 >> 16),
            }

            if (opcode in ops) {
                var op = ops[opcode];
                thiz.log(oldPc, inst, op.str(parts))
                op.exec(parts);
            } else {
                console.log(oldpc, "noop")
            }
        }

        function step() {
            var inst = u32(thiz.pc);
            var oldPc = thiz.pc;
            if (e != 0) return;
            thiz.pc = a32(thiz.pc + 4);
            exec(inst, oldPc);
            uregs[0] = 0; // clear reg0

            thiz.tsc += 1;
            if (thiz.ttl > 0) {
                thiz.ttl -= 1;
                if (thiz.ttl == 0) {
                    e = errDeath;
                }
            }
        }

        this.error = function() { return e; }
        this.pc = pageHead(1);
        this.ttl = 0;
        this.tsc = 0;

        this.mapPage = function(a, p) {
            var id = pageId(a);
            if (id == 0) return;
            pages[pageId(a)] = p;
        }

        this.resume = function() { e = 0; while (e == 0) step(); return e; }
        this.step = function() { e = 0; step(); return e; }
        this.printi = function(i) { console.log(i); }
        this.log = function(pc, inst, s) { }
        this.asm = asm;
    }

    function PageWriter(page) {
        var p = page;
        var t = this;
        this.off = 0;

        function offinc(i) { t.off = pageOff(t.off + i); }
        this.i8 = function(v) { p.pi8(this.off, v); offinc(1); }
        this.u8 = function(v) { p.pi8(this.off, v); offinc(1); }
        this.i16 = function(v) { p.pi16(this.off, v); offinc(2); }
        this.u16 = function(v) { p.pu16(this.off, v); offinc(2); }
        this.i32 = function(v) { p.pi32(this.off, v); offinc(4); }
        this.u32 = function(v) { p.pu32(this.off, v); offinc(4); }
        this.f64 = function(v) { p.pf64(this.off, v); offinc(8); }
    }

    exports.p8 = {
        errHalt: errHalt,
        errAddr: errAddr,
        errDeath: errDeath,
        VM: VM,
        PageWriter: PageWriter,
        Page: Page,
        pageHead: pageHead,
    }
})();
