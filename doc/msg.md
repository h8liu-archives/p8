** message **

- Messaging is the way a vm communicates with the world outside. 
- Messages are sent in via system pipes
- Messages are only received by pulling

recv
- input: pipe id y
- input: map addr p
- output: return value p
  - a value
- output: return error x
  - pipe empty (0x1)
  - mapping failure (when y is non zero)  (0x2)
  - pipe invalid (0x4)

send
- input: pipe id y
- input: map addr p
- input: value q
- output: return error x
  - pipe full (0x1)
  - share failure (when y in non zero)
  - pipe invalid (0x4)

syscall
- input: call x
- input: parameter y
- input: parameter p
- input: parameter q
- return: x
- return value y
- return value p
- return value q

syscalli
- input: call i
- input: parameter y
...

default pipes
- pipe 0, int stream // a stream of uint64 integers
- pipe 1, tty stream // a stream of uint64 integers

- 0 reseved
- 1 new page (addr, perm) -> error
- 2 del page (addr) -> error
- 3 pipe recv
- 4 pipe send

so, to send a char out via vm
- save the char in a register, say register 1

addi 1, 0, 4 // send out
addi 2, 0, 0 // pipe id
ld 3, 14, 0 // load the int
syscalli 1, 2, 3, 0

