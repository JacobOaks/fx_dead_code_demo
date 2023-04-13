This repo shows how Dig callbacks [this PR](https://github.com/uber-go/dig/pull/377) and fx Running events [this commit](https://github.com/JacobOaks/fx/commit/1532d1bd3c7d3926ee3e98d8d15e47206e7192ef#diff-a04d0b8fa4df9f46eb6a4ee7c75554d41ab5c2343bea18ba09fc62538953d9a7)
can be used to easily detect dead Fx constructors/decorators.

This is done through a custom `fxevent.Logger` that wraps your usual logger,
but additionally keeps track of functions that haven't been run yet.
This is implemented in `dead_fn_wrapper.go`.

The actual Fx app is created in `main.go`. The app contains a couple
functions that aren't used, which we print out, providing output:

```
Found 5 dead functions:
- MyModule:main.uint16ToUnit32()
- MySubModule:fx_dead_code_demo/subpkg.StringToBool()
- go.uber.org/fx.(*App).shutdowner-fm()
- go.uber.org/fx.New.func1()
- go.uber.org/fx.(*App).dotGraph-fm()
```

The bottom three are internal Fx functions that get provided to every app.
