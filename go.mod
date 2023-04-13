module fx_dead_code_demo

go 1.20

replace go.uber.org/fx => github.com/JacobOaks/fx v0.0.0-20230323170247-1532d1bd3c7d

require go.uber.org/fx v0.0.0-00010101000000-000000000000

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace go.uber.org/dig => github.com/JacobOaks/dig v1.15.1-0.20230322221759-1f96d68bb026
