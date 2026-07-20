package inference

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Engine struct {
	modelPath string
}

func NewEngine(modelPath string) *Engine {
	return &Engine{modelPath: modelPath}
}

func (e *Engine) RunGPUInference(ctx context.Context, rawFrame []byte) (string, float64, error) {
	prahariLogger.Info(ctx, "Scheduled GPU accelerated model tensor inference execution step")
	// return mock detection label: no_helmet, confidence 0.96
	return "no_helmet", 0.96, nil
}
