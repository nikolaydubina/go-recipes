package color

type Color struct{ c uint }

var (
	Undefined = Color{}
	Red       = Color{1}
	Green     = Color{2}
	Blue      = Color{3}
)
