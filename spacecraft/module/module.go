package module

import "math"

type Module struct {
	Mass int
}

func (m *Module) TotalFuel() (fuel int) {
	fuel = m.fuelForMass(m.Mass)

	fuelForFuel := m.fuelForMass(fuel)
	for fuelForFuel >= 0 {
		fuel += fuelForFuel
		fuelForFuel = m.fuelForMass(fuelForFuel)
	}

	return fuel
}

func (m *Module) fuelForMass(mass int) (fuel int) {
	return int(math.Floor(float64(mass)/3) - 2)
}
