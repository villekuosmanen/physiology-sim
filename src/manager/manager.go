package manager

import (
	"context"

	"github.com/villekuosmanen/physiology-sim/src/body"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
)

type BodySimManager struct {
	cancel context.CancelFunc
}

func (m *BodySimManager) ResetSim(signalCtx context.Context, connManager body.Broadcaster) {
	if m.cancel != nil {
		m.cancel()
	}

	ctx, cancel := context.WithCancel(signalCtx)
	m.cancel = cancel

	body := body.ConstructBody(connManager)
	body.SetMetabolicRate(metabolism.METHeavyCardio)

	// launch in a separate goroutine
	go body.Run(ctx, 10, true, false)
}
