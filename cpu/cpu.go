package cpu

type CPU struct {
	reg *register
}

func New() *CPU {
	cpu := &CPU{
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

func (cpu CPU) getAF() uint16 {
	return uint16(cpu.reg.a)<<8 | uint16(cpu.reg.f)
}

func (cpu CPU) getBC() uint16 {
	return uint16(cpu.reg.b)<<8 | uint16(cpu.reg.c)
}

func (cpu CPU) getDE() uint16 {
	return uint16(cpu.reg.d)<<8 | uint16(cpu.reg.e)
}

func (cpu CPU) getHL() uint16 {
	return uint16(cpu.reg.h)<<8 | uint16(cpu.reg.l)
}

func (cpu *CPU) setAF(val uint16) {
	cpu.reg.a = uint8(0xFF & (val >> 8))
	cpu.reg.f = uint8(0xFF & val)
}

func (cpu *CPU) setBC(val uint16) {
	cpu.reg.b = uint8(0xFF & (val >> 8))
	cpu.reg.c = uint8(0xFF & val)
}

func (cpu *CPU) setDE(val uint16) {
	cpu.reg.d = uint8(0xFF & (val >> 8))
	cpu.reg.e = uint8(0xFF & val)
}

func (cpu *CPU) setHL(val uint16) {
	cpu.reg.h = uint8(0xFF & (val >> 8))
	cpu.reg.l = uint8(0xFF & val)
}

func (cpu *CPU) setZeroFlag() {
	cpu.reg.f |= 0x80
}

func (cpu *CPU) setSubFlag() {
	cpu.reg.f |= 0x40
}

func (cpu *CPU) setHalfCarryFlag() {
	cpu.reg.f |= 0x20
}

func (cpu *CPU) setCarryFlag() {
	cpu.reg.f |= 0x10
}

func (cpu *CPU) resetZeroFlag() {
	cpu.reg.f &= 0x7F
}

func (cpu *CPU) resetSubFlag() {
	cpu.reg.f &= 0xBF
}

func (cpu *CPU) resetHalfCarryFlag() {
	cpu.reg.f &= 0xDF
}

func (cpu *CPU) resetCarryFlag() {
	cpu.reg.f &= 0xEF
}

func (cpu CPU) getZeroFlag() bool {
	return ((cpu.reg.f & 0x80) >> 7) == 1
}

func (cpu CPU) getSubFlag() bool {
	return ((cpu.reg.f & 0x40) >> 6) == 1
}

func (cpu CPU) getHalfCarryFlag() bool {
	return ((cpu.reg.f & 0x20) >> 5) == 1
}

func (cpu CPU) getCarryFlag() bool {
	return ((cpu.reg.f & 0x10) >> 4) == 1
}
