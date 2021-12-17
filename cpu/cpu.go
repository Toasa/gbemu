package cpu

type CPU struct {
	reg *register
}

func New() *CPU {
    cpu := &CPU {
        reg: &register{},
    }
    return cpu
}

type register struct {
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

func (r register) getAF() uint16 {
	return uint16(r.a)<<8 | uint16(r.f)
}

func (r register) getBC() uint16 {
	return uint16(r.b)<<8 | uint16(r.c)
}

func (r register) getDE() uint16 {
	return uint16(r.d)<<8 | uint16(r.e)
}

func (r register) getHL() uint16 {
	return uint16(r.h)<<8 | uint16(r.l)
}

func (r *register) setAF(val uint16) {
	r.a = uint8(0xFF & (val >> 8))
	r.f = uint8(0xFF & val)
}

func (r *register) setBC(val uint16) {
	r.b = uint8(0xFF & (val >> 8))
	r.c = uint8(0xFF & val)
}

func (r *register) setDE(val uint16) {
	r.d = uint8(0xFF & (val >> 8))
	r.e = uint8(0xFF & val)
}

func (r *register) setHL(val uint16) {
	r.h = uint8(0xFF & (val >> 8))
	r.l = uint8(0xFF & val)
}

func (r *register) setZeroFlag() {
	r.f |= 0x80
}

func (r *register) setSubFlag() {
	r.f |= 0x40
}

func (r *register) setHalfCarryFlag() {
	r.f |= 0x20
}

func (r *register) setCarryFlag() {
	r.f |= 0x10
}

func (r *register) resetZeroFlag() {
	r.f &= 0x7F
}

func (r *register) resetSubFlag() {
	r.f &= 0xBF
}

func (r *register) resetHalfCarryFlag() {
	r.f &= 0xDF
}

func (r *register) resetCarryFlag() {
	r.f &= 0xEF
}

func (r register) getZeroFlag() bool {
	return ((r.f & 0x80) >> 7) == 1
}

func (r register) getSubFlag() bool {
	return ((r.f & 0x40) >> 6) == 1
}

func (r register) getHalfCarryFlag() bool {
	return ((r.f & 0x20) >> 5) == 1
}

func (r register) getCarryFlag() bool {
	return ((r.f & 0x10) >> 4) == 1
}
