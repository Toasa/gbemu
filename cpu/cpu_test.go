package cpu

import (
    "testing"
)

func TestGetRegister (t *testing.T) {
    r := register {
        a: 0x11, b: 0x22, c: 0x33, d: 0x44,
        e: 0x55, f: 0x66, h: 0x77, l: 0x88,
    }

    if r.getAF() != 0x1166 {
        t.Errorf("getAF() failed")
    }
    if r.getBC() != 0x2233 {
        t.Errorf("getBC() failed")
    }
    if r.getDE() != 0x4455 {
        t.Errorf("getDE() failed")
    }
    if r.getHL() != 0x7788 {
        t.Errorf("getHL() failed")
    }
}

func TestSetRegister (t *testing.T) {
    r := register {}

    r.setAF(0x1166)
    r.setBC(0x2233)
    r.setDE(0x4455)
    r.setHL(0x7788)

    if r.a != 0x11 || r.f != 0x66 {
        t.Errorf("setAF() failed")
    }
    if r.b != 0x22 || r.c != 0x33 {
        t.Errorf("setBC() failed")
    }
    if r.d != 0x44 || r.e != 0x55 {
        t.Errorf("setDE() failed")
    }
    if r.h != 0x77 || r.l != 0x88 {
        t.Errorf("setHL() failed")
    }
}

func TestZeroFlagRegister (t *testing.T) {
    r := register {}

    r.setZeroFlag()
    if r.f != 0x80 {
        t.Errorf("setZeroFlag() failed")
    }
    if !r.getZeroFlag() {
        t.Errorf("getZeroFlag() failed")
    }
}

func TestSubFlagRegister (t *testing.T) {
    r := register {}

    r.setSubFlag()
    if r.f != 0x40 {
        t.Errorf("setSubFlag() failed")
    }
    if !r.getSubFlag() {
        t.Errorf("getSubFlag() failed")
    }
}

func TestHalfCarryFlagRegister (t *testing.T) {
    r := register {}

    r.setHalfCarryFlag()
    if r.f != 0x20 {
        t.Errorf("setHalfCarryFlag() failed")
    }
    if !r.getHalfCarryFlag() {
        t.Errorf("getHalfCarryFlag() failed")
    }
}

func TestCarryFlagRegister (t *testing.T) {
    r := register {}

    r.setCarryFlag()
    if r.f != 0x10 {
        t.Errorf("setCarryFlag() failed")
    }
    if !r.getCarryFlag() {
        t.Errorf("getCarryFlag() failed")
    }
}

