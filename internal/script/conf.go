package script

const (
	MethodHepburn  method = "Hepburn"
	MethodKunrei   method = "Kunrei"
	MethodPassport method = "Passport"
)

const (
	Mode_a mode = "a"
	ModeE  mode = "E"
	ModeH  mode = "H"
	ModeK  mode = "K"
)

type Conf struct {
	Method method
	Mode   mode
}

type method string

type mode string
