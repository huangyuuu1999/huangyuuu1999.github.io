# ISA指令集架构

ISA == Instruction Set Architecture 指令集架构
ISA定义了寄存器，指令，寻址方式，基本数据类型，异常中断等。

最重要的就是有哪些寄存器，有哪些指令。

CISC 复杂指令集
RISC 精简指令集

riscv是精简指令集。

## RISCV的ISA命名方式
riscv有很多版本的ISA，例如RV32，RV64
这里的32，64是指令集字款，就是通用寄存器的宽度，注意不是指令的宽度。
riscv是分层设计的，所以很多指令模块都是渐进式的，比如有一个I集合只包含整除操作的指令，F表示单精度浮点数的操作指令，D是双精度浮点数的指令，等等，按需组合。

A是原子指令集合 C是压缩指令集合...

IMAFD = G， RV64GC表示支持 IMAFD，C这些指令集，的64位。

## 寄存器
32个通用寄存器，x0-x31, PC寄存器（程序指针）
RV32的寄存器宽度是32位，RV64是64位，但是指令都是32位的，这一点要注意。

这些通用寄存器，的使用规范，在ABI（application binary interface）中定义，例如每个寄存器实际上有自己的名字，和使用时的惯例。

## 特权级别
用两个比特来表示特权级别，有三种特权级别，machine，supervisor，user。
supervisor是用来实现具有虚拟内存和进程概念的操作系统的级别，就是os课上学到的内核模式。
machain类似于裸金属模式，级别最高，能操作的寄存器最多。

CSR(control status register) 控制状态寄存器。不同的特权级别具备各自的csr。