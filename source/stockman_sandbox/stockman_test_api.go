package sandbox

import (
	"context"
	core "stockman/source/stockman_core"
)

type FirstTestDTO struct {
	TestFieldDTO string `json:"test_field_dto"`
}

func FirstServiceTestFn(ctx context.Context, event *core.Event) {
	s := "Hello services!"
	dto := FirstTestDTO{TestFieldDTO: s}
	event.SetOutput(dto)
	event.NotifyOutputChanged()
}
