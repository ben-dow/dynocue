package main

import (
	"dynocue/internal/appdef"
	"dynocue/internal/localapp"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DynoCueService struct {
	appdef.DynoCueApplication
	app *application.App
}

func (d *DynoCueService) setApp(app *application.App) {
	d.app = app
}

func (d *DynoCueService) evCallback(ev string, data interface{}) {
	d.app.EmitEvent(ev, data)
}

func (d *DynoCueService) NewLocal(path string) error {
	app, err := localapp.NewLocalDynoCue(path, d.evCallback)
	d.DynoCueApplication = app
	return err
}

func (d *DynoCueService) OpenLocal(path string) error {
	app, err := localapp.OpenLocalDynoCue(path, d.evCallback)
	d.DynoCueApplication = app
	return err
}

func NewDynoCueService() *DynoCueService {
	return &DynoCueService{
		DynoCueApplication: appdef.NoopDynoCueApplication{},
	}
}
