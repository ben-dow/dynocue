package app

type NoopDynoCueApplication struct {
	DynoCueApplication
}

func NewNoopDynoCueApplication() *NoopDynoCueApplication {
	return &NoopDynoCueApplication{}
}
