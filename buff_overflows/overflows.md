# Buffer Overflow Notes
## For OSCP and beyond
Notes for detecting and exploiting buffer overflow vulnerabilities.

## Further reading: 
- [Assembly Programming](https://tutorialspoint.com/assembly_programming/)
- [More on Registers](https://wiki.skullsecurity.org/Registers)

## Memory Addresses

- EBX: Base pointer for memory addresses.
- ECX: Loop, shift, rotation counter
- EDX: I/O port addressing, multiplication, division
- ESI: Pointer addressing of data and source in string copy operations
- EDI: Pointer addressing of data and destination in string copy operations
- ESP: Since data in the stack is dynamic the stack pointer holds the last used address (on top of the stack) and holds a pointer to it.
- EBP: Since the stack has a constant flux of information, functions can somtimes loose track of their own stack frame, which stores all arguments, cariables and the return address. This pointer hold the address of the top of the stack, which means that functions can quickly access their own stack frame by using offsets. 
- EIP: The instruction pointer points to the next code instruction that will be executed, because of this it is one of the most important registers. 

## Useful shortcuts (Immunity Debugger)
- F2 - Set breakpoint
- F7/F8 - Step into/Step over