package timezone

// Zone defines time offset details for localization.
type Zone struct {
	ID         string `json:"id"`
	Code       string `json:"code"` // UTC, IST, EST
	OffsetSecs int    `json:"offset_secs"`
}
