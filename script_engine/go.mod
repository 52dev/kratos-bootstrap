module github.com/52dev/kratos-bootstrap/script_engine

go 1.24.6

replace github.com/52dev/kratos-bootstrap/api => ../api

require (
	github.com/52dev/go-scripts v1.0.0
	github.com/52dev/kratos-bootstrap/api v1.0.1
)

require google.golang.org/protobuf v1.36.11 // indirect
