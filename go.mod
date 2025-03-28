module github.com/samber/lo

go 1.22.0

toolchain go1.23.7

//
// Dev dependencies are excluded from releases. Please check CI.
//

require (
	github.com/stretchr/testify v1.10.0
	github.com/thoas/go-funk v0.9.3
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.60.0
	go.uber.org/goleak v1.2.1
	golang.org/x/text v0.22.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.opentelemetry.io/otel/metric v1.35.0 // indirect
	go.opentelemetry.io/otel/trace v1.35.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
