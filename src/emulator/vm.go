package emulator

type VM struct {
	Memory [MAX_MEMORY]uint16
	Reg    [R_COUNT]uint16
}

func NewVM() *VM {
	return &VM{} // Return a pointer to a new VM struct
}

func (v *VM) Mem_read(address uint16) uint16 {
	return v.Memory[address]
}

func (v *VM) UpdateFlags(regIndex uint16) {
	if v.Reg[regIndex] == 0 {
		v.Reg[R_COND] = uint16(FL_ZRO)
	} else if (v.Reg[regIndex] >> 15) == 1 {
		v.Reg[R_COND] = uint16(FL_NEG)
	} else {
		v.Reg[R_COND] = uint16(FL_POS)
	}
}

func (v *VM) CheckInstr(opcode uint16, instruction uint16) {
	switch opcode {
	case uint16(OP_ADD):
		v.add(instruction)
	case uint16(OP_AND):

	case uint16(OP_BR):

	case uint16(OP_JMP):

	case uint16(OP_JSR):

	case uint16(OP_LD):

	case uint16(OP_LDI):

	case uint16(OP_LDR):

	case uint16(OP_LEA):

	case uint16(OP_NOT):

	case uint16(OP_ST):

	case uint16(OP_STI):

	case uint16(OP_STR):

	case uint16(OP_TRAP):

	case uint16(OP_RES):
	case uint16(OP_RTI):
	default:
		//bad opcode
	}
}
