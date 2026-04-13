module github.com/52dev/kratos-bootstrap/database/mongodb

go 1.25.0

replace github.com/52dev/kratos-bootstrap/api => ../../api

require (
	github.com/52dev/go-crud/mongodb v1.0.1
	github.com/52dev/kratos-bootstrap/api v1.0.0
	github.com/go-kratos/kratos/v2 v2.9.2
	go.mongodb.org/mongo-driver/v2 v2.5.0
)

require (
	github.com/52dev/go-crud/api v1.0.1 // indirect
	github.com/52dev/go-crud/pagination v1.0.1 // indirect
	github.com/52dev/go-utils v1.0.1 // indirect
	github.com/52dev/go-utils/mapper v1.0.1 // indirect
	github.com/google/gnostic v0.7.1 // indirect
	github.com/google/gnostic-models v0.7.1 // indirect
	github.com/jinzhu/copier v0.4.0 // indirect
	github.com/klauspost/compress v1.18.4 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.2.0 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.einride.tech/aip v0.81.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/crypto v0.50.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
