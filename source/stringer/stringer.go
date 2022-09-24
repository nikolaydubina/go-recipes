package painkiller

//go:generate stringer -type=Pill -linecomment

type Pill int

const (
	Placebo Pill = iota
	Ibuprofen
	Paracetamol
	PillAspirin   // Aspirin
	Acetaminophen = Paracetamol
)

// "Acetaminophen"
var s string = Acetaminophen.String()
