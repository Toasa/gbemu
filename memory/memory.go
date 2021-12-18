package memory

type Memory struct {
	mem [0x10000]uint8
}

func New() *Memory {
	return new(Memory)
}

func (m Memory) ReadByte(addr uint16) uint8 {
	return m.mem[addr]
}
