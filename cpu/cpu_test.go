package cpu

import (
    "testing"
)

func TestGetRegister (t *testing.T) {
    cpu := New()
    cpu.reg = &register {
        a: 0x11, b: 0x22, c: 0x33, d: 0x44,
        e: 0x55, f: 0x66, h: 0x77, l: 0x88,
    }

    if cpu.getAF() != 0x1166 {
        t.Errorf("getAF() failed")
    }
    if cpu.getBC() != 0x2233 {
        t.Errorf("getBC() failed")
    }
    if cpu.getDE() != 0x4455 {
        t.Errorf("getDE() failed")
    }
    if cpu.getHL() != 0x7788 {
        t.Errorf("getHL() failed")
    }
}

func TestSetRegister (t *testing.T) {
    cpu := New()

    cpu.setAF(0x1166)
    cpu.setBC(0x2233)
    cpu.setDE(0x4455)
    cpu.setHL(0x7788)

    if cpu.reg.a != 0x11 || cpu.reg.f != 0x66 {
        t.Errorf("setAF() failed")
    }
    if cpu.reg.b != 0x22 || cpu.reg.c != 0x33 {
        t.Errorf("setBC() failed")
    }
    if cpu.reg.d != 0x44 || cpu.reg.e != 0x55 {
        t.Errorf("setDE() failed")
    }
    if cpu.reg.h != 0x77 || cpu.reg.l != 0x88 {
        t.Errorf("setHL() failed")
    }
}

func TestZeroFlagRegister (t *testing.T) {
    cpu := New()

    cpu.setZeroFlag()
    if cpu.reg.f != 0x80 {
        t.Errorf("setZeroFlag() failed")
    }
    if !cpu.getZeroFlag() {
        t.Errorf("getZeroFlag() failed")
    }
}

func TestSubFlagRegister (t *testing.T) {
    cpu := New()

    cpu.setSubFlag()
    if cpu.reg.f != 0x40 {
        t.Errorf("setSubFlag() failed")
    }
    if !cpu.getSubFlag() {
        t.Errorf("getSubFlag() failed")
    }
}

func TestHalfCarryFlagRegister (t *testing.T) {
    cpu := New()

    cpu.setHalfCarryFlag()
    if cpu.reg.f != 0x20 {
        t.Errorf("setHalfCarryFlag() failed")
    }
    if !cpu.getHalfCarryFlag() {
        t.Errorf("getHalfCarryFlag() failed")
    }
}

func TestCarryFlagRegister (t *testing.T) {
    cpu := New()

    cpu.setCarryFlag()
    if cpu.reg.f != 0x10 {
        t.Errorf("setCarryFlag() failed")
    }
    if !cpu.getCarryFlag() {
        t.Errorf("getCarryFlag() failed")
    }
}

