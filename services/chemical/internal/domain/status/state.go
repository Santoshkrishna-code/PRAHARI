package status

// Code represents chemical lifecycle state.
type Code string

const (
	CodeRequested           Code = "REQUESTED"
	CodeTechnicalReview     Code = "TECHNICAL_REVIEW"
	CodeSafetyReview        Code = "SAFETY_REVIEW"
	CodeEnvironmentalReview Code = "ENVIRONMENTAL_REVIEW"
	CodeApproved            Code = "APPROVED"
	CodePurchased           Code = "PURCHASED"
	CodeReceived            Code = "RECEIVED"
	CodeInspection          Code = "INSPECTION"
	CodeStored              Code = "STORED"
	CodeIssued              Code = "ISSUED"
	CodeInUse               Code = "IN_USE"
	CodeTransferred         Code = "TRANSFERRED"
	CodeReturned            Code = "RETURNED"
	CodeWasteClassification Code = "WASTE_CLASSIFICATION"
	CodeWasteProcessing     Code = "WASTE_PROCESSING"
	CodeDisposed            Code = "DISPOSED"
	CodeClosed              Code = "CLOSED"
	CodeRejected            Code = "REJECTED"
	CodeExpired             Code = "EXPIRED"
	CodeQuarantined         Code = "QUARANTINED"
	CodeRecalled            Code = "RECALLED"
)

func (c Code) String() string {
	return string(c)
}
