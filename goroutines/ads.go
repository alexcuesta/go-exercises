package goroutines

type AdSource int

// Stringer interface
func (a AdSource) String() string {
	switch a {
	case Google:
		return "Google"
	case Amazon:
		return "Amazon"
	case Meta:
		return "Meta"
	default:
		return "Unkown"
	}
}

const (
	Google AdSource = iota
	Meta
	Amazon
)

var AllAdSources = []AdSource{Google, Meta, Amazon}

type Ad struct {
	id   int
	name string
}
