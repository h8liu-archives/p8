**Project P8**

This project plans to construct a simple virtual computer.

The goal of the project, if any, is to build a simple virtual world where all
its code pieces are modularized, easy to understand, reason about and hence
easy to maintain and reuse. In other words, code quality is over everything,
all code pieces should present crystal clear human-understandable logic.
Performance and size optimizations are almost always the last concerns.

So, basically, this is like a minecraft game for a CS SysNet PhD (hopefully
more insteresting and meaningful than those other computer games that ends up
with absolutely nothing). 

**Tentative Plan**

Design new arch and languages:

- P8 is a register-based machine that runs a simple RISC ISA
- IR8 is an SSA-like intermediate representation for a compiler
- G is a high-level programming language

Write some code for the new arch and languages:

- `p8/arch` P8 architecture
- `p8/vm` a simulator for P8
- `p8/asm` a P8 assembler and deassembler
- `p8/i8` an IR8 interpreter
- `p8/i8c` a compiler from IR8 to P8
- `p8/gc` a compiler from G to P8 via IR8
- `p8/web` a web-based user interface for the machine
- `p8/web/vm` the vm interpreter but in javascript
- `p8/cc` a full-fledged C compiler, compile to P8 via IR8
- `p8/goc` a full-fledged golang compiler, compile to P8 via IR8

At this time, all the previous code written in go language should be able to
port into P8 framework without modification. (An alternative plan could be
rewrite everything above in G language.)

I know this project just sounds crazy... so please don't expect me to finish it
in 100 years.

**Recruitement**

If you are also interested and want to help, please contact me. liulonnie@gmail.com
