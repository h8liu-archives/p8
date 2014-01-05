/*
Package arch defines the op code for P8 RISC architecture.

A P8 CPU has a PC register and 16 64-bit registers. Register $0 is always 0.

It also has a 64-bit time stamp counter (TSC) read-only reg that increases by 1
every cycle, and a 64-bit time to live (TTL) read-only reg that decreases by 1
every cycle if it is not zero already.  In the simulator, both registers
(might) change by 1 every instruction. When the TTL register changes from 1 to
0, the machine halts.

Address length is also 64-bit. On memory alignment error or invalid
errors, the machine halts. Memory space is virtual, split into 4K pages. A page
can we executable or not, writable or not. Page 0 must always be invalid.

All memory accesses must be proper aligned.

P8 is little endian.

Each instruction is 64-bit long. The highest 16 bits is the opcode. The
instruction is a jump when the highest bit of the opcode is 1, the jump will
save the PC in $15 if the second highest bit of the opcode is also 1. For
details of the opcode, see comments of the opcode definitions.

TODO: In the future, it also will support 64-bit floating point calculations.

There is no ring protection or interrupt handling mechanisms within a VM.

TODO: Processes will be put in separate VMs, VMs communicates via register
based messages and shared pages. So yes, a page can be mapped to multiple VMs.

TODO: VM0 is the kernel VM. It listens on all kinds of system events (by default),
and manage other VMs via system calls (open, kill, pause, resume, map memory,
map io devices). On bootaing, Page1 will be loaded into VM0 as writable and
executable. When VM0 halts, the machine halts.
*/
package risc

// system
const (
	Halt  = iota // halt
	Rdtsc        // x = tsc
	Rdttl        // x = ttl
)

// calculations
const (
	Add  = 0x100 + iota // x = p + q
	Addi                // x = y + signed(i)
	Sub                 // x = p - q
	Lui                 // x[high] = i, set high 32 bits of x
	And                 // x = p & q
	Andi                // x = y & unsigned(i)
	Or                  // x = p | q
	Ori                 // x = y | unsigned(i)
	Xor                 // x = p xor q
	Nor                 // x = p nor q
	Slt                 // x = p < q ? 1 : 0
	Slti                // x = (y < signed(i)) ? 1 : 0
	Sll                 // x = p << q!
	Srl                 // x = p >> q!, unsigned
	Sra                 // x = p >> q!, signed
	Sllv                // x = p << q
	Srlv                // x = p >> q, unsigned
	Srav                // x = p >> q, signed
)

// jumps
const (
	Jr  = 0x200 + iota // pc = p
	Beq                // if x == y, pc += signed(i)*8
	Bne                // if x != y, pc += signed(i)*8
)

// muls and divs
const (
	Mul  = 0x300 + iota // x = p * q, signed
	Mulu                // x = p * q, unsigned
	Div                 // (x, y) = (p / q, p % q), signed
	Divu                // (x, y) = (p / q, p % q), unsigned
)

// memory ops
const (
	Ld  = 0x400 + iota // x = [y + signed(i)], double word
	Lw                 // x = [y + signed(i)], signed word
	Lwu                // x = [y + signed(i)], unsigned word
	Lh                 // x = [y + signed(i)], signed half word
	Lhu                // x = [y + signed(i)], unsigned half word
	Lb                 // x = [y + signed(i)], signed byte
	Lbu                // x = [y + signed(i)], unsigned byte
	Sd                 // [y + signed(i)] = x, double word
	Sw                 // [y + signed(i)] = x, word
	Sh                 // [y + signed(i)] = x, half word
	Sb                 // [y + signed(i)] = x, byte
)

// immediate jumps
const (
	J   = 0x8000 // pc = I<<3
	Jal = 0x4000 // $15=pc, pc = I<<3
)
