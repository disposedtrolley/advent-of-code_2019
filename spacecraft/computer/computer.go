package computer

type Opcode int

const (
	ADD  Opcode = 1
	MULT Opcode = 2
	HALT Opcode = 99
)

type Computer struct {
	memory          []int
	ip              int
	instructionSize int
}

func NewComputer(instructionSize int) *Computer {
	return &Computer{ip: 0, instructionSize: instructionSize}
}

func (c *Computer) LoadProgram(p []int) {
	c.ip = 0
	c.memory = p
}

func (c *Computer) advanceIP() {
	c.ip += c.instructionSize
}

func (c *Computer) nextInstruction() Opcode {
	opcode := Opcode(c.read(c.ip))
	return opcode
}

func (c *Computer) processInstruction(opcode Opcode) (ended bool) {
	if opcode == HALT {
		return true
	}

	inputA := c.read(c.read(c.ip + 1))
	inputB := c.read(c.read(c.ip + 2))
	outputIndex := c.read(c.ip + 3)

	switch opcode {
	case ADD:
		c.write(outputIndex, inputA+inputB)
	case MULT:
		c.write(outputIndex, inputA*inputB)
	}

	return false
}

func (c *Computer) read(addr int) int {
	return c.memory[addr]
}

func (c *Computer) write(addr int, val int) {
	c.memory[addr] = val
}

func (c *Computer) Execute() (result int) {
	opcode := c.nextInstruction()
	for {
		if shouldHalt := c.processInstruction(opcode); shouldHalt {
			break
		}
		c.advanceIP()
		opcode = c.nextInstruction()
	}

	return c.read(0)
}
