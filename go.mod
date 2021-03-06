module mic_srv_office

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.5.1
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/protobuf v1.26.0
)
