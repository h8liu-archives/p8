**p8**

The project plan to construct a simple virtual computer.

The goal of the project, if any, is to build a simple virtual world where all
its code pieces are modularized, easy to understand, reason about and hence
easy to maintain and reuse. In other words, code quality is over everything,
all code pieces should present crystal clear human-understandable logic.
Performance and size optimizations are almost always the last things that need
to consider.

So, basically, this is like a minecraft game for a CS SysNet PhD (hopefully
more insteresting and meaningful than those other computer games that ends up
with absolutely nothing). 

**plan (tentative)**

Design new arch and languages:

- p8 is a register-based machine that runs a simple RISC ISA
- i8 is an SSA-like intermediate representation for a compiler
- G is a high-level programming language

Write some code for the new arch and languages:

- `p8/vm` a simulator for p8
- `p8/dasm` a p8 deassembler
- `p8/asm` a p8 assembler
- `p8/i8` an i8 interpreter
- `p8/i8c` a compiler from i8 to p8
- `p8/gc` a compiler from G to p8 via i8
- `p8/web` a web-based user interface for the machine
- `p8/web/vm` the vm interpreter but in javascript
- `p8/cc` a full-fledged C compiler, compile to p8 via i8
- `p8/goc` a full-fledged golang compiler, compile to p8 via i8

At this time, all the previous code written in go language should be able to
port into p8 framework without modification. (An alternative plan could be
rewrite everything above in G language.)

I know this project just sounds crazy... so please don't expect me to finish it
in 100 years.

**recruitement**

If you are also interested and want to help, please contact me. liulonnie@gmail.com
