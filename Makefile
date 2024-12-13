all:
	make clean
	make flush_redis
	make proto_build
	make c2s
	make c2s_manager
	make asset


proto_build:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative ./proto_defs/proto_defs.proto

flush_redis:
	redis-cli --scan --pattern '*' | xargs redis-cli DEL

asset:
	go clean -cache
	go build -o asset/bin/asset client/*.go
	GOOS=windows GOARCH=amd64 go build -o asset/bin/asset.exe  asset/*.go

c2s:
	go clean -cache
	go build -o server/bin/server server/*.go

c2s_manager:
	go clean -cache 
	go build -o manager/bin/manager manager/*.go

.IGNORE clean :
	rm */bin/*
	rm ./proto_defs/*.go
	rm -rf ~/.customC2/*