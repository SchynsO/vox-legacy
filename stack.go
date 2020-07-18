package main


type Stack struct {
    data [0x100]uint8
    ptr  int
}


func (s Stack) Push (value uint8) {
    s.data[s.ptr] = value
    s.ptr += 1
}

func (s Stack) Pull () uint8 {
    s.ptr -= 1
    val := s.data[s.ptr]
    return val
}

func (s Stack) PushAddress (addr uint) {
    s.data[s.ptr    ] = uint8(addr >> 8)
    s.data[s.ptr + 1] = uint8(addr)
    s.ptr += 2
}

func (s Stack) PullAddress () uint {
    s.ptr -= 2
    high := uint(s.data[s.ptr    ])
    low  := uint(s.data[s.ptr + 1])
    return (high << 8) | low
}


func (s Stack) InBound () bool {
    return 0 <= s.ptr && s.ptr < 0x100
}
