package constants

type ToxicType string

const (
	ToxicTypeHigh   ToxicType = "htoxic"
	ToxicTypeMedium ToxicType = "mtoxic"
	ToxicTypeLow    ToxicType = "ltoxic"
	NonToxic        ToxicType = "nontoxic"
)

func DetermineToxicType(toxicity float64) ToxicType {
	if toxicity >= 0.9 {
		return ToxicTypeHigh
	} else if toxicity >= 0.75 {
		return ToxicTypeMedium
	} else if toxicity >= 0.60 {
		return ToxicTypeLow
	}
	return NonToxic
}
