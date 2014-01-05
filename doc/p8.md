# bytecode design

16 32bit registers
$0 is always 0
// regs
// 0ccc xypq
0000. halt
0001. jr pc=p
0002. add x=p+q
0003. sub x=p-q 
0004. and x=p&q
0005. or  x=p|q
0006. xor x=p^q
0007. nor x=^(p|q)
0008. slt x=p<q
0009. sll x=p<<!q
000A. srl x=p>>!q
000B. sra x=p>>!q // signed
000C. sllv x=p<<q
000D. srlv x=p>>q
000E. srav x=p>>q // signed
000F. // reserved

0010. mul  [x;y]=p*q // signed
0011. mulu [x;y]=p*q // unsigned
0012. div  (x,y)=(p/q,p%q) // signed
0013. divu (x,y)=(p/q,p%q) // unsigned

// immediates
// 1cxy iiii
10xy. addi x=y+se(i)
11xy. andi x=y&ze(i)
12xy. ori  x=y|ze(i)
13xy. slti x=y<se(i)
14xy. lw   x=[y+se(i)]
15xy. lh  x=[y+se(i)] // signed
16xy. lhu x=[y+se(i)]
17xy. lb  x=[y+se(i)] // signed
18xy. lbu x=[y+se(i)]
19xy. lui x=ze(i) << 16
1Axy. sw  [y+se(i)]=x
1Bxy. sh [y+se(i)]=x
1Cxy. sb [y+se(i)]=x
1Dxy. beq if (x==y) pc+=se(i)<<2
1Exy. bne if (x!=y) pc+=se(i)<<2

// fpus // reserved
// 2... ....

// syscall, ios // reserved
// 3... ....

// reserved
// (4-7)

// (8-B) j pc=I<<2
// (C-F) jal $15=pc; pc=I<<2
