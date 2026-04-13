module github.com/52dev/kratos-bootstrap/registry

go 1.24.6

replace github.com/52dev/kratos-bootstrap/api => ../api

require (
	github.com/52dev/kratos-bootstrap/api v1.0.0
	github.com/go-kratos/kratos/v2 v2.9.2
)

require google.golang.org/protobuf v1.36.11 // indirect
