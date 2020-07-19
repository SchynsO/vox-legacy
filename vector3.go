package main


// represent a data structure with 3 components
type Vector3 struct {
    x, y, z uint
}

func (v Vector3) Set (x, y, z uint) {
    v.x = x
    v.y = y
    v.z = z
}

func (v Vector3) Set8 (x, y, z uint8) {
    v.x = uint(x)
    v.y = uint(y)
    v.z = uint(z)
}

// convert a single byte into three components
func (v Vector3) SetByte (byte uint8) {
    b := uint(byte)
    v.x = b >> 4 & 0x3
    v.y = b >> 2 & 0x3
    v.z = b      & 0x3
}

func (v Vector3) Get () (uint, uint, uint) {
    return v.x, v.y, v.z
}

func (v Vector3) Get8 () (uint8, uint8, uint8) {
    return uint8(v.x), uint8(v.y), uint8(v.z)
}

// keep two bits of each components to obtain a byte
func (v Vector3) GetByte () uint8 {
    x := v.x & 0x3
    y := v.y & 0x3
    z := v.z & 0x3
    return uint8(x << 4 | y << 2 | z)
}


func (v1 Vector3) Add (v2 Vector3) Vector3 {
    var v3 Vector3
    v3.x = v1.x + v2.x
    v3.y = v1.y + v2.y
    v3.z = v1.z + v2.z
    return v3
}

func (v1 Vector3) Sub (v2 Vector3) Vector3 {
    var v3 Vector3
    v3.x = v1.x - v2.x
    v3.y = v1.y - v2.y
    v3.z = v1.z - v2.z
    return v3
}

func (v Vector3) Mod (m uint) Vector3 {
    var w Vector3
    w.x = v.x % m
    w.y = v.y % m
    w.z = v.z % m
    return w
}

func (v Vector3) Mask (m uint) Vector3 {
    var w Vector3
    w.x = v.x & m
    w.y = v.y & m
    w.z = v.z & m
    return w
}

func (v Vector3) ShiftL (b uint) Vector3 {
    var w Vector3
    w.x = v.x << b
    w.y = v.y << b
    w.z = v.z << b
    return w
}

func (v Vector3) ShiftR (b uint) Vector3 {
    var w Vector3
    w.x = v.x >> b
    w.y = v.y >> b
    w.z = v.z >> b
    return w
}


/**/

type Byte3 struct {
    x, y, z uint8
}

func (v Byte3) Set (x, y, z uint) {
    v.x = uint8(x)
    v.y = uint8(y)
    v.z = uint8(z)
}

func (v Byte3) Set8 (x, y, z uint8) {
    v.x = x
    v.y = y
    v.z = z
}

// convert a single byte into three components
func (v Byte3) SetByte (byte uint8) {
    v.x = byte >> 4 & 0x3
    v.y = byte >> 2 & 0x3
    v.z = byte      & 0x3
}

func (v Byte3) Get () (uint, uint, uint) {
    return uint(v.x), uint(v.y), uint(v.z)
}

func (v Byte3) Get8 () (uint8, uint8, uint8) {
    return v.x, v.y, v.z
}

// keep two bits of each components to obtain a byte
func (v Byte3) GetByte () uint8 {
    x := v.x & 0x3
    y := v.y & 0x3
    z := v.z & 0x3
    return x << 4 | y << 2 | z
}

/**/

type Bool3 struct {
    x, y, z bool
}

func (v Bool3) SetByte (byte uint8) {
    v.x = (byte & 0x4) != 0
    v.y = (byte & 0x2) != 0
    v.z = (byte & 0x1) != 0
}

func (v Bool3) GetByte () uint8 {
    b := 0
    if v.x {b |= 0x4}
    if v.y {b |= 0x2}
    if v.z {b |= 0x1}
    return uint8(b)
}
