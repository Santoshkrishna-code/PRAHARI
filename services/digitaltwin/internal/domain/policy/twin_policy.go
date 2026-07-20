package policy

// ValidateSimulationParameters checks that simulation parameters are safe to execute.
func ValidateSimulationParameters(params string) bool {
	if len(params) == 0 {
		return false
	}
	if len(params) > 50000 {
		return false // reject oversized payloads
	}
	return true
}

// ValidateTelemetryQuality checks that a telemetry signal quality is acceptable.
func ValidateTelemetryQuality(quality string) bool {
	return quality == "GOOD" || quality == "UNCERTAIN"
}
