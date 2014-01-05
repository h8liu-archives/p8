I plan to write a C compiler in Golang here. For learning and fun.

I know that there are a lot of great C compilers out there, and there are also
great compiler frameworks like LLVM. This project only aims to write a compiler
where the program is clearly architectured with reusable modules, in a
beautiful idiomatic mordern language.

Implementing a full C standard from the start would be hard, even for C89/C90.
So my plan is just implement a small very small subset of C first and grow it
gradually.

The compiler will compile the small C language to something like SSA, and then
to a very simple RISC that I defined. There will be SSA and RISC interpreters
and/or simulators for running the compiled programs. I might consider compile
it to x86 binaries in the future, but I feel like I need to make it running
first, for debugging and debugging purposes.

I will probably start writing the RISC and the SSA first.
