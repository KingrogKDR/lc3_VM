package emulator

func SignExtend(x uint16, bitCount int) uint16 {
	if (x >> (uint16(bitCount) - 1) & 1) == 1 {
		mask := ^((1 << bitCount) - 1)
		x |= uint16(mask)
	}
	return x
}

func (v *VM) add(instr uint16) {
	r0 := (instr >> 9) & 0x7       // DR (3-bit)
	r1 := (instr >> 6) & 0x7       // SR1 (3-bit)
	imm_flag := (instr >> 5) & 0x1 // immediate flag (1-bit)

	if imm_flag == 1 {
		imm_flag = SignExtend(instr&0x1F, 5)
		v.Reg[r0] = v.Reg[r1] + imm_flag
	} else {
		r2 := instr & 0x7 // SR2 (3-bit)
		v.Reg[r0] = v.Reg[r1] + v.Reg[r2]
	}

	v.UpdateFlags(r0)

}
