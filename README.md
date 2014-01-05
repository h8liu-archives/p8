**p8**

The project plan to construct a simple virtual computer.

New Languages:

- p8 is a register-based machine that runs a simple RISC ISA
- i8 is an SSA-like intermediate representation for a compiler
- G is a high-level programming language

Parts that will be writen in Go language:

- `p8/vm` a simulator for p8
- `p8/asm` a p8 assembler and deassembler
- `p8/i8` an i8 interpreter
- `p8/i8c` a compiler from i8 to p8
- `p8/gc` a compiler from G to p8 via i8
- `p8/web` a web-based user interface for the machine
- `p8/web/vm` the vm interpreter but in javascript
- `p8/cc` a full-fledged C compiler, compile to p8 via i8
- `p8/goc` a full-fledged golang compiler, compile to p8 via i8

Parts that will be writen in G language:

- `g/sys` a thin system layer to use in p8 machine, like an OS

At this time, all the previous code written in go language should be able to
port into p8 framework without modification.

An alternative plan could be rewrite everything above in G language

This project just sounds crazy... so please don't expect me to finish it...

The goal of the project is build a simple virtual world where all its code
pieces are easy to understand, reasonn about and reuse. Code quality is over
everything. Performance and size optimizations are almost always the last 
things that need to consider.
