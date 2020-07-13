default rel

section .rodata
msg: db "Hello, world!", 0

section .text
global hello
hello:
    lea rax, [msg]
    ret
