package main

import (
	"fmt"
	"strings"

	"go.uber.org/fx/fxevent"
)

type function struct {
	name  string
	trace []string
}

// deadFunctionWrapper wraps an [fxevent.Logger] and also keeps track of
// which Fx constructors/decorators haven't yet run.
type deadFunctionWrapper struct{
	fns map[string]*function

	wrapped fxevent.Logger
}

var _ fxevent.Logger = &deadFunctionWrapper{}

func getKey(module, function string) string {
	key := function
	if module != "" {
		key = module + ":" + key
	}
	return key
}

func (dfw *deadFunctionWrapper) addFn(module, name string, trace string) {
	key := getKey(module, name)
	dfw.fns[key] = &function{
		name:  key,
		trace: strings.Split(trace, ";"),
	}
}

func (dfw *deadFunctionWrapper) rmFn(module, name string) {
	delete(dfw.fns, getKey(module, name))
}

// LogEvent updates our running list of provided/decorated functions,
// and then delegates actual logging to the wrapped logger.
func (dfw *deadFunctionWrapper) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Provided:
		dfw.addFn(e.ModuleName, e.ConstructorName, e.StackTrace)
	case *fxevent.Decorated:
		dfw.addFn(e.ModuleName, e.DecoratorName, e.StackTrace)
	case *fxevent.Run:
		dfw.rmFn(e.ModuleName, e.Name)
	}
	dfw.wrapped.LogEvent(event)
}

// PrintReport reports all functions that have not been invoked by Fx yet.
func (dfw *deadFunctionWrapper) PrintReport() {
	fmt.Printf("Found %d dead functions:\n", len(dfw.fns))
	for _, fn := range dfw.fns {
		fmt.Printf(" - %q from:\n", fn.name)
		for i, line := range fn.trace {
			tabs := "\t"
			if i != 0 {
				tabs += "\t"
			}
			fmt.Printf("%v%v\n", tabs, line)
		}
		fmt.Println()
	}
}

// NewDeadFunctionWrapper 
func NewDeadFunctionWrapper(wrapped fxevent.Logger) *deadFunctionWrapper {
	return &deadFunctionWrapper{
		fns:     make(map[string]*function),
		wrapped: wrapped,
	}
}
