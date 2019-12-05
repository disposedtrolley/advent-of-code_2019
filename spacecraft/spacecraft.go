package spacecraft

import (
	"fmt"

	"github.com/disposedtrolley/advent-of-code_2019/spacecraft/module"
)

type Spacecraft struct {
	Modules []*module.Module
}

func (s *Spacecraft) AddModule(m *module.Module) {
	s.Modules = append(s.Modules, m)
}

func (s *Spacecraft) TotalFuelRequired() (fuel int) {
	for _, m := range s.Modules {
		fuel += m.TotalFuel()
		fmt.Printf("totalFuel: %d\n", fuel)
	}

	return fuel
}
