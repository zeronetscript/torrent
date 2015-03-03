// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"cmd/internal/obj"
	"cmd/internal/obj/i386"
)
import "cmd/internal/gc"

var (
	AX               = RtoB(i386.REG_AX)
	BX               = RtoB(i386.REG_BX)
	CX               = RtoB(i386.REG_CX)
	DX               = RtoB(i386.REG_DX)
	DI               = RtoB(i386.REG_DI)
	SI               = RtoB(i386.REG_SI)
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [i386.ALAST]gc.ProgInfo{
	obj.ATYPE:     gc.ProgInfo{gc.Pseudo | gc.Skip, 0, 0, 0},
	obj.ATEXT:     gc.ProgInfo{gc.Pseudo, 0, 0, 0},
	obj.AFUNCDATA: gc.ProgInfo{gc.Pseudo, 0, 0, 0},
	obj.APCDATA:   gc.ProgInfo{gc.Pseudo, 0, 0, 0},
	obj.AUNDEF:    gc.ProgInfo{gc.Break, 0, 0, 0},
	obj.AUSEFIELD: gc.ProgInfo{gc.OK, 0, 0, 0},
	obj.ACHECKNIL: gc.ProgInfo{gc.LeftRead, 0, 0, 0},
	obj.AVARDEF:   gc.ProgInfo{gc.Pseudo | gc.RightWrite, 0, 0, 0},
	obj.AVARKILL:  gc.ProgInfo{gc.Pseudo | gc.RightWrite, 0, 0, 0},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Intel opcode.
	obj.ANOP:        gc.ProgInfo{gc.LeftRead | gc.RightWrite, 0, 0, 0},
	i386.AADCL:      gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.AADCW:      gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.AADDB:      gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AADDL:      gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AADDW:      gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AADDSD:     gc.ProgInfo{gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.AADDSS:     gc.ProgInfo{gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.AANDB:      gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AANDL:      gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AANDW:      gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	obj.ACALL:       gc.ProgInfo{gc.RightAddr | gc.Call | gc.KillCarry, 0, 0, 0},
	i386.ACDQ:       gc.ProgInfo{gc.OK, AX, AX | DX, 0},
	i386.ACWD:       gc.ProgInfo{gc.OK, AX, AX | DX, 0},
	i386.ACLD:       gc.ProgInfo{gc.OK, 0, 0, 0},
	i386.ASTD:       gc.ProgInfo{gc.OK, 0, 0, 0},
	i386.ACMPB:      gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ACMPL:      gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ACMPW:      gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ACOMISD:    gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ACOMISS:    gc.ProgInfo{gc.SizeF | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ACVTSD2SL:  gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTSD2SS:  gc.ProgInfo{gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTSL2SD:  gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTSL2SS:  gc.ProgInfo{gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTSS2SD:  gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTSS2SL:  gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTTSD2SL: gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ACVTTSS2SL: gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.ADECB:      gc.ProgInfo{gc.SizeB | RightRdwr, 0, 0, 0},
	i386.ADECL:      gc.ProgInfo{gc.SizeL | RightRdwr, 0, 0, 0},
	i386.ADECW:      gc.ProgInfo{gc.SizeW | RightRdwr, 0, 0, 0},
	i386.ADIVB:      gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	i386.ADIVL:      gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	i386.ADIVW:      gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	i386.ADIVSD:     gc.ProgInfo{gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.ADIVSS:     gc.ProgInfo{gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.AFLDCW:     gc.ProgInfo{gc.SizeW | gc.LeftAddr, 0, 0, 0},
	i386.AFSTCW:     gc.ProgInfo{gc.SizeW | gc.RightAddr, 0, 0, 0},
	i386.AFSTSW:     gc.ProgInfo{gc.SizeW | gc.RightAddr | gc.RightWrite, 0, 0, 0},
	i386.AFADDD:     gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFADDDP:    gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFADDF:     gc.ProgInfo{gc.SizeF | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFCOMD:     gc.ProgInfo{gc.SizeD | gc.LeftAddr | gc.RightRead, 0, 0, 0},
	i386.AFCOMDP:    gc.ProgInfo{gc.SizeD | gc.LeftAddr | gc.RightRead, 0, 0, 0},
	i386.AFCOMDPP:   gc.ProgInfo{gc.SizeD | gc.LeftAddr | gc.RightRead, 0, 0, 0},
	i386.AFCOMF:     gc.ProgInfo{gc.SizeF | gc.LeftAddr | gc.RightRead, 0, 0, 0},
	i386.AFCOMFP:    gc.ProgInfo{gc.SizeF | gc.LeftAddr | gc.RightRead, 0, 0, 0},
	i386.AFUCOMIP:   gc.ProgInfo{gc.SizeF | gc.LeftAddr | gc.RightRead, 0, 0, 0},
	i386.AFCHS:      gc.ProgInfo{gc.SizeD | RightRdwr, 0, 0, 0}, // also SizeF

	i386.AFDIVDP:  gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFDIVF:   gc.ProgInfo{gc.SizeF | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFDIVD:   gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFDIVRDP: gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFDIVRF:  gc.ProgInfo{gc.SizeF | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFDIVRD:  gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFXCHD:   gc.ProgInfo{gc.SizeD | LeftRdwr | RightRdwr, 0, 0, 0},
	i386.AFSUBD:   gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFSUBDP:  gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFSUBF:   gc.ProgInfo{gc.SizeF | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFSUBRD:  gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFSUBRDP: gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFSUBRF:  gc.ProgInfo{gc.SizeF | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFMOVD:   gc.ProgInfo{gc.SizeD | gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	i386.AFMOVF:   gc.ProgInfo{gc.SizeF | gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	i386.AFMOVL:   gc.ProgInfo{gc.SizeL | gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	i386.AFMOVW:   gc.ProgInfo{gc.SizeW | gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	i386.AFMOVV:   gc.ProgInfo{gc.SizeQ | gc.LeftAddr | gc.RightWrite, 0, 0, 0},

	// These instructions are marked as RightAddr
	// so that the register optimizer does not try to replace the
	// memory references with integer register references.
	// But they do not use the previous value at the address, so
	// we also mark them RightWrite.
	i386.AFMOVDP:  gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightWrite | gc.RightAddr, 0, 0, 0},
	i386.AFMOVFP:  gc.ProgInfo{gc.SizeF | gc.LeftRead | gc.RightWrite | gc.RightAddr, 0, 0, 0},
	i386.AFMOVLP:  gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.RightAddr, 0, 0, 0},
	i386.AFMOVWP:  gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.RightWrite | gc.RightAddr, 0, 0, 0},
	i386.AFMOVVP:  gc.ProgInfo{gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.RightAddr, 0, 0, 0},
	i386.AFMULD:   gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFMULDP:  gc.ProgInfo{gc.SizeD | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AFMULF:   gc.ProgInfo{gc.SizeF | gc.LeftAddr | RightRdwr, 0, 0, 0},
	i386.AIDIVB:   gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	i386.AIDIVL:   gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	i386.AIDIVW:   gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.SetCarry, AX | DX, AX | DX, 0},
	i386.AIMULB:   gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	i386.AIMULL:   gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.ImulAXDX | gc.SetCarry, 0, 0, 0},
	i386.AIMULW:   gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.ImulAXDX | gc.SetCarry, 0, 0, 0},
	i386.AINCB:    gc.ProgInfo{gc.SizeB | RightRdwr, 0, 0, 0},
	i386.AINCL:    gc.ProgInfo{gc.SizeL | RightRdwr, 0, 0, 0},
	i386.AINCW:    gc.ProgInfo{gc.SizeW | RightRdwr, 0, 0, 0},
	i386.AJCC:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJCS:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJEQ:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJGE:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJGT:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJHI:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJLE:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJLS:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJLT:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJMI:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJNE:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJOC:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJOS:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJPC:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJPL:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	i386.AJPS:     gc.ProgInfo{gc.Cjmp | gc.UseCarry, 0, 0, 0},
	obj.AJMP:      gc.ProgInfo{gc.Jump | gc.Break | gc.KillCarry, 0, 0, 0},
	i386.ALEAL:    gc.ProgInfo{gc.LeftAddr | gc.RightWrite, 0, 0, 0},
	i386.AMOVBLSX: gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.AMOVBLZX: gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.AMOVBWSX: gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.AMOVBWZX: gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.AMOVWLSX: gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.AMOVWLZX: gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Conv, 0, 0, 0},
	i386.AMOVB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	i386.AMOVL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	i386.AMOVW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	i386.AMOVSB:   gc.ProgInfo{gc.OK, DI | SI, DI | SI, 0},
	i386.AMOVSL:   gc.ProgInfo{gc.OK, DI | SI, DI | SI, 0},
	i386.AMOVSW:   gc.ProgInfo{gc.OK, DI | SI, DI | SI, 0},
	obj.ADUFFCOPY: gc.ProgInfo{gc.OK, DI | SI, DI | SI | CX, 0},
	i386.AMOVSD:   gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	i386.AMOVSS:   gc.ProgInfo{gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},

	// We use MOVAPD as a faster synonym for MOVSD.
	i386.AMOVAPD:  gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move, 0, 0, 0},
	i386.AMULB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.SetCarry, AX, AX, 0},
	i386.AMULL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.SetCarry, AX, AX | DX, 0},
	i386.AMULW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.SetCarry, AX, AX | DX, 0},
	i386.AMULSD:   gc.ProgInfo{gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.AMULSS:   gc.ProgInfo{gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.ANEGB:    gc.ProgInfo{gc.SizeB | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.ANEGL:    gc.ProgInfo{gc.SizeL | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.ANEGW:    gc.ProgInfo{gc.SizeW | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.ANOTB:    gc.ProgInfo{gc.SizeB | RightRdwr, 0, 0, 0},
	i386.ANOTL:    gc.ProgInfo{gc.SizeL | RightRdwr, 0, 0, 0},
	i386.ANOTW:    gc.ProgInfo{gc.SizeW | RightRdwr, 0, 0, 0},
	i386.AORB:     gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AORL:     gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AORW:     gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.APOPL:    gc.ProgInfo{gc.SizeL | gc.RightWrite, 0, 0, 0},
	i386.APUSHL:   gc.ProgInfo{gc.SizeL | gc.LeftRead, 0, 0, 0},
	i386.ARCLB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ARCLL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ARCLW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ARCRB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ARCRL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ARCRW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.AREP:     gc.ProgInfo{gc.OK, CX, CX, 0},
	i386.AREPN:    gc.ProgInfo{gc.OK, CX, CX, 0},
	obj.ARET:      gc.ProgInfo{gc.Break | gc.KillCarry, 0, 0, 0},
	i386.AROLB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.AROLL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.AROLW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ARORB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ARORL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ARORW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASAHF:    gc.ProgInfo{gc.OK, AX, AX, 0},
	i386.ASALB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASALL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASALW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASARB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASARL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASARW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASBBB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ASBBL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ASBBW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry | gc.UseCarry, 0, 0, 0},
	i386.ASETCC:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETCS:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETEQ:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETGE:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETGT:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETHI:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETLE:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETLS:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETLT:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETMI:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETNE:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETOC:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETOS:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETPC:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETPL:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASETPS:   gc.ProgInfo{gc.SizeB | RightRdwr | gc.UseCarry, 0, 0, 0},
	i386.ASHLB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASHLL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASHLW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASHRB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASHRL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASHRW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.ShiftCX | gc.SetCarry, 0, 0, 0},
	i386.ASTOSB:   gc.ProgInfo{gc.OK, AX | DI, DI, 0},
	i386.ASTOSL:   gc.ProgInfo{gc.OK, AX | DI, DI, 0},
	i386.ASTOSW:   gc.ProgInfo{gc.OK, AX | DI, DI, 0},
	obj.ADUFFZERO: gc.ProgInfo{gc.OK, AX | DI, DI, 0},
	i386.ASUBB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.ASUBL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.ASUBW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.ASUBSD:   gc.ProgInfo{gc.SizeD | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.ASUBSS:   gc.ProgInfo{gc.SizeF | gc.LeftRead | RightRdwr, 0, 0, 0},
	i386.ATESTB:   gc.ProgInfo{gc.SizeB | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ATESTL:   gc.ProgInfo{gc.SizeL | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.ATESTW:   gc.ProgInfo{gc.SizeW | gc.LeftRead | gc.RightRead | gc.SetCarry, 0, 0, 0},
	i386.AUCOMISD: gc.ProgInfo{gc.SizeD | gc.LeftRead | gc.RightRead, 0, 0, 0},
	i386.AUCOMISS: gc.ProgInfo{gc.SizeF | gc.LeftRead | gc.RightRead, 0, 0, 0},
	i386.AXCHGB:   gc.ProgInfo{gc.SizeB | LeftRdwr | RightRdwr, 0, 0, 0},
	i386.AXCHGL:   gc.ProgInfo{gc.SizeL | LeftRdwr | RightRdwr, 0, 0, 0},
	i386.AXCHGW:   gc.ProgInfo{gc.SizeW | LeftRdwr | RightRdwr, 0, 0, 0},
	i386.AXORB:    gc.ProgInfo{gc.SizeB | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AXORL:    gc.ProgInfo{gc.SizeL | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
	i386.AXORW:    gc.ProgInfo{gc.SizeW | gc.LeftRead | RightRdwr | gc.SetCarry, 0, 0, 0},
}

func proginfo(p *obj.Prog) (info gc.ProgInfo) {
	info = progtable[p.As]
	if info.Flags == 0 {
		gc.Fatal("unknown instruction %v", p)
	}

	if (info.Flags&gc.ShiftCX != 0) && p.From.Type != obj.TYPE_CONST {
		info.Reguse |= CX
	}

	if info.Flags&gc.ImulAXDX != 0 {
		if p.To.Type == obj.TYPE_NONE {
			info.Reguse |= AX
			info.Regset |= AX | DX
		} else {
			info.Flags |= RightRdwr
		}
	}

	// Addressing makes some registers used.
	if p.From.Type == obj.TYPE_MEM && p.From.Name == obj.NAME_NONE {
		info.Regindex |= RtoB(int(p.From.Reg))
	}
	if p.From.Index != i386.REG_NONE {
		info.Regindex |= RtoB(int(p.From.Index))
	}
	if p.To.Type == obj.TYPE_MEM && p.To.Name == obj.NAME_NONE {
		info.Regindex |= RtoB(int(p.To.Reg))
	}
	if p.To.Index != i386.REG_NONE {
		info.Regindex |= RtoB(int(p.To.Index))
	}

	return info
}