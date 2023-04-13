package main

import "go.uber.org/fx/fxevent"

// deadFunctionWrapper wraps an [fxevent.Logger] and also keeps track of
// which Fx constructors/decorators haven't yet run.
type deadFunctionWrapper struct{
	fns map[string]struct{}

	wrapped fxevent.Logger
}

var _ fxevent.Logger = &deadFunctionWrapper{}

func (dfw *deadFunctionWrapper) getKey(module, function string) string {
	key := function
	if module != "" {
		key = module + ":" + key
	}
	return key
}

// LogEvent updates our running list of provided/decorated functions,
// and then delegates actual logging to the wrapped logger.
func (dfw *deadFunctionWrapper) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Provided:
		dfw.fns[dfw.getKey(e.ModuleName, e.ConstructorName)] = struct{}{}
	case *fxevent.Decorated:
		dfw.fns[dfw.getKey(e.ModuleName, e.DecoratorName)] = struct{}{}
	case *fxevent.Run:
		delete(dfw.fns, dfw.getKey(e.ModuleName, e.Name))
	}
	dfw.wrapped.LogEvent(event)
}

// Dead prints out all functions that haven't been called.
func (dfw *deadFunctionWrapper) Dead() []string {
	dead := make([]string, 0, len(dfw.fns))
	for fn := range dfw.fns {
		dead = append(dead, fn)
	}
	return dead
}

// NewDeadFunctionWrapper 
func NewDeadFunctionWrapper(wrapped fxevent.Logger) *deadFunctionWrapper {
	return &deadFunctionWrapper{
		fns:     make(map[string]struct{}),
		wrapped: wrapped,
	}
}
