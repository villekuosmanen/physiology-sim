package circulation

type Blood struct {
	// contains everything in a blood
	Quantity         float64
	OxygenSaturation float64
}

// Extract removes a given fraction of blood in the system.
// It modifies the existing quantity and returns what was extracted.
func (b *Blood) Extract(fraction float64) Blood {
	qty := b.Quantity * fraction
	b.Quantity -= qty

	return Blood{
		Quantity:         qty,
		OxygenSaturation: b.OxygenSaturation,
	}
}

// Merge merges the given two Blood objects
func (b *Blood) Merge(a Blood) {
	total := b.Quantity + a.Quantity
	bFraction := b.Quantity / total
	oxygenSat := (b.OxygenSaturation * bFraction) + (a.OxygenSaturation * (1 - bFraction))

	b.Quantity = total
	if total == 0 {
		b.OxygenSaturation = 0
	} else {
		b.OxygenSaturation = oxygenSat
	}
}

// RemoveFrom removes the given fraction's worth of blood from input.
func RemoveFrom(b Blood, fraction float64) Blood {
	qty := b.Quantity * fraction
	return Blood{
		Quantity:         qty,
		OxygenSaturation: b.OxygenSaturation,
	}
}
