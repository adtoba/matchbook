package book

type Side uint8

const (
	SideUnknown Side = iota
	Buy
	Sell
)

func (s Side) better(a, b int64) bool {
	switch s {
	case Buy:
		return a > b
	case Sell:
		return a < b
	default:
		return false
	}
}

func (s Side) String() string {
	switch s {
	case Buy:
		return "buy"
	case Sell:
		return "sell"
	default:
		return "unknown"
	}
}
