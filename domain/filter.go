package domain

type Filter struct {
	Eq string
}

type FilterType int

const (
	FilterType_UNKNOWN FilterType = iota
	FilterType_EQ
	FilterType_REGEX
)

func (f FilterType) String() string {
	switch f {
	case FilterType_EQ:
		return "EQ"
	case FilterType_REGEX:
		return "REGEX"
	default:
		return "UnKnown"
	}
}
