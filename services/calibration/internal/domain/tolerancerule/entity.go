package tolerancerule

// Rule defines acceptable measurement tolerance ranges for an instrument class (e.g. ±0.5% span).
type Rule struct {
	ID             string  `json:"id"`
	InstrumentType string  `json:"instrument_type"`
	MinLimit       float64 `json:"min_limit"`
	MaxLimit       float64 `json:"max_limit"`
	UnitOfMeasure  string  `json:"unit_of_measure"`
}
