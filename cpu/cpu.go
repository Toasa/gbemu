package cpu

type Register struct {
    a uint8
    b uint8
    c uint8
    d uint8
    e uint8

    // The flags register. Upper 4 bits show the result of the most recent
    // instruction. Lower 4 bits are always zero. The Upper 4bits are as
    // follows:
    //    Bit    Explanation
    // ---------------------------
    //    7      Zero flag
    //    6      Subtraction flag
    //    5      Half carry flag
    //    4      Carry flag
    // ---------------------------
    f uint8

    h uint8
    l uint8
}

func (r Register) getAF() uint16 {
    return uint16(r.a) << 8 | uint16(r.f)
}

func (r Register) getBC() uint16 {
    return uint16(r.b) << 8 | uint16(r.c)
}

func (r Register) getDE() uint16 {
    return uint16(r.d) << 8 | uint16(r.e)
}

func (r Register) getHL() uint16 {
    return uint16(r.h) << 8 | uint16(r.l)
}

func (r *Register) setAF(val uint16) {
    r.a = uint8(0xFF & (val >> 8))
    r.f = uint8(0xFF & val)
}

func (r *Register) setBC(val uint16) {
    r.b = uint8(0xFF & (val >> 8))
    r.c = uint8(0xFF & val)
}

func (r *Register) setDE(val uint16) {
    r.d = uint8(0xFF & (val >> 8))
    r.e = uint8(0xFF & val)
}

func (r *Register) setHL(val uint16) {
    r.h = uint8(0xFF & (val >> 8))
    r.l = uint8(0xFF & val)
}

func (r *Register) setZeroFlag() {
    r.f |= 0x80
}

func (r *Register) setSubFlag() {
    r.f |= 0x40
}

func (r *Register) setHalfCarryFlag() {
    r.f |= 0x20
}

func (r *Register) setCarryFlag() {
    r.f |= 0x10
}

func (r Register) getZeroFlag() bool {
    return ((r.f & 0x80) >> 7) == 1
}

func (r Register) getSubFlag() bool {
    return ((r.f & 0x40) >> 6) == 1
}

func (r Register) getHalfCarryFlag() bool {
    return ((r.f & 0x20) >> 5) == 1
}

func (r Register) getCarryFlag() bool {
    return ((r.f & 0x10) >> 4) == 1
}
