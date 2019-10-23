package settings

import (
	"github.com/argoproj/argo-cd/engine/pkg"
	"github.com/argoproj/argo-cd/engine/pkg/apis/application/v1alpha1"
)

type noop_callbacks struct {
}

func (n noop_callbacks) OnSyncCompleted(appName string, state v1alpha1.OperationState) error {
	return nil
}

func NewNoOpCallbacks() pkg.Callbacks {
	return &noop_callbacks{}
}
