package localization

// Config defines local language and measurement metrics settings.
type Config struct {
	ID            string `json:"id"`
	TenantID      string `json:"tenant_id"`
	Language      string `json:"language"`       // en, es, fr
	DateFormat    string `json:"date_format"`    // YYYY-MM-DD
	UnitSystem    string `json:"unit_system"`    // METRIC, IMPERIAL
	NumberDecimal string `json:"number_decimal"` // DOT, COMMA
}
