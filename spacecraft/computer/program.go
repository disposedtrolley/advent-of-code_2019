package computer

type Opcode int

const (
	ADD  Opcode = 1
	MULT Opcode = 2
	HALT Opcode = 99
)

type IntcodeProgram struct {
	instructions []int
	index        int
	ipOffset     int
}

func NewIntcodeProgram(instructions []int, ipOffset int) *IntcodeProgram {
	return &IntcodeProgram{instructions: instructions, index: -(ipOffset), ipOffset: ipOffset}
}

func (p *IntcodeProgram) nextOpcode() Opcode {
	p.index += p.ipOffset
	return Opcode(p.readFromPos(p.index))
}

func (p *IntcodeProgram) processInstruction(opcode Opcode) (ended bool) {
	if opcode == HALT {
		return true
	}

	inputA := p.readFromPos(p.readFromPos(p.index + 1))
	inputB := p.readFromPos(p.readFromPos(p.index + 2))
	outputIndex := p.readFromPos(p.index + 3)

	switch opcode {
	case ADD:
		p.writeToPos(outputIndex, inputA+inputB)
	case MULT:
		p.writeToPos(outputIndex, inputA*inputB)
	}

	return false
}

func (p *IntcodeProgram) readFromPos(pos int) int {
	return p.instructions[pos]
}

func (p *IntcodeProgram) writeToPos(pos int, val int) {
	p.instructions[pos] = val
}

func (p *IntcodeProgram) Execute() (result int) {
	// Restore 1202 program alarm state
	//p.writeToPos(1, 12)
	//p.writeToPos(2, 2)

	opcode := p.nextOpcode()
	for {
		if shouldHalt := p.processInstruction(opcode); shouldHalt {
			break
		}
		opcode = p.nextOpcode()
	}

	return p.readFromPos(0)
}
