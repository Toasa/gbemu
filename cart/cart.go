package cart

import (
	"fmt"
	"gbemu/log"
	"io/ioutil"
)

func ReadCartridge(filepath string) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	title := string(f[0x0134:0x0143])
	cartType := cartTypes[f[0x0147]]
	romSize := romSizes[f[0x0148]]
	ramSize := ramSizes[f[0x0149]]

	fmt.Println("title:", title)
	fmt.Println("cartType:", cartType)
	fmt.Println("romSize:", romSize.size)
	fmt.Println("ramSize:", ramSize)

	var x byte = 0
	for i := 0x0134; i <= 0x014C; i++ {
		x = x - f[i] - 1
	}
	if x != f[0x014D] {
		log.Fatal("Checksum calculation failed")
	}
	fmt.Println("Checksum OK")
}

var cartTypes = map[byte]string{
	0x00: "ROM ONLY",
	0x01: "MBC1",
	0x02: "MBC1+RAM",
	0x03: "MBC1+RAM+BATTERY",
	0x05: "MBC2",
	0x06: "MBC2+BATTERY",
	0x08: "ROM+RAM 1",
	0x09: "ROM+RAM+BATTERY 1",
	0x0B: "MMM01",
	0x0C: "MMM01+RAM",
	0x0D: "MMM01+RAM+BATTERY",
	0x0F: "MBC3+TIMER+BATTERY",
	0x10: "MBC3+TIMER+RAM+BATTERY 2",
	0x11: "MBC3",
	0x12: "MBC3+RAM 2",
	0x13: "MBC3+RAM+BATTERY 2",
	0x19: "MBC5",
	0x1A: "MBC5+RAM",
	0x1B: "MBC5+RAM+BATTERY",
	0x1C: "MBC5+RUMBLE",
	0x1D: "MBC5+RUMBLE+RAM",
	0x1E: "MBC5+RUMBLE+RAM+BATTERY",
	0x20: "MBC6",
	0x22: "MBC7+SENSOR+RUMBLE+RAM+BATTERY",
	0xFC: "POCKET CAMERA",
	0xFD: "BANDAI TAMA5",
	0xFE: "HuC3",
	0xFF: "HuC1+RAM+BATTERY",
}

var romSizes = map[byte]struct {
	size   string
	nBanks int
}{
	0x00: {"32 KByte", 2},
	0x01: {"64 KByte", 4},
	0x02: {"128 KByte", 8},
	0x03: {"256 KByte", 16},
	0x04: {"512 KByte", 32},
	0x05: {"1 MByte", 64},
	0x06: {"2 MByte", 128},
	0x07: {"4 MByte", 256},
	0x08: {"8 MByte", 512},
	0x52: {"1.1 MByte", 72},
	0x53: {"1.2 MByte", 80},
	0x54: {"1.5 MByte", 96},
}

var ramSizes = map[byte]string{
	0x00: "0",
	0x01: "-",
	0x02: "8 KB",
	0x03: "32 KB",
	0x04: "128 KB",
	0x05: "64 KB",
}
