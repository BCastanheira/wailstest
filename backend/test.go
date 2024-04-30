package backend

import (
	"log"
)

func (t *Test) CreateTest() {
	log.Println("Creating...")
	t.EventHandler.EmitEvent("env_launched")
}
