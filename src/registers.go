package main

type Registers int
type Cond_flags int

const (
	R_R0 Registers = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC
	R_COND
	R_COUNT
)

const (
	FL_POS Cond_flags = 1 << iota
	FL_ZRO
	FL_NEG
)

const PC_START = 0x3000
