package main

import (
	"dynocue/pkg/api"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type DynoCueService struct {
	api.Api
	apiType  string
	wailsApp *application.App
}

func NewDynoCueService() *DynoCueService {
	return &DynoCueService{
		Api:     &api.Noop{},
		apiType: "noop",
	}
}

func (d *DynoCueService) setApp(app *application.App) {
	d.wailsApp = app
}

func (d *DynoCueService) evCallback(ev string, data interface{}) {
	d.wailsApp.EmitEvent(ev, data)
}

func (d *DynoCueService) ConnectAsHost() {
	d.apiType = "client"
}

func (d *DynoCueService) CreateNewAsHost(path string) {
	d.apiType = "host"
}

func (d *DynoCueService) OpenAsHost(path string) {
	d.apiType = "host"
}
