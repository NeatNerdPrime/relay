package runner

import (
	"context"

	logging "github.com/puppetlabs/insights-logging"
	"github.com/puppetlabs/nebula/pkg/errors"
	"github.com/puppetlabs/nebula/pkg/io"
	"github.com/puppetlabs/nebula/pkg/state"
)

type ActionRuntime interface {
	IO() *io.IO
	Logger() logging.Logger
	StateManager() state.Manager
}

type ActionRunner interface {
	Run(ctx context.Context, resourceID string, runtime ActionRuntime, variables map[string]string) errors.Error
	Decoder() Decoder
}

type Decoder interface {
	Decode(b []byte) errors.Error
}

func NewRunner(kind RunnerKind) (ActionRunner, error) {
	switch kind {
	case RunnerKindGKEClusterProvisioner:
		return &GKEClusterProvisioner{}, nil
	case RunnerKindShell:
		return &Shell{}, nil
	case RunnerKindWorkflow:
		return &Workflow{}, nil
	}

	return nil, errors.NewWorkflowRunnerNotFound(string(kind))
}