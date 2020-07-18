package main

type Memory struct {
    data [0x10000]uint8
}

// return a single byte from the memory
func (ram Memory) GetByte (index uint) uint {
    return uint(ram.data[index])
}

// return two bytes from the memory (usually an address)
func (ram Memory) GetAddress (index uint) uint {
    high := uint(ram.data[index    ])
    low  := uint(ram.data[index + 1])
    return (high << 8) | low
}


// write a byte in the memory
func (ram Memory) Write (index, value uint) {
    ram.data[index] = uint8(value)
}
