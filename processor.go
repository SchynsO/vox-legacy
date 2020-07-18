package main


import (
    "math/bits"
)


type Processor struct {
    reg  [4]uint8 // A, X, Y, Z
    flag [4]bool  // zero, overflow, carry, negative
    ptr     uint
    ram    *Memory
    stack  *Stack
}


// read and execute one instruction from the memory
func (cpu Processor) Cycle () {
    // read the byte at the specified location
    inst := cpu.ram.GetByte(cpu.ptr)
    reg  := inst & 0x3 // register to use
    cpu.ptr += 1 // move to next byte

    // lower values are NOP
    if between(0x40, inst, 0xA8)  {
        if        inst < 0x48 { // STR
            if inst < 0x44 { // STR from registers
                addr := cpu.ram.GetAddress(cpu.ptr)
                cpu.reg[reg] = uint8(cpu.ram.GetByte(addr))
                cpu.ptr += 2
            } else         { // STR with indexing
                cpu.reg[ 0 ] = uint8(cpu.readMR(inst))
            }
        } else if inst < 0x54 { // LOD
            if inst < 0x4C        { // LOD in registers
                addr := cpu.ram.GetAddress(cpu.ptr)
                cpu.reg[reg] = uint8(cpu.ram.GetByte(addr))
                cpu.ptr += 2
            } else if inst < 0x50 { // LOD with indexing
                cpu.reg[ 0 ] = uint8(cpu.readMR(inst))
            } else                { // LOD numbers
                cpu.reg[reg] = uint8(cpu.ram.GetByte(cpu.ptr))
                cpu.ptr += 1
            }
        } else if inst < 0x58 { // PSH
            cpu.stack.Push(cpu.reg[reg])
        } else if inst < 0x5C { // PLL
            cpu.reg[reg] = cpu.stack.Pull()
        } else if inst < 0x60 { // JMP & RTN
            if inst < 0x5E { // JMP
                addr := cpu.ram.GetAddress(cpu.ptr)
                if inst == 0x5D {cpu.stack.PushAddress(cpu.ptr + 1)}
                cpu.ptr = addr
            } else         { // RTN
                cpu.ptr = cpu.stack.PullAddress()
            }
        } else if inst < 0x70 { // TRS
            reg2 := (inst & 0xC) >> 2
            cpu.reg[reg] = cpu.reg[reg2]
        } else if inst < 0x78 { // INC
            cpu.writeMR(inst, cpu.readMR(inst) + 1)
        } else if inst < 0x80 { // DEC
            cpu.writeMR(inst, cpu.readMR(inst) - 1)
        } else if inst < 0x88 { // SHL
            cpu.writeMR(inst, cpu.readMR(inst) << 1)
        } else if inst < 0x90 { // SHR
            cpu.writeMR(inst, cpu.readMR(inst) >> 1)
        } else if inst < 0x98 { // ROL
            read := uint8(cpu.readMR(inst))
            cpu.writeMR(inst, uint(bits.RotateLeft8(read,  1)))
        } else if inst < 0xA0 { // ROR
            read := uint8(cpu.readMR(inst))
            cpu.writeMR(inst, uint(bits.RotateLeft8(read, -1)))
        } else                { // NOT
            cpu.writeMR(inst, ^cpu.readMR(inst))
        }
    } else if between(0xB0, inst, 0xF0) {
        if        inst < 0xB8 { // BRC
            // detect if branching
            cond := cpu.flag[reg]
            not  := inst >= 0xB4
            if (cond && !not) || (!cond && not) {
                cpu.ptr = cpu.ram.GetAddress(cpu.ptr)
            } else {
                cpu.ptr += 2
            }
        } else if inst < 0xBC { // SET
            cpu.flag[reg] = true
        } else if inst < 0xC0 { // CLR
            cpu.flag[reg] = false
        } else if inst < 0xC8 { // ADD
            cpu.reg[0] += uint8(cpu.readMR(inst))
        } else if inst < 0xD0 { // SUB
            cpu.reg[0] -= uint8(cpu.readMR(inst))
        } else if inst < 0xD8 { // AND
            cpu.reg[0] &= uint8(cpu.readMR(inst))
        } else if inst < 0xE0 { // IOR
            cpu.reg[0] |= uint8(cpu.readMR(inst))
        } else if inst < 0xE8 { // XOR
            cpu.reg[0] ^= uint8(cpu.readMR(inst))
        } else                { // CMP
            // TODO...
        }
    }
}

// helper to read from memory or registers
func (cpu Processor) readMR (inst uint) uint {
    reg := inst & 0x3
    if (inst & 0x4) == 0 { // read from a register
        return uint(cpu.reg[reg])
    } else {
        addr := cpu.ram.GetAddress(cpu.ptr)
        if reg == 0 { // read from memory
            return cpu.ram.GetByte(addr)
        } else { // read from memory with index
            return cpu.ram.GetByte(addr + uint(cpu.reg[reg]))
        }
        cpu.ptr += 2
    }
    return 0
}

// helper to write to memory or registers
func (cpu Processor) writeMR (inst, value uint) {
    reg := inst & 0x3
    if (inst & 0x4) == 0 { // write to a register
        cpu.reg[reg] = uint8(value)
    } else {
        addr := cpu.ram.GetAddress(cpu.ptr)
        if reg == 0 { // write to memory
            cpu.ram.Write(addr, value)
        } else { // write to memory with index
            cpu.ram.Write(addr + uint(cpu.reg[reg]), value)
        }
        cpu.ptr += 2
    }
}


// specify if the pointer has reached the end of memory
func (cpu Processor) ReachedEnd () bool {
    if cpu.ptr >= 0x10000 {
        cpu.ptr = 0x0
        return true
    }
    return false
}


func between (min, val, max uint) bool {
    return min <= val && val < max
}
