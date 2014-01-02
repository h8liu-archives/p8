# stage 1

- only integer registers, no fpu
- assembler and simulator
- simulator dumps all the registers at the end of running of the program

# go-stage1 design

// go state1 is a language that is similar to assumbly language
// and should be easy to compile to our bytecode
func main() {
    mov(1, 2)
    mov(2, 3)
    movi(3, 5)
    add(3, 4, 5)
loop:
    if eq(2, 4) { goto loop }
    if ne(2, 3) { goto loop }
    
    halt()
    addi(3, 4, 37323)
    
    goto somewhere
}

