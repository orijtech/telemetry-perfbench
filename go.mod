module github.com/orijtech/telemetry-perfbench

go 1.14

require (
	github.com/rs/zerolog v1.18.0
	github.com/sirupsen/logrus v1.5.0
	go.uber.org/zap v1.14.1
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/tools v0.0.0-20191029190741-b9c20aec41a5
)

replace golang.org/x/tools => ./internal/x_tools
