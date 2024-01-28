package circulation

type Blood struct {
	// contains everything in a blood
	Quantity         float64 // percentage (100 is max amount)
	OxygenSaturation float64 // between 0 - 1
	LacticAcid       float64 // from 0

	// Neurotransmitters
	// Norepinephrine regulates Sympathetic Nervous System
	Norepinephrine float64 // between 0 - 1
}

// Extract removes a given fraction of blood in the system.
// It modifies the existing quantity and returns what was extracted.
func (b *Blood) Extract(fraction float64) Blood {
	qty := b.Quantity * fraction
	b.Quantity -= qty

	return Blood{
		Quantity:         qty,
		OxygenSaturation: b.OxygenSaturation,
		LacticAcid:       b.LacticAcid,
		Norepinephrine:   b.Norepinephrine,
	}
}

// Merge merges the given two Blood objects
func (b *Blood) Merge(a Blood) {
	total := b.Quantity + a.Quantity
	bFraction := b.Quantity / total
	if bFraction < 0.0001 {
		bFraction = 0
	} else if bFraction > 10000 {
		bFraction = 1
	}

	oxygenSat := (b.OxygenSaturation * bFraction) + (a.OxygenSaturation * (1 - bFraction))
	lactic := (b.LacticAcid * bFraction) + (a.LacticAcid * (1 - bFraction))
	norepinephrine := (b.Norepinephrine * bFraction) + (a.Norepinephrine * (1 - bFraction))

	b.Quantity = total
	if total == 0 {
		b.OxygenSaturation = 0
		b.LacticAcid = 0
		b.Norepinephrine = 0
	} else {
		b.OxygenSaturation = oxygenSat
		b.LacticAcid = lactic
		b.Norepinephrine = norepinephrine
	}
}

// RemoveFrom removes the given fraction's worth of blood from input.
func RemoveFrom(b Blood, fraction float64) Blood {
	qty := b.Quantity * fraction
	return Blood{
		Quantity:         qty,
		OxygenSaturation: b.OxygenSaturation,
		LacticAcid:       b.LacticAcid,
		Norepinephrine:   b.Norepinephrine,
	}
}
