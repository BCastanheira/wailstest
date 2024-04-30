package backend

import(
	"log"
	"context"
)

var eh = NewEventHandler()

func (t *Test) CreateTest() {
	log.Println("Creating...")
  	eh.EmitEvent("env_launched")
}

func (t *Test) startup(ctx context.Context) {
	t.ctx = ctx
}
