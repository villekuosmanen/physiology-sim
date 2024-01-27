package circulation

type Blood struct {
	// contains everything in a blood
	Quantity float64
}

// Extract removes a given fraction of blood in the system.
// It modifies the existing quantity and returns what was extracted.
func (b *Blood) Extract(fraction float64) Blood {
	qty := b.Quantity * fraction
	b.Quantity -= qty

	return Blood{
		Quantity: qty,
	}
}

// Merge merges the given two Blood objects
func (b *Blood) Merge(a Blood) {
	b.Quantity += a.Quantity
}

// DivideBlood divides a blood into even fractions, as represented by the blood unit returned.
func DivideBlood(a Blood, shares int) Blood {
	qty := a.Quantity / float64(shares)
	return Blood{
		Quantity: qty,
	}
}
