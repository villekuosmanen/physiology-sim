package body

import (
	"os"
	"time"

	"github.com/villekuosmanen/physiology-sim/src/organ"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
)

// TODO this should not be a constant
const heartRate = 80

type Body struct {
	Heart            *organ.Heart
	Aorta            *circulation.Vessel
	SuperiorVenaCava *circulation.Vessel
	InferiorVenaCava *circulation.Vessel

	PulmonaryArtery *circulation.Vessel
	PulmonaryVein   *circulation.Vessel
	Lungs           *organ.Lungs

	Brain       *organ.Brain
	Liver       *organ.Liver
	LeftKidney  *organ.Kidney
	RightKidney *organ.Kidney

	// To be added
	// digestive tract
	// Portal Vein

	LeftBreast  *organ.TorsoPart
	RightBreast *organ.TorsoPart
	Abdomen     *organ.TorsoPart

	RightArm *organ.Limb
	LeftArm  *organ.Limb
	RightLeg *organ.Limb
	LeftLeg  *organ.Limb
}

func ConstructBody() *Body {
	// heart and veins
	heart := organ.ConstructHeart()
	superiorVenaCava := circulation.ConstructVessel([]circulation.BloodConsumer{
		&heart.RightAtrium,
	}, false)
	inferiorVenaCava := circulation.ConstructVessel([]circulation.BloodConsumer{
		&heart.RightAtrium,
	}, false)

	// major organs
	brain := organ.ConstructBrain(superiorVenaCava)
	liver := organ.ConstructLiver(inferiorVenaCava)
	leftKidney := organ.ConstructKidney(inferiorVenaCava)
	rightKidney := organ.ConstructKidney(inferiorVenaCava)

	// limbs and torso
	leftBreast := organ.ConstructTorsoPart(0.7, superiorVenaCava)
	rightBreast := organ.ConstructTorsoPart(0.7, superiorVenaCava)
	abdomen := organ.ConstructTorsoPart(0.5, inferiorVenaCava)

	rightArm := organ.ConstructLimb(0.8, superiorVenaCava)
	leftArm := organ.ConstructLimb(0.8, superiorVenaCava)
	rightLeg := organ.ConstructLimb(0.8, inferiorVenaCava)
	leftLeg := organ.ConstructLimb(0.8, inferiorVenaCava)

	// lungs and pulmonary veins
	pulmonaryVein := circulation.ConstructVessel([]circulation.BloodConsumer{
		&heart.LeftAtrium,
	}, false)
	lungs := organ.ConstructLungs(pulmonaryVein)

	// arteries
	pulmonaryArtery := circulation.ConstructVessel([]circulation.BloodConsumer{
		lungs,
	}, true)
	aorta := circulation.ConstructVessel([]circulation.BloodConsumer{
		brain,
		liver,
		leftKidney,
		rightKidney,
		&heart.Myocardium,
		leftBreast,
		rightBreast,
		abdomen,
		rightArm,
		leftArm,
		rightLeg,
		leftLeg,
	}, true)

	// set consumers to heart
	heart.SetConsumers(aorta, pulmonaryArtery)

	return &Body{
		Heart:            heart,
		Aorta:            aorta,
		SuperiorVenaCava: superiorVenaCava,
		InferiorVenaCava: inferiorVenaCava,
		PulmonaryArtery:  pulmonaryArtery,
		PulmonaryVein:    pulmonaryVein,
		Lungs:            lungs,
		Brain:            brain,
		Liver:            liver,
		LeftKidney:       leftKidney,
		RightKidney:      rightKidney,
		LeftBreast:       leftBreast,
		RightBreast:      rightBreast,
		Abdomen:          abdomen,
		RightArm:         rightArm,
		LeftArm:          leftArm,
		RightLeg:         rightLeg,
		LeftLeg:          leftLeg,
	}
}

func (b *Body) Run(frequency float64, realtime bool, sigs <-chan os.Signal) {
	// run forever in given Hz
	untilNextHeartbeat := 0.0

	var t *time.Ticker
	if realtime {
		t = time.NewTicker(time.Second / time.Duration(frequency))
	} else {
		// 100 times faster than otherwise
		t = time.NewTicker(time.Second / (time.Duration(frequency) * 100))
	}

	for {
		select {
		case <-t.C:
			if untilNextHeartbeat <= 0 {
				b.Heart.Beat()
				untilNextHeartbeat = heartRate * frequency
			} else {
				untilNextHeartbeat -= 1
			}
			b.Act()

		case <-sigs:
			return
		}
	}
}

func (b *Body) Act() {
	b.Heart.Myocardium.Act()
	b.Aorta.Act()

	b.PulmonaryArtery.Act()
	b.Lungs.Act()
	b.InferiorVenaCava.Act()

	b.Brain.Act()
	b.Liver.Act()
	b.LeftKidney.Act()
	b.RightKidney.Act()

	b.LeftBreast.Act()
	b.RightBreast.Act()
	b.Abdomen.Act()

	b.RightArm.Act()
	b.LeftArm.Act()
	b.RightLeg.Act()
	b.LeftLeg.Act()

	b.SuperiorVenaCava.Act()
	b.InferiorVenaCava.Act()
}
