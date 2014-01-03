**Something New**

So I just saw some (very prelimitary) documents talking about how to write a
LLVM backend for CPU0, which is a very simple ISA. So, it might be just easier
in the end to just write a simmulator for CPU0 (or something similar), and then
we can leverage the large chunk of LLVM stuff, and compile C programs.

Another way I am thinking about is to port tcc (A Tiny C Compiler), try to
understand the code and rewrite it in my favourite golang, and by then, I will
have a full fledged compiler that I can say that I understand how it works.

Neither way has a huge learning curve, but comparing with the developing effort
of writing a compiler from scratch, it might be worth trying. LLVM is a of
course a huge trunk of code, but somehow modularized into reusable tools. tcc
is in fact not that tiny anyway. I personally prefer tcc somehow because that
allows you to have control on everything, rather than sit on a large chunk
of modules that you don't (and probably never going to) understand how it works.

**What is `m8`**

I am still brainstorming on this, but basically I want to build a simple
Virtual machine. It does not need to be all from scratch, but it needs to be
simple enough that I can understand how it works exactly.

Computer systems today are huge and complex, and even as a senior PhD student
on sysnet/networks computer science, I only understand how things work in
general concepturally, yet always feel like I have no idea how it works in
details. This often makes me uncomfortable, in several ways.

First is security. Lots of security features are essentially based on merely
trust on the implementations. However, when trust is gone for some reasons,
it is practically impossible to check and verity the features, since the
implementations are often complex and hard to understand. Even with open source
programs, the source can be easily too complicate to reason about. Instead of
open, what we really need here is understandable source programs.

Second is education. Schools and universities teach simple concepts on
architecturs, compilers, operating systems, networks, and students writes toy
programs based on the teaching materials in separate courses. None of these
toys work together as a complete system that students can understand how they
really interacts. Then the students are thrown to real systems which are much
more complex and impossible to understand the whole picture.

Third is research. Research is about proposing and validating prototype ideas
with scientific evaluation. However, experiments and benchmarks on real systems
are just hard to repeat becauss modern computers are complex and depends on all
kinds of random factors. As a result, when evaluating new research ideas, every
researcher just build its own benchmark setup and how that others believe the
results validates the idea. 

Final one is personal. I need to tell a story for this. 

When I was in collage, I was once in charge of developing a platform for an
annual game AI competition. The competition goes like this: every year, a small
group of organizers (often tech nerd students like me) design a simple game
(often based ancient games like PacMan, Tetris, etc.), while every participants
submits an AI for the game. All the AIs run in a realtime gaming arena based on
the game rules, compete for the winning prices. Almost every year, a new but
imperfect game AI arena is developed. Sometimes the inter-procedure
communication was based on linking DLLs files, sometimes based on pipes, and
sometimes on networks links. However, since it is always based on existing
operating systems environment (a.k.a. Windows), fariness on computational
resource allocation has always been a problem, especially when the game
approaches the finals and heats up. Since cycles allocated to each AIs are
limited, a really competitive AI often wants to squeeze all the cycles
available, but at the same time can easily miss the action deadline if it is
not careful, and from the arena platform side, providing a real-time API that
can tell the API how many cycles is in its pocket is very hard in most
non-realtime OS runtimes, not saying to guarantee fairness among different
running AIs. If it was using DLLs running on the same machine, an AI might miss
an action deadline due to scheduling gitters; if it was running on separate
machines, an AI might miss the deadline due to network latencies. Now imagine
the final comes, two AIs compete for the throne, and one extremely powerful AI
misses an action deadline on a critical movement and lose the game; this
actually happened.

I blame the platform for this issue, and when I was in charge on implementing
this, I always wanted to fix that, but find it very hard to fix perfectly. It
becomes one of my complexes. It sounds like a feature request that is so
simple: allocate a certain but hard number of CPU cycles, memory spaces to a
single-threaded program written in C(or any other popular language) and run it
in a sandbox. Over these years, I continue searching for an SDK like this but
fails. 

**Why not just use Lua?**

Lua might be the mature environment that best fits my requirements, but after
looking into it for a while, I gave up.

First, it (seems) does not come with an instruction counter. Although this
might be easy to add.

Second, it it already pretty complicated. Lua 5.x has around 20000 lines of C
code, which is not easy to understand and reason. This is what I hate about
large softwares, they evolve over time, and becomes a large pile of code that
tangled together, and it is normally impossible to just use part of the code
to work with something new because it is not well modularized. Even if it is,
the modular structure is usually implicit, inter-module interfaces are often
not well-defined, and hence hard for a third person to understand it completely
and use just part of it to work with something new.

So I think I will just start from scratch and see if I can do a better job on
writing human understandable programs.
