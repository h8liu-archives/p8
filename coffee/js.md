** requirement of the project **

- a sandbox that runs a single program
- can limit the number of instructions a sandbox can use
- the sandbox can communicate with the outside
- several sandboxes can run at the same time
- should be easy to develop programs for the sandbox
- the sandbox should be accessible via the web

** solution **

- write a simple virtual machine
- since it should be accessible from the web, it should be written in javascript (rather than golang)

** some problems **
javascript only do 32bit integers well, but we have a 64bit integers

we can redesign the asm though
for example, we can say, all registers are float64

registers are all 32bits
32bit instructions
8 32bit registers

32bit instruction
7bit op; 9bit (rd, rx, ry); 16bit im

syscalls
rd
0 halt
1 print
2 input
3 rand

4 readtime
5 readtll
6 newpage
7 delpage

8 ropen
9 rclose
10 rsend
11 rrecv

12 popen
13 pclose
14 psend
15 precv

// standard channels
// 0:  cap = 1
// 1: sync sys message exchange       cap = 0
// 2: async sys message exchange      cap = 1

// pipe capa

0 halt
1 output  rx im
2 outputp rx im
2 input   rd im
3 inputp  rd im

1 add     rd = rx + ry
2 sub     rd = rx - ry

3 and     rd = rx & ry
4 or      rd = rx | ry
5 xor     rd = rx ^ ry
6 nor     rd = ~(rx | ry)

7 slt     rd = rx < ry ? 1 : 0
8 sll     rd = rx << ry
9 srl     rd = rx >> ry
10 sra     rd = rx >>> ry

20 slti    rd = (rx < ims) ? 1 : 0
22 slli    rd = rx << imu
23 srli    rd = rx >> imu (unsigned)
24 srai    rd = rx >>> imu (signed)

17 andi    rd = rx & imu
18 ori     rd = rx | imu
19 addi    rd = rx + ims
21 lui     rd = imu << 16

25 beq     if rx == ry: pc += ims
26 bne     if rx != ry: pc += ims
27 beqal   if rx == ry: pc = rd
28 bneql   if rx != ry: pc = rd

29 lw      rd = [rx + ims]
30 lhs     rd = [rx + ims]
31 lhu     
32 lbs
33 lbu
34 sw
35 sh
36 sb

11 mul     rd = rx * ry
12 mulu    rd = rx * ry (unsigned)
13 div     rd = rx / ry
14 divu    rd = rx / ry
15 mod     rd = rx % ry
16 modu    rd = rx % ry 

37 fadd    fd = fp + fq
38 fsub    fd = fp - fq
39 fmul    fd = fp * fq
40 fdiv    fd = fp / fq

41 fslt    rd = (fp < fq) ? 1 : 0
42 lf      fd = [rx + ims]
43 sf      [rx + ims] = fd

// max to 63

10- j        pc=I<<2
11- jal      $15=pc; pc=I<<2

