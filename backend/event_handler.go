package backend

import(
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type EventHandler struct {
    Ctx context.Context
}

func NewEventHandler() *EventHandler {
    return &EventHandler{}
}

func (eh *EventHandler) Startup(ctx context.Context) {
    eh.Ctx = ctx
}

func (eh *EventHandler) EmitEvent(event string) {
    runtime.EventsEmit(eh.Ctx, event)
}