# bytecode design

16 32bit registers
$0 is always 0

// i0: 0000 000c
0. halt

// i4: 0000 00cx
1. jr pc=x

// i8: 0000 0cyx

// i12: 0000 czyx
1. add x=y+z
2. sub x=y-z 
3. and x=y&z
4. or  x=y|z
5. xor x=y^z
6. nor x=^(y|z)
7. slt x=y<z
8. sll x=y<<!z
9. srl x=y>>!z
10. sra x=y>>!z // signed
11. sllv x=y<<z
12. srlv x=y>>z
13. srav x=y>>z // signed

// i16: 000c qpyx
1. mul  [xy]=p*q // signed
2. mulu [xy]=p*q
3. div  (x,y)=(p/q,p%q) // signed
4. divu (x,y)=(p/q,p/q)

// i20: 00ci iiix
1. lui x=ze(i) << 16

// i24: 0cii iiyx
1. addi x=y+se(i)
2. andi x=y&ze(i)
3. ori  x=y|ze(i)
4. slti x=y<se(i)
5. ld   x=[y+se(i)]
6. ldh  x=[y+se(i)] // signed
7. ldhu x=[y+se(i)]
8. ldb  x=[y+se(i)] // signed
9. ldbu x=[y+se(i)]
10. st  [y+se(i)]=x
11. sth [y+se(i)]=x
12. stb [y+se(i)]=x
13. beq if (x==y) pc+=se(i)<<2
14. bne if (x!=y) pc+=se(i)<<2

// i28: ciii iiii
- 0001 syscall (parameters)
- 01ii. j pc=i<<2
- 10ii. jal $15=pc; pc=i<<2