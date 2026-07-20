package config

// FeatureFlagConfig maps standard boolean toggles enabling or disabling system capabilities.
type FeatureFlagConfig struct {
	EnableMultiTenancy            bool `env:"FF_ENABLE_MULTI_TENANCY" envDefault:"false"`
	EnableWorkerVitalsAnalytics   bool `env:"FF_ENABLE_WORKER_VITALS" envDefault:"false"`
	EnableDigitalTwinSimulations  bool `env:"FF_ENABLE_DIGITAL_TWIN" envDefault:"false"`
	EnableAutomaticPermitApproval bool `env:"FF_ENABLE_AUTO_PERMIT" envDefault:"false"`
}
