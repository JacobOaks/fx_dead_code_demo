package main

import (
	"fmt"
	"os"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"fx_dead_code_demo/subpkg"
)



// Random functions we will provide to Fx
func getInt() int { return 5 }
func doubleInt(a int) int { return a * 2 }
func intToFloat(a int) float64 { return float64(a) }
func uint16ToUint32(a uint16) uint32 { return uint32(a) }
func printString(s string) { fmt.Println(s) }

var myModule = fx.Module(
	"MyModule",
	subpkg.SubModule,
	fx.Provide(
		getInt,         // Used
		intToFloat,     // Used
		uint16ToUint32, // Not Used
	),
	fx.Decorate(doubleInt), // Used
	fx.Invoke(printString),
)

func main() {
	var dfw *deadFunctionWrapper
	app := fx.New(
		myModule,
		fx.WithLogger(func() fxevent.Logger { 
			dfw = NewDeadFunctionWrapper(
				&fxevent.ConsoleLogger{
					W: os.Stderr,
				},
			)
			return dfw
		}),
	)

	if err := app.Err(); err != nil {
		os.Exit(1)
	}

	dead := dfw.Dead()
	fmt.Printf("Found %d dead functions:\n", len(dead))
	for _, fn := range dead {
		fmt.Printf(" - %v\n", fn)
	}

	// Output:

	// Found 5 dead functions:
	// - MyModule:main.uint16ToUnit32()
	// - MySubModule:fx_dead_code_demo/subpkg.StringToBool()
	// - go.uber.org/fx.(*App).shutdowner-fm()
	// - go.uber.org/fx.New.func1()
	// - go.uber.org/fx.(*App).dotGraph-fm()

	// (bottom three are internal Fx functions that get provided to every app)
}


