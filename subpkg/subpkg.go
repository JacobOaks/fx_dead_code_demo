package subpkg

import (
	"fmt"

	"go.uber.org/fx"
)

var SubModule = fx.Module(
	"MySubModule",
	fx.Provide(FloatToString, StringToBool),
)

func FloatToString(f float64) string {
	return fmt.Sprint(f)
}

func StringToBool(s string) bool {
	return s == "true"
}
