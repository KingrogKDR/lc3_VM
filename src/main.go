package main

const MAX_MEMORY = 1 << 16

var is_Running bool

type VM struct {
	memory [MAX_MEMORY]uint16
	reg    [R_COUNT]uint16
}

func (v VM) mem_read(address uint16) uint16 {
	return v.memory[address]
}

func main() {
	vm := VM{}
	vm.reg[R_COND] = uint16(FL_ZRO)

	vm.reg[R_PC] = PC_START
	is_Running = true

	for is_Running {
		instr := vm.mem_read(vm.reg[R_PC])
		vm.reg[R_PC]++
		op := instr >> 12

		switch op {
		case uint16(OP_ADD):
			//add

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
}
