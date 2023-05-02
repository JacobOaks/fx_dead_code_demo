module fx_dead_code_demo

go 1.20

require go.uber.org/fx v0.0.0-00010101000000-000000000000

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace go.uber.org/fx => github.com/JacobOaks/fx v0.0.0-20230502170936-64fdf63ce654

replace go.uber.org/dig => go.uber.org/dig v1.16.2-0.20230501184430-027aa21628f4
