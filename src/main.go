package main

import (
	"fmt"
	"os"

	"github.com/KingrogKDR/lc3_VM/src/emulator"
)

const PC_START = 0x3000

var is_Running bool

func main() {

	if len(os.Args) < 2 {
		fmt.Println("lc3 Image failed....")
		os.Exit(2)
	}

	for i := 1; i < len(os.Args); i++ {
		imagePath := os.Args[i]
		if !ReadImage(imagePath) {
			fmt.Printf("failed to load image: %s\n", imagePath)
			os.Exit(1)
		}
	}

	fmt.Println("All images loaded successfully. Starting LC-3 emulation...")

	myVm := emulator.NewVM()

	myVm.Reg[emulator.R_COND] = uint16(emulator.FL_ZRO)

	myVm.Reg[emulator.R_PC] = PC_START
	is_Running = true

	for is_Running {
		instr := myVm.Mem_read(myVm.Reg[emulator.R_PC])
		myVm.Reg[emulator.R_PC]++
		op := instr >> 12
		myVm.CheckInstr(op, instr)
	}
}
