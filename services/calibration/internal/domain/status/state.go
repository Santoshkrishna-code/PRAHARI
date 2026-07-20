package status

// Code represents Calibration lifecycle state.
type Code string

const (
	CodeRegistered            Code = "REGISTERED"
	CodeScheduled             Code = "SCHEDULED"
	CodeCalibrationStarted    Code = "CALIBRATION_STARTED"
	CodeMeasurementRecorded   Code = "MEASUREMENT_RECORDED"
	CodeToleranceVerification Code = "TOLERANCE_VERIFICATION"
	CodeCertificateGenerated  Code = "CERTIFICATE_GENERATED"
	CodeApproved              Code = "APPROVED"
	CodeActive                Code = "ACTIVE"
	CodeFailed                Code = "FAILED"
	CodeOutOfTolerance        Code = "OUT_OF_TOLERANCE"
	CodeRetired               Code = "RETIRED"
)

func (c Code) String() string {
	return string(c)
}
