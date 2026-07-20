package secrets

// RotationEvent maps standard JSON payloads received by AWS Secrets Manager rotation lambdas.
type RotationEvent struct {
	SecretID           string `json:"SecretId"`
	ClientRequestToken string `json:"ClientRequestToken"`
	Step               string `json:"Step"` // e.g. createSecret, setSecret, testSecret, finishSecret
}

// RotationStep defines types for rotation callbacks.
type RotationStep string

const (
	StepCreateSecret RotationStep = "createSecret"
	StepSetSecret    RotationStep = "setSecret"
	StepTestSecret   RotationStep = "testSecret"
	StepFinishSecret RotationStep = "finishSecret"
)
