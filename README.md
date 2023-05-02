This repo shows how Dig callbacks from [this PR](https://github.com/uber-go/dig/pull/377) and fx Running events from [this PR](https://github.com/uber-go/fx/pull/1077) can be used to easily identify dead Fx constructors/decorators.

This is done through a custom `fxevent.Logger` that wraps your usual logger,
but additionally keeps track of functions that haven't been run yet.
This is implemented in `dead_fn_wrapper.go`.

The actual Fx app is created in `main.go`. The app contains a couple
functions that aren't used, which we print out, providing output:

```
Found 5 dead functions:
 - "MySubModule:fx_dead_code_demo/subpkg.StringToBool()" from:
        fx_dead_code_demo/subpkg.init (/home/user/go/src/github.com/JacobOaks/fx_dead_code_demo/subpkg/subpkg.go:11)
                 runtime.doInit (/opt/go/root/src/runtime/proc.go:6506)
                 runtime.doInit (/opt/go/root/src/runtime/proc.go:6483)
                 runtime.main (/opt/go/root/src/runtime/proc.go:233)

 - "MyModule:main.uint16ToUint32()" from:
        main.init (/home/user/go/src/github.com/JacobOaks/fx_dead_code_demo/main.go:25)
                 runtime.doInit (/opt/go/root/src/runtime/proc.go:6506)
                 runtime.main (/opt/go/root/src/runtime/proc.go:233)

 - "go.uber.org/fx.New.func1()" from:
        go.uber.org/fx.New (/home/user/go-repos/pkg/mod/github.com/!jacob!oaks/fx@v0.0.0-20230502170936-64fdf63ce654/app.go:475)
                 main.main (/home/user/go/src/github.com/JacobOaks/fx_dead_code_demo/main.go:36)
                 runtime.main (/opt/go/root/src/runtime/proc.go:250)

 - "go.uber.org/fx.(*App).shutdowner-fm()" from:
        go.uber.org/fx.New (/home/user/go-repos/pkg/mod/github.com/!jacob!oaks/fx@v0.0.0-20230502170936-64fdf63ce654/app.go:475)
                 main.main (/home/user/go/src/github.com/JacobOaks/fx_dead_code_demo/main.go:36)
                 runtime.main (/opt/go/root/src/runtime/proc.go:250)

 - "go.uber.org/fx.(*App).dotGraph-fm()" from:
        go.uber.org/fx.New (/home/user/go-repos/pkg/mod/github.com/!jacob!oaks/fx@v0.0.0-20230502170936-64fdf63ce654/app.go:475)
                 main.main (/home/user/go/src/github.com/JacobOaks/fx_dead_code_demo/main.go:36)
                 runtime.main (/opt/go/root/src/runtime/proc.go:250)
```

The bottom three are internal Fx functions that get provided to every app.
