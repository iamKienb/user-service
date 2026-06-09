package module

import (
	"user-worker-module/internal/application/port"
	"user-worker-module/internal/application/processor"
)

type ApplicationModule struct {
	EventProcessor port.EventProcessor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	return &ApplicationModule{
		EventProcessor: processor.NewUserEventProcessor(infra.ESRepo, infra.WorkerCache),
	}
}
