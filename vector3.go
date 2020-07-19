package main


// represent a data structure with 3 components
type Vector3 struct {
    coords [3]uint
}

func (v Vector3) Set (x, y, z uint) {
    v.coords[0] = x
    v.coords[1] = y
    v.coords[2] = z
}

func (v Vector3) Set8 (x, y, z uint8) {
    v.coords[0] = uint(x)
    v.coords[1] = uint(y)
    v.coords[2] = uint(z)
}

// convert a single byte into three components
func (v Vector3) SetByte (byte uint8) {
    b := uint(byte)
    v.coords[0] = b >> 4 & 0x3
    v.coords[1] = b >> 2 & 0x3
    v.coords[2] = b      & 0x3
}

func (v Vector3) Get () (uint, uint, uint) {
    return v.coords[0], v.coords[1], v.coords[2]
}

func (v Vector3) Get8 () (uint8, uint8, uint8) {
    return uint8(v.coords[0]), uint8(v.coords[1]), uint8(v.coords[2])
}

// keep two bits of each components to obtain a byte
func (v Vector3) GetByte () uint8 {
    x := v.coords[0] & 0x3
    y := v.coords[1] & 0x3
    z := v.coords[2] & 0x3
    return uint8(x << 4 | y << 2 | z)
}

/**/

type Byte3 struct {
    coords [3]uint8
}

func (v Byte3) Set (x, y, z uint) {
    v.coords[0] = uint8(x)
    v.coords[1] = uint8(y)
    v.coords[2] = uint8(z)
}

func (v Byte3) Set8 (x, y, z uint8) {
    v.coords[0] = x
    v.coords[1] = y
    v.coords[2] = z
}

// convert a single byte into three components
func (v Byte3) SetByte (byte uint8) {
    v.coords[0] = byte >> 4 & 0x3
    v.coords[1] = byte >> 2 & 0x3
    v.coords[2] = byte      & 0x3
}

func (v Byte3) Get () (uint, uint, uint) {
    return uint(v.coords[0]), uint(v.coords[1]), uint(v.coords[2])
}

func (v Byte3) Get8 () (uint8, uint8, uint8) {
    return v.coords[0], v.coords[1], v.coords[2]
}

// keep two bits of each components to obtain a byte
func (v Byte3) GetByte () uint8 {
    x := v.coords[0] & 0x3
    y := v.coords[1] & 0x3
    z := v.coords[2] & 0x3
    return x << 4 | y << 2 | z
}

/**/

type Bool3 struct {
    bools [3]bool
}

func (v Bool3) SetByte (byte uint8) {
    v.bools[0] = (byte & 0x4) != 0
    v.bools[1] = (byte & 0x2) != 0
    v.bools[2] = (byte & 0x1) != 0
}

func (v Bool3) GetByte () uint8 {
    b := 0
    if v.bools[0] {b |= 0x4}
    if v.bools[1] {b |= 0x2}
    if v.bools[2] {b |= 0x1}
    return uint8(b)
}
