package manager

import (
	"context"

	"github.com/villekuosmanen/physiology-sim/src/body"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type BodySimManager struct {
	cancel context.CancelFunc
	body   *body.Body
}

func (m *BodySimManager) ResetSim(signalCtx context.Context, connManager body.Broadcaster) {
	if m.cancel != nil {
		m.cancel()
	}

	ctx, cancel := context.WithCancel(signalCtx)
	m.cancel = cancel

	body := body.ConstructBody(connManager)
	body.SetMetabolicRate(metabolism.METRest)
	m.body = body

	// launch in a separate goroutine
	go body.Run(ctx, false)
}

func (m *BodySimManager) SetExerciseLevel(level float64) {
	if m.body != nil {
		m.body.SetMetabolicRate(metabolism.MET(level))
	}
}

func (m *BodySimManager) ToggleFastForvard() {
	if m.body != nil {
		m.body.ToggleFastForvard()
	}
}
