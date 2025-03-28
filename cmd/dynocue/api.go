package main

import "changeme/internal"

type DynoCueHeadless struct{}

func NewDynoCueHeadless() *DynoCueHeadless {
	return &DynoCueHeadless{}
}

type DynoCueService struct {
	internal.DynoCueApi
}

func NewDynoCueService() *DynoCueService {
	return &DynoCueService{}
}
