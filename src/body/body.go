package body

import (
	"fmt"
	"os"
	"time"

	"github.com/villekuosmanen/physiology-sim/src/organ"
	"github.com/villekuosmanen/physiology-sim/src/systems/circulation"
	"github.com/villekuosmanen/physiology-sim/src/systems/metabolism"
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
	superiorVenaCava := circulation.ConstructVessel(
		"Superior Vena Cava",
		circulation.VesselSizeLarge,
		[]circulation.ConsumerWithBloodSupply{
			{
				Consumer:    &heart.RightAtrium,
				BloodSupply: 1,
			},
		},
		false,
	)
	inferiorVenaCava := circulation.ConstructVessel(
		"Inferior Vena Cava",
		circulation.VesselSizeLarge,
		[]circulation.ConsumerWithBloodSupply{
			{
				Consumer:    &heart.RightAtrium,
				BloodSupply: 1,
			},
		},
		false,
	)

	// major organs
	brain := organ.ConstructBrain(superiorVenaCava)
	liver := organ.ConstructLiver(inferiorVenaCava)
	leftKidney := organ.ConstructKidney("Left Kidney", inferiorVenaCava)
	rightKidney := organ.ConstructKidney("Right Kidney", inferiorVenaCava)

	// limbs and torso
	leftBreast := organ.ConstructTorsoPart(
		"Left Breast",
		0.7,
		superiorVenaCava,
	)
	rightBreast := organ.ConstructTorsoPart(
		"Right Breast",
		0.7,
		superiorVenaCava,
	)
	abdomen := organ.ConstructTorsoPart(
		"Abdomen",
		0.5,
		inferiorVenaCava,
	)

	rightArm := organ.ConstructLimb(
		"Right Arm",
		0.8,
		superiorVenaCava,
	)
	leftArm := organ.ConstructLimb(
		"Left Arm",
		0.8,
		superiorVenaCava,
	)
	rightLeg := organ.ConstructLimb(
		"Right Leg",
		0.8,
		inferiorVenaCava,
	)
	leftLeg := organ.ConstructLimb(
		"Left Leg",
		0.8,
		inferiorVenaCava,
	)

	// lungs and pulmonary veins
	pulmonaryVein := circulation.ConstructVessel(
		"Pulmonary Vein",
		circulation.VesselSizeMedium,
		[]circulation.ConsumerWithBloodSupply{
			{
				Consumer:    &heart.LeftAtrium,
				BloodSupply: 1,
			},
		},
		false,
	)
	lungs := organ.ConstructLungs(pulmonaryVein)

	// arteries
	pulmonaryArtery := circulation.ConstructVessel(
		"Pulmonary Artery",
		circulation.VesselSizeMedium,
		[]circulation.ConsumerWithBloodSupply{
			{
				Consumer:    lungs,
				BloodSupply: 1,
			},
		},
		true,
	)
	aorta := circulation.ConstructVessel(
		"Aorta",
		circulation.VesselSizeHuge,
		[]circulation.ConsumerWithBloodSupply{
			{
				Consumer:    brain,
				BloodSupply: 5,
			},
			{
				Consumer:    liver,
				BloodSupply: 4,
			},
			{
				Consumer:    leftKidney,
				BloodSupply: 3,
			},
			{
				Consumer:    rightKidney,
				BloodSupply: 3,
			},
			{
				Consumer:    &heart.Myocardium,
				BloodSupply: 0.125,
			},
			{
				Consumer:    leftBreast,
				BloodSupply: 1,
			},
			{
				Consumer:    rightBreast,
				BloodSupply: 1,
			},
			{
				Consumer:    abdomen,
				BloodSupply: 2,
			},
			{
				Consumer:    rightArm,
				BloodSupply: 1,
			},
			{
				Consumer:    leftArm,
				BloodSupply: 1,
			},
			{
				Consumer:    rightLeg,
				BloodSupply: 1,
			},
			{
				Consumer:    leftLeg,
				BloodSupply: 1,
			},
		},
		true,
	)

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
	// get heart rate in frequency
	heartRateFreq := (heartRate / 60) * frequency

	// run forever in given Hz
	untilNextHeartbeat := 0.0

	var t *time.Ticker
	var i int64
	if realtime {
		t = time.NewTicker(time.Second / time.Duration(frequency))
	} else {
		// 100 times faster than otherwise
		t = time.NewTicker(time.Second / (time.Duration(frequency) * 100))
	}

	// Before we start printing stats, converge by running 1,000,000 runs
	fmt.Println("Starting simulation...")
	for i := 0; i < 1_000_000; i++ {
		if untilNextHeartbeat <= 0 {
			b.Heart.Beat()
			untilNextHeartbeat = heartRateFreq
		} else {
			untilNextHeartbeat -= 1
		}

		b.Act()
		i++
	}

	for {
		select {
		case <-t.C:
			if untilNextHeartbeat <= 0 {
				b.Heart.Beat()
				untilNextHeartbeat = heartRateFreq
			} else {
				untilNextHeartbeat -= 1
			}

			b.PrintStats()
			b.Act()
			i++

		case <-sigs:
			return
		}
	}
}

func (b *Body) Act() {
	b.Heart.Myocardium.Act()
	b.Aorta.Act()

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

	b.PulmonaryArtery.Act()
	b.Lungs.Act()
	b.PulmonaryVein.Act()

	b.SuperiorVenaCava.Act()
	b.InferiorVenaCava.Act()
}

func (b *Body) PrintStats() {
	total := 0.0
	fmt.Println("********************************")

	heartStats := b.Heart.MonitorHeart()
	for _, hs := range heartStats {
		total += hs.BloodQuantity
		hs.Print()
	}
	s := b.Aorta.Monitor()
	total += s.BloodQuantity
	s.Print()

	s = b.Brain.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.Liver.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.LeftKidney.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.RightKidney.Monitor()
	total += s.BloodQuantity
	s.Print()

	s = b.LeftBreast.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.RightBreast.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.Abdomen.Monitor()
	total += s.BloodQuantity
	s.Print()

	s = b.RightArm.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.LeftArm.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.RightLeg.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.LeftLeg.Monitor()
	total += s.BloodQuantity
	s.Print()

	s = b.PulmonaryArtery.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.Lungs.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.PulmonaryVein.Monitor()
	total += s.BloodQuantity
	s.Print()

	s = b.SuperiorVenaCava.Monitor()
	total += s.BloodQuantity
	s.Print()
	s = b.InferiorVenaCava.Monitor()
	total += s.BloodQuantity
	s.Print()

	fmt.Printf("****** TOTAL: %.2f *********\n", total)
	fmt.Println("********************************")
	fmt.Println("")
}

func (b *Body) SetMetabolicRate(m metabolism.MET) {
	// TODO:
	// Myocardium metabolic rate should be set by heart rate
	// The faster the heart rate the more it consumes
	b.LeftArm.Muscle.SetMetabolicRate(m)
	b.RightArm.Muscle.SetMetabolicRate(m)
	b.LeftLeg.Muscle.SetMetabolicRate(m)
	b.RightArm.Muscle.SetMetabolicRate(m)

	b.Abdomen.Muscle.SetMetabolicRate(m)
	b.RightBreast.Muscle.SetMetabolicRate(m)
	b.LeftBreast.Muscle.SetMetabolicRate(m)
}
