package main

import (
	"dynocue/internal/app"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DynoCueService struct {
	app.DynoCueApplication
	app *application.App
}

func (d *DynoCueService) setApp(app *application.App) {
	d.app = app
}

func (d *DynoCueService) evCallback(ev string, data interface{}) {
	d.app.EmitEvent(ev, data)
}

func (d *DynoCueService) NewLocal() {
	d.DynoCueApplication = app.NewLocalDynoCue(d.evCallback)
}

func (d *DynoCueService) OpenLocal() {
	d.DynoCueApplication = app.NewLocalDynoCue(d.evCallback)
}

func NewDynoCueService() *DynoCueService {
	return &DynoCueService{
		DynoCueApplication: app.NoopDynoCueApplication{},
	}
}
