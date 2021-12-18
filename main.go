package main

import (
	"fmt"
	"gbemu/cart"
	"gbemu/cpu"
)

func main() {
	_ = cpu.CPU{}

	romPath := "rom/gb-test-roms/cpu_instrs/cpu_instrs.gb"
	cart.ReadCartridge(romPath)
	fmt.Println("OK")
}
